// Package timers provides the internal timer struct that is used by the DB
// functions to create, query, update and delete timers are also provided
package timers

import (
	"errors"
	"fmt"
	"time"

	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
)

// Timer struct represents a timer row inside the DB
type Timer struct {
	Id          int
	Title       string
	TimeElapsed int
	CreatedAt   time.Time `pg:"default:now()"`
}

// String returns a string that describes the timer, mainly for debugging purposes
func (timer Timer) String() string {
	return fmt.Sprintf("Timer<%d %s %d>", timer.Id, timer.Title, timer.TimeElapsed)
}

// Save stores a new timer into the DB
func (timer *Timer) Save() error {
	_, err := postgres.Db.Model(timer).Insert()
	if err != nil {
		panic(err)
	}

	return err
}

// Update updates the title and timer_elapsed for an existing timer
func (timer *Timer) Update() error {
	_, err := postgres.Db.Model(timer).WherePK().Column("title", "time_elapsed").Update()
	if err != nil {
		panic(err)
	}

	err = postgres.Db.Model(timer).WherePK().Select()
	if err != nil {
		panic(err)
	}

	return err
}

// GetTimers returns all the timers stored in the DB
func GetTimers() []Timer {
	var timers []Timer
	err := postgres.Db.Model(&timers).Select()
	if err != nil {
		panic(err)
	}

	return timers
}

// DeleteTimer deletes the timer that has the provided ID
func DeleteTimer(id int) error {
	res, err := postgres.Db.Exec("DELETE FROM timers WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	if res.RowsAffected() != 1 {
		err = errors.New("Could not find a timer to delete.")
	}

	return err
}
