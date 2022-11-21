package main

import (
	"fmt"
	"app/psql"
	"github.com/gin-gonic/gin"
	"app/api"
)

func main() {
	// debug or release mode
	gin.SetMode(gin.ReleaseMode)
	psql.ConnectDB()
	
	fmt.Println("The server is running!")
	routes.RegisterRoutes()
}