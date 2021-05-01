package timers

import (
	"fmt"
	"time"

	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
)

type Timer struct {
	Id          int
	Title       string
	TimeElapsed int
	CreatedAt   time.Time `pg:"default:now()"`
}

func (timer Timer) String() string {
	return fmt.Sprintf("Timer<%d %s %d>", timer.Id, timer.Title, timer.TimeElapsed)
}

func (timer Timer) Save() (Timer, error) {
	_, err := postgres.Db.Model(&timer).Insert()
	if err != nil {
		panic(err)
	}

	return timer, err
}

func (timer Timer) Update() (Timer, error) {
	_, err := postgres.Db.Model(&timer).WherePK().Column("title", "time_elapsed").Update()
	if err != nil {
		panic(err)
	}

	err = postgres.Db.Model(&timer).WherePK().Select()
	if err != nil {
		panic(err)
	}

	return timer, err
}

func GetTimers() []Timer {
	var timers []Timer
	err := postgres.Db.Model(&timers).Select()
	if err != nil {
		panic(err)
	}

	return timers
}

func DeleteTimer(id int) error {
	_, err := postgres.Db.Exec("DELETE FROM timers WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	return err
}
