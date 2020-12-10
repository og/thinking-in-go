package tig_sql_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

var _ = `
^sql.DB^ 能与各种sql驱动结合，并且能管理**连接池**。

^import _ "github.com/go-sql-driver/mysql"^ 是加载 mysql 驱动，可以根据时间情况决定加载什么样的驱动。

一些其他驱动： github.com/go-sql-driver/mysql github.com/lib/pq

import 后需要在 sql.Open() 的第一个参数写上驱动名字，以使用此驱动连接数据库。
sql.Open() 的第二个参数是驱动对应的连接方式（用户名 密码 地址 库名等）。

注意：sql.DB 对象被设计为长期存在的，例如在 http 服务中，不要每次查询数据都使用 sql.Open 和使用完立即 db.Close()

sql.DB 维护了数据流的连接池，可以将 sql.DB 保存在包变量中，或者使用依赖注入管理 sql.DB (依赖注入:https://github.com/og/thinking-in-go/tree/master/di)

`
func TestDBPing(t *testing.T) {
	db, err := sql.Open("mysql", "root:somepass@tcp(127.0.0.1:3306)/thinking-in-go") ; if err != nil {
		panic(err)
	} else { defer db.Close() }
	err = db.Ping(); if err != nil {panic(err)}
}

func TestQuery(t *testing.T) {
	db, err := sql.Open("mysql", "root:somepass@tcp(127.0.0.1:3306)/thinking-in-go") ; if err != nil {
		panic(err)
	} ; defer db.Close()
	err = db.Ping(); if err != nil {panic(err)}
	rows, err := db.Query(`SELECT id, name FROM query`) ; if err != nil {
		panic(err)
	} else {
		defer rows.Close()
	}
	for rows.Next() {
		// rows.Close 之前会一直占用连接池中的连接，可以使用 time.Sleep 然后再执行中使用 show processlist 测试
		// time.Sleep(time.Second*4)
		var id int
		var name string
		err := rows.Scan(&id, &name) ; if err != nil {
			panic(err)
		}
		log.Print(id, " ", name)
	}
	err = rows.Err() ; if err != nil {
		panic(err)
	}
}