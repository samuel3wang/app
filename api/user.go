package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
func Aaaaa(c *gin.Context){
	c.String(http.StatusOK, "hello")
	fmt.Println("hello")
}