package routes

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"fmt"
)
type nameoutputs struct {
	Name 	string 		`json:"name"`
	Outputs []string 	`json:"outputs"`
}
func allItems(c *gin.Context){
	var res nameoutputs
	arr := [...]string{"apple", "banana", "watermelon"}
	
	res.Name = "All Fruits"
	res.Outputs = arr[:]

	fmt.Println("asdasda")
	c.JSON(200, res)
}
