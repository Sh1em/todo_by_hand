package handler

import (
	todo "todo_app"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Fatalf("error sign-up : %s", err.Error())
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
