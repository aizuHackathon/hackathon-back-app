package db

import (
	"fmt"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

func GetSession() (*dbr.Session, error) {
	conn, err := dbr.Open("postgres", "host=postgres user=admin password=admin dbname=test sslmode=disable", nil)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}
	sess := conn.NewSession(nil)

	return sess, nil
}
