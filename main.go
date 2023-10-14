package main

import (
	"Demo/database"

	"github.com/gin-gonic/gin"
)

func init() { //main() 函數之前被調用
	database.DBconnect()
}

func main() {
	r := gin.Default()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
