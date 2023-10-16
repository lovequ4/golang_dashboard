package router

import (
	"Demo/controllers"
	"Demo/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	product := router.Group("/products")

	product.POST("", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.ProductCreate(c)
	})

	product.GET("", middlewares.AuthMiddleware(), func(c *gin.Context) {
		controllers.AllProducts(c)
	})

	product.PUT(":id", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.ProductPut(c)
	})

	product.DELETE(":id", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.ProductDelete(c)
	})
}
