package psql

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"	
)
const(
	host = "localhost"
	port = 5432
	user = "samuel"
	password = "3336"
	dbname = "samuel"
)

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func GetDB() {

	// connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// Close DB
	defer db.Close()
	
	// c := gin.Default()

	// ---Insert--- 
	// insert := `INSERT INTO "userinfo"("uid", "username", "department") values($1, $2, $3)`
	// _, err = db.Exec(insert, 5,"elsa","qq")
	// checkError(err)
	
	// ---Update---
	// update := `UPDATE "userinfo" SET "username"=$1, "department"=$2 WHERE "uid"=$3`
	// _, err = db.Exec(update, "julia", "music", 21)
	// checkError(err)

	// ---Delete---
	// delete := `DELETE FROM "userinfo" WHERE "uid"=$1`
	// _, err = db.Exec(delete, 4)
	// checkError(err)

	// ---Queries---
	// rows, err := db.Query(`SELECT * FROM "userinfo"`)
	// checkError(err)

	// defer rows.Close()
	// for rows.Next(){
	// 	var uid int64
	// 	var username string
	// 	var department string
	// 	var created interface{}
		
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkError(err)

	// 	fmt.Println(uid, username, department, &created)
	// }

// routes -> postgres -> api 
	//---Retrieve---
	// user := new(User)
	
	// rows := db.QueryRow(`SELECT * FROM "userinfo" where uid = 2`)
	
	// if err := rows.Scan(&user.Id, &user.Username, &user.Department); err != nil {
	// 	fmt.Println("Failed")
	// 	return
	// }
	// fmt.Println("success", *user)
}