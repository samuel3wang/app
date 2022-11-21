package routes

import (
	"app/psql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func Lookup(c *gin.Context){
	c.String(http.StatusOK, "hello")
	fmt.Println("hello")
}

func add(c *gin.Context) {
	
	insert := `INSERT INTO "userinfo"("uid", "username", "department") values($1, $2, $3)`
	_, err := psql.DB.Exec(insert, 10, "elsa10", "qq10")
	checkError(err)
	
	return
}