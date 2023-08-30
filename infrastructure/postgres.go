package infrastructure

import (
	"database/sql"
	"fmt"
	"ginweb/ent"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*ent.Client, error) {
	// user := "postgres"
	// password := "sa"
	// host := "host.docker.internal"
	// port := "5432"
	// dbname := "template"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		err = fmt.Errorf("[infrastructure.Open] failed to open connection to database: %w", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), err
}
