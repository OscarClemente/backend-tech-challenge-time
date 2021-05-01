package migration

import (
	"log"
	"time"

	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
	"github.com/OscarClemente/backend-tech-challenge-time/model/timers"
	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema() error {
	models := []interface{}{
		(*timers.Timer)(nil),
	}

	for _, model := range models {
		err := postgres.Db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        false,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func PopulateWithTestData() error {
	var timersInDb []timers.Timer
	err := postgres.Db.Model(&timersInDb).Select()
	if err != nil {
		panic(err)
	}

	if len(timersInDb) > 0 {
		// Already populated
		log.Println("DB Already populated")
		log.Println(timersInDb)
		return nil
	}

	log.Println("Adding timers")

	timer1 := timers.Timer{
		Title:       "Test timer 1",
		TimeElapsed: 65371,
		CreatedAt:   time.Now(),
	}

	timer2 := timers.Timer{
		Title:       "Test timer 2",
		TimeElapsed: 1275,
		CreatedAt:   time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	_, err = postgres.Db.Model(&timer1).Insert()
	_, err = postgres.Db.Model(&timer2).Insert()

	err = postgres.Db.Model(&timersInDb).Select()
	if err != nil {
		panic(err)
	}

	log.Println(timersInDb)

	return nil
}
