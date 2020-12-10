package globalDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func init () {
	// 不要使用 db, err := sql.Open() 这样会导致申明的 db 是 init 函数的局部变量，而不是赋值 db
	var err error
	DB, err = sql.Open("mysql", "root:somepass@tcp(127.0.0.1:3306)/thinking-in-go") ; if err != nil {
		panic(err)
	}
	// 这里不要 Close() 因为当 init 执行完毕就 Close 会导致 DB 不可用
	//	defer DB.Close()
	err = DB.Ping(); if err != nil {panic(err)}
}
