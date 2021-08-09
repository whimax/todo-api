package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
