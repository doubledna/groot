package ansible

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/hibiken/asynq"
	"groot/internal/config"
	"groot/internal/zlog"
	"os"
	"time"
)

const (
	TypeAnsible = "ansible"
)

type AnsibleTaskPayload struct {
	Repo      string
	Branch    string
	Playbook  string
	Inventory string
}

func HandleAnsibleTask(ctx context.Context, t *asynq.Task) error {
	var htp AnsibleTaskPayload
	if err := json.Unmarshal(t.Payload(), &htp); err != nil {
		return err
	}
	// get config
	privateKeyFile := config.GetString("ansible.privateKeyFile")
	dataDir := config.GetString("ansible.dataDir")
	authenticate := config.GetString("ansible.authenticate")

	// authenticate
	var authMethod transport.AuthMethod
	// using privateKey
	if authenticate == "privateKey" {
		_, err := os.Stat(privateKeyFile)
		if err != nil {
			zlog.Errorf("read private key file", err)
			return err
		}
		authMethod, err = ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
		if err != nil {
			zlog.Errorf("generate public key failed", err)
			return err
		}
	} else if authenticate == "kerberos" {
		// using kerberos

	}

	// pull git repository
	repoDir := dataDir + time.Now().Format("2006-01-02_15-04-05")
	_, err := git.PlainClone(repoDir, false, &git.CloneOptions{
		URL:           htp.Repo,
		Auth:          authMethod,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", htp.Branch)),
		SingleBranch:  true,
		Progress:      os.Stdout,
	})
	if err != nil {
		zlog.Errorf("pull git repository", err)
		return err
	}

	// run ansible
	inventoryPath := repoDir + "/" + htp.Inventory
	playbookPath := repoDir + "/" + htp.Playbook
	err = runPlaybook(inventoryPath, playbookPath)
	if err != nil {
		zlog.Errorf("run ansible playbook error", err)
		return err
	}

	return nil
}

func runPlaybook(inventorys, playbooks string) error {
	// get config
	user := config.GetString("ansible.user")
	becomeUser := config.GetString("ansible.becomeUser")

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Become: true,
		//Connection: "local",
		User:       user,
		BecomeUser: becomeUser,
		Inventory:  inventorys,
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbooks),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
	)

	err := exec.Execute(context.Background())
	if err != nil {
		return err
	}
	return nil
}
