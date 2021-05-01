package postgres

import (
	"github.com/go-pg/pg/v10"
)

var Db *pg.DB

func New() {
	Db = pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5431",
		User:     "postgres",
		Password: "dbpass",
		Database: "timerdb",
	})
}
