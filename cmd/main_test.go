package main

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_getWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: "world is mine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWorld(); got != tt.want {
				t.Errorf("getWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getItems(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/sample")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	ctx := context.TODO()
	cnn, err := db.Conn(ctx)

	if err != nil {
		panic(err)
	}

	items := getItems(context.Background(), cnn)

	t.Log(items)
	if len(items) < 1 {
		t.Errorf("items,0")
	}
}
