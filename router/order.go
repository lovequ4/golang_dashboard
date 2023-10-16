package router

import (
	"Demo/controllers"
	"Demo/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	order := router.Group("/orders")

	order.POST("", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.OrderCreate(c)
	})

	order.GET("", middlewares.AuthMiddleware(), func(c *gin.Context) {
		controllers.AllOrders(c)
	})

	order.PUT(":id", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.OrderPut(c)
	})

	order.DELETE(":id", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.OrderDelete(c)
	})
}
