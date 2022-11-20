package routes

import (
	//"app/psql"
	//"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Lookup(c *gin.Context){
	c.String(http.StatusOK, "hello")
	fmt.Println("hello")
}

func add(c *gin.Context) {
	
	insert := `INSERT INTO "userinfo"("uid", "username", "department") values($1, $2, $3)`
	_, err = db.Exec(insert, 10, "elsa9", "qq9")
	if err != nil {
		panic(err)
	}
	return
}