package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

func GetSession() *dbr.Session {
	conn, err := dbr.Open("postgres", "host=postgres user=admin password=admin dbname=test sslmode=disable", nil)
	if err != nil {
		return nil
	}
	sess := conn.NewSession(nil)

	return sess
}
