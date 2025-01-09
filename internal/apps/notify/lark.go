package notify

import (
	"encoding/json"
	"groot/internal/config"
	response2 "groot/internal/response"
)

// Lark send message via lark
func (m *Message) Lark() error {
	/* demo
	messageTitle := InnerContent{Tag: "text", Text: "Title: " + p.PullRequest.Title + "\n"}
	messageUrlName := InnerContent{Tag: "text", Text: "URL: "}
	messageUrl := InnerContent{Tag: "a", Text: p.PullRequest.URL + "\n", Href: p.PullRequest.URL}
	messageRepository := InnerContent{Tag: "text", Text: "Repo: " + p.Repository.Name + "\n"}
	messageMRDest := InnerContent{Tag: "text", Text: "MRDest: " + p.PullRequest.Head.Label + " -> " + p.PullRequest.Base.Label + "\n"}
	messageCommitter := InnerContent{Tag: "text", Text: "Committer: " + p.Sender.Login + "\n"}
	messageReviewerName := InnerContent{Tag: "text", Text: "Reviewer: "}
	messageReviewer := InnerContent{Tag: "at", Text: reviewer}
	createAt := p.PullRequest.CreatedAt.Add(8 * time.Hour)
	messageCreateAt := InnerContent{Tag: "text", Text: "\nCreateAt: " + createAt.Format("2006-01-02 15:04:05")}
	*/
	//zhCN := ZhCN{Title: "merge request", Content: [][]InnerContent{{messageTitle, messageUrlName, messageUrl, messageRepository, messageMRDest, messageCommitter, messageReviewerName, messageReviewer, messageCreateAt}}}
	zhCN := ZhCN{Title: "Task: ", Content: [][]InnerContent{{}}}
	post := Post{ZhCN: zhCN}
	content := Content{Post: post}
	data := LarkMessage{MsgType: "post", Content: content}

	messageBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := config.GetString("notify.larkWebhook")
	err = response2.PostRequest(url, messageBytes)
	if err != nil {
		return err
	}
	return nil
}
