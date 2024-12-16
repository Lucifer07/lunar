package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService *userService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (u *UserHandler) GetAll(c *gin.Context) {
	users, err := u.userService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
