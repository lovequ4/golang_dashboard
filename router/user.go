package router

import (
	"Demo/controllers"
	"Demo/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/users")

	user.POST("/signup", func(c *gin.Context) {
		controllers.SignUp(c)
	})

	user.POST("/signin", func(c *gin.Context) {
		controllers.SignIn(c)
	})

	user.GET("", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.AllUsers(c)
	})

	user.PUT(":id", middlewares.AuthMiddleware("admin"), func(c *gin.Context) {
		controllers.UserPut(c)
	})
}
