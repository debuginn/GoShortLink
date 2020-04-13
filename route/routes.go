package routes

import (
	"GoShortLink/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/auth/register", controller.Register)

	return r
}
