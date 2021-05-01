package postgres

import (
	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func New(address string) {
	Db = pg.Connect(&pg.Options{
		Addr:     address,
		User:     "postgres",
		Password: "dbpass",
		Database: "timerdb",
	})
}
