package v1

import (
	genv1 "groot/gen/v1"
	"groot/internal/models/user"
	userrepo "groot/internal/repository/user"
	"groot/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *TaskStore) CreateUser(c *gin.Context) {
	var newUser genv1.CreateUser
	err := c.Bind(&newUser)
	if err != nil {
		ErrorFormat(c, response.CodeUserCreatePostDataFormatError.WithError(err))
		return
	}

	var user user.User
	user.Username = newUser.Username
	user.Password = newUser.Password

	result, err := userrepo.CreateUser(user)
	if err != nil {
		ErrorFormat(c, response.CodeUserCreateFailed.WithError(err))
		return
	}

	getUser(c, result, *response.CodeSuccess)
}

func getUser(c *gin.Context, users []user.User, response response.Response) {
	result := make([]genv1.User, 0, len(users))
	for _, value := range users {
		result = append(result, genv1.User{
			Id:       int64(value.ID),
			Username: value.Username,
		})
	}

	resp := genv1.GetUser{
		Code:      int64(response.Code),
		Message:   response.Message,
		Reference: response.Reference,
		Error:     response.Error,
		Data:      result,
	}
	c.JSON(http.StatusOK, resp)
}
