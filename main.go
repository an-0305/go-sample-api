package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sample-api/ent"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world!")

}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://postgres:test@127.0.0.1")

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
