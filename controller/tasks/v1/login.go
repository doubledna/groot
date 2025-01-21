package v1

import (
	genv1 "groot/gen/v1"
	"groot/internal"
	"groot/internal/models/user"
	"groot/internal/response"
	"groot/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"
)

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (t *TaskStore) LoginUser(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		ErrorFormat(c, response.CodeLoginFailed.WithError(err))
		return
	}

	u := user.User{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := LoginCheck(u.Username, u.Password)
	if err != nil {
		ErrorFormat(c, response.CodeLoginFailed.WithError(err))
		return
	}

	tokens := []genv1.Token{
		{
			Token:    token,
			Username: u.Username,
		},
	}

	SuccessFormat(c, tokens, *response.CodeSuccess)
}

func LoginCheck(username string, password string) (string, error) {
	u := user.User{}
	var err error
	err = internal.DB.Model(&user.User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyPassword(password string, userPassword string) error {
	if password != userPassword {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func SuccessFormat(c *gin.Context, tokens []genv1.Token, response response.Response) {
	result := make([]genv1.Token, 0, len(tokens))
	for _, value := range tokens {
		result = append(result, genv1.Token{
			Token:    value.Token,
			Username: value.Username,
		})
	}

	resp := genv1.GetToken{
		Code:      int64(response.Code),
		Message:   response.Message,
		Reference: response.Reference,
		Error:     response.Error,
		Data:      result,
	}

	c.JSON(http.StatusOK, resp)
}
