package main

import (
	"Demo/database"
	_ "Demo/docs"
	"Demo/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() { //main() 函數之前被調用
	database.DBconnect()
}

// @title  MyGolang
// @version 1.0
// @description  Swagger API.
// @host localhost:8080
func main() {
	r := gin.Default()

	router.UserRoutes(r) // 引入user router
	router.ProductRoutes(r)
	router.OrderRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Run(":8080")
}
