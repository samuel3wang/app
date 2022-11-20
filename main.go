package main

import (
	"fmt"
	"app/psql"
	"github.com/gin-gonic/gin"
	"app/api"
)
// const(
// 	host = "localhost"
// 	port = 5432
// 	user = "samuel"
// 	password = "3336"
// 	dbname = "samuel"
// )

func main() {
	// debug or release mode
	gin.SetMode(gin.ReleaseMode)
	psql.GetDB()
	

	// // connection string.
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// // Open DB
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// // Close DB
	// defer db.Close()

	// routes.Add(db, err)
	
	fmt.Println("The server is running!")
	routes.RegisterRoutes()
}