package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	//	"math"
	"net/http"
	"path"
	"strconv"
)

func main() {
	fib_server()
}

func find_or_create_fib_index(index int) (fib_value int) {
	db := db_connection()
	var value int
	err := db.QueryRow("SELECT value FROM fibonacci WHERE index=$1", index).Scan(&value)
	switch {
	case err == sql.ErrNoRows:
		switch {
		case index <= 1:
			db.QueryRow("insert into fibonacci values($1, $2) returning value", index, index).Scan(&value)
		default:
			value = find_or_create_fib_index(index-1) + find_or_create_fib_index(index-2)
			db.Exec("insert into fibonacci values($1, $2)", index, value)
		}
	case err != nil:
		log.Fatal(err)
	}
	db.Close()
	return value
}

func fib_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url_path, fib_max := path.Split(r.URL.Path[1:])
		n, err := strconv.Atoi(fib_max)
		if url_path != "fibonacci/" {
			panic("format url as fibonacci/:n")
		}
		if err != nil {
			panic(err)
		}
		fib_sequence := make([]int, n)
		for i := 0; i < n; i++ {
			fib_sequence[i] = find_or_create_fib_index(i)
		}
		fmt.Fprintln(w, fib_sequence)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func db_connection() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=weekly_workshop dbname=weekly_workshop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
