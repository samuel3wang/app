package routes

import (
	// "fmt"
	// "app/psql"
	// "database/sql"
	
	"github.com/gin-gonic/gin"
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
	// 	v1.GET("/", Lookup)
	 	v1.GET("/", add)
	// 	v1.PUT("/", )
	// 	v1.DELETE("/:", )

	}
	v2 := c.Group("/fruit")
	{
		v2.GET("/", allItems)
	}
	c.Run(":3000")
}