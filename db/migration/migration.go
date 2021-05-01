// package migration provides schema and default values for the db
package migration

import (
	"log"
	"time"

	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
	"github.com/OscarClemente/backend-tech-challenge-time/model/timers"
	"github.com/go-pg/pg/v10/orm"
)

// CreateSchema creates the timers table in the database
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

// PopulateWithTestData populates an empty database with test values
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

	now := time.Now()

	timer1 := timers.Timer{
		Title:       "Today's session",
		TimeElapsed: 65371,
		CreatedAt:   now,
	}

	timer2 := timers.Timer{
		Title:       "Yesterday's session",
		TimeElapsed: 12751,
		CreatedAt:   now.AddDate(0, 0, -1),
	}

	timer3 := timers.Timer{
		Title:       "May 1st session",
		TimeElapsed: 42753,
		CreatedAt:   time.Date(2021, 5, 1, 10, 0, 0, 0, time.UTC),
	}

	timer4 := timers.Timer{
		Title:       "March's session",
		TimeElapsed: 81275,
		CreatedAt:   time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC),
	}

	_, err = postgres.Db.Model(&timer1).Insert()
	_, err = postgres.Db.Model(&timer2).Insert()
	_, err = postgres.Db.Model(&timer3).Insert()
	_, err = postgres.Db.Model(&timer4).Insert()

	err = postgres.Db.Model(&timersInDb).Select()
	if err != nil {
		panic(err)
	}

	log.Println(timersInDb)

	return nil
}
