package main

import (
	"fmt"
	"app/api"
	"app/psql"
	"github.com/gin-gonic/gin"
)

func main() {
	// debug or release mode
	gin.SetMode(gin.ReleaseMode)
	psql.GetDB()
	fmt.Println("The server is running!")
	routes.RegisterRoutes()
}