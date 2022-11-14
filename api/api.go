package routes

import (
	// "fmt"

	// "app/psql"
	// "database/sql"

	// "database/sql"

	
	// "github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(){
	c := gin.Default()
	
	c.GET("/", func(ctx *gin.Context)  {
		res := "Q_Q test Hello Q_Q"
		// psql.SelectUser(&sql.DB{})
		ctx.String(200, res)
	})

	v1 := c.Group("/user")
	{
		v1.GET("/", Aaaaa)
	}
	v2 := c.Group("/fruit")
	{
		v2.GET("/", 		allItems)
		// v2.GET("/order", 	)
		// v2.POST("/order", 	)
		// v2.PUT("/order", 	)
		// v2.DELETE("/order", )
	}
	c.Run(":3000")
}