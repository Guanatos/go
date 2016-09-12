package main

import ( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql","jihanki:vendo123@tcp(127.0.0.1:3306)/jihanki?charset=utf8")
	fmt.Println("db:",db)
	fmt.Println("err:",err)
}
