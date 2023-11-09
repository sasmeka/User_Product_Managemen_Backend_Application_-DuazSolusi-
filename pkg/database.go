package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Postgres_Database() (*sqlx.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", host, user, password, dbName, port)

	return sqlx.Connect("postgres", config)

}
