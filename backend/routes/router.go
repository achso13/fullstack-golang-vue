package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	router.POST("/api/login", controllers.Login)
	router.POST("/api/register", controllers.Register)

	return router
}
