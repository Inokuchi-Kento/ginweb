package infrastructure

import (
	"database/sql"
	"fmt"
	"ginweb/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*ent.Client, error) {
	user := "postgres"
	password := "sa"
	host := "localhost"
	port := "5555"
	dbname := "template"
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		err = fmt.Errorf("[infrastructure.Open] failed to open connection to database: %w", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), err
}
