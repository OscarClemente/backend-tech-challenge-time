package timers

import (
	"fmt"
	"time"
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
