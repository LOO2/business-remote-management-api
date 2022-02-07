package delivery

import (
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type Config struct {
	R *gin.Engine
}
