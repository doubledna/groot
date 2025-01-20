package v1

import (
	"go/token"
	genv1 "groot/gen/v1"
	"groot/internal/models/user"
	"groot/internal/response"

	"net/http"

	"github.com/gin-gonic/gin"
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

	SuccessFormat(c, *response.CodeSuccess, token)
}

func LoginCheck(username string, password string) (string, error) {
	return "token", nil
}

func SuccessFormat(c *gin.Context, tokens []token.Token, response response.Response, data interface{}) {
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
