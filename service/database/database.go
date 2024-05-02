package database

import (
	"clashapiv2-api/environment"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect()(*sql.DB, error) {
	env := environment.LoadEnvironment()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.DbHost, env.DbPort, env.DbUser, env.DbPass, env.DbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}