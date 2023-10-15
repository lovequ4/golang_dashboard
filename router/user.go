package router

import (
	"Demo/controllers"

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

}
