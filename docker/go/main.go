package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}
	http2.VerboseLogs = true
	http2.ConfigureServer(srv, nil)
	log.Fatal(srv.ListenAndServeTLS("/go/tls/tls.crt", "/go/tls/tls.key"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	/** URLパス表示 */
	w.Write([]byte("url path is " + r.URL.Path[1:] + "\n"))

	/** DB接続 */
	var dbConnectQuery string
	dbConnectQuery = "root:" + os.Getenv("GO_DB_PASSWORD") + "@tcp(" + os.Getenv("GO_DB_HOST") + ":3306)/ucwork"
	db, err := sql.Open("mysql", dbConnectQuery)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM user") //
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			w.Write([]byte(columns[i] + ": " + value + "\n"))
		}
		fmt.Println("-----------------------------------")
	}
}
