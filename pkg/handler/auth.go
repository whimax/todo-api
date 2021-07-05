package handler

import (
	"github.com/gin-gonic/gin"
	todo "todo-app/pkg"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {

	}
}

func (h *Handler) signIn(c *gin.Context) {

}
