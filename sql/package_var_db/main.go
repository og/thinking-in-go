package main

import (
	globalDB "github.com/og/thinking-in-go/sql/package_var_db/db"
	"log"
	"net/http"
)
// 在程序运行前 运行中 关闭后 都通过 show processlist 查看数据库连接情况
func main () {
	// 在程序退出时 close
	defer globalDB.DB.Close()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		pingErr := globalDB.DB.Ping()
		var msg string
		if pingErr != nil { msg = "error" } else { msg = "ok" }
		_, err := writer.Write([]byte(msg)) ; if err != nil {
			panic(err)
		}
		// 可以使用多个 goroutine 执行很多次 db.Stats()，再通过 show processlist 查看连接池情况
		// for i:=0;i<100;i++ {
		// 	go func() {
		// 		globalDB.DB.Ping()
		// 	}()
		// }
	})
	addr := ":8421"
	log.Print("listen: http://127.0.0.1" + addr)
	err := http.ListenAndServe(addr, nil) ; if err != nil {
		panic(err)
	}
}
