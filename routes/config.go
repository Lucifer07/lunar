package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes(engine *gin.Engine) *routes {
	return &routes{
		router: engine,
	}
}
