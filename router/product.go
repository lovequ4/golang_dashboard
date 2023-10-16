package router

import (
	"Demo/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	product := router.Group("/products")

	product.POST("", func(c *gin.Context) {
		controllers.ProductCreate(c)
	})

	product.GET("", func(c *gin.Context) {
		controllers.AllProducts(c)
	})

	product.PUT(":id", func(c *gin.Context) {
		controllers.ProductPut(c)
	})

	product.DELETE(":id", func(c *gin.Context) {
		controllers.ProductDelete(c)
	})
}
