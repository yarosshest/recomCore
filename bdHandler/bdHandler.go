package bdHandler

import (
	"context"
	"database/sql"
	"log"
	"recomCore/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Close(client *ent.Client) {
	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func DbInit() {
	// Open new connection
	client := Open("postgresql://postgres:postgres@localhost/recomCore")
	defer Close(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	err := client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
