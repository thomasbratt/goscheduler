package goscheduler

import (
	"time"
)

// job contains information about a scheduled task.
type job struct {
	Action     func() bool
    Count      int
	Interval   time.Duration
    RunAt      time.Time
}

func (j *job) UpdateRunAt() {
	j.RunAt = time.Now().Add(j.Interval)
}
