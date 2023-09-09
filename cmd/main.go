package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/sample")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	cnn, err := db.Conn(ctx)

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/items", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%v", getItems(r.Context(), cnn))
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, getWorld())
	})

	http.HandleFunc("/404", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(404)
	})

	http.ListenAndServe(":8088", nil)
}

func getWorld() string {
	return "world is mine"
}

func getItems(ctx context.Context, cnn *sql.Conn) []string {
	r, err := cnn.QueryContext(ctx, "select * from items")
	defer r.Close()
	if err != nil {
		panic(err)
	}

	res := []string{}
	for r.Next() {
		var (
			id   int
			name string
		)
		err := r.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		res = append(res, fmt.Sprintf("{'id':'%d', 'name':'%s'}", id, name))
	}
	return res
}
