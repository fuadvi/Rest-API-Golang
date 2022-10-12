package routers

import (
	"assement_02/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()
	router.POST("/items", controllers.CreateItems)
	router.POST("/order", controllers.CreateOrder)

	return router
}
