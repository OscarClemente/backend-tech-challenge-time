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

func (t Timer) String() string {
	return fmt.Sprintf("Timer<%d %s %d>", t.Id, t.Title, t.TimeElapsed)
}

func GetTimers() []Timer {
	var timers []Timer
	err := postgres.Db.Model(&timers).Select()
	if err != nil {
		panic(err)
	}

	return timers
}
