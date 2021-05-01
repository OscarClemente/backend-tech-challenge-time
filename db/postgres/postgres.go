// package postgres provides funtionality to connect to a postgres database using go-pg
package postgres

import (
	"github.com/go-pg/pg/v10"
)

// Db stores a database connection
var Db *pg.DB

// New creates and saves a database connection to the Db variable
func New(address string) {
	Db = pg.Connect(&pg.Options{
		Addr:     address,
		User:     "postgres",
		Password: "dbpass",
		Database: "timerdb",
	})
}
