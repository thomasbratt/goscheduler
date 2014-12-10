// Package goscheduler provides a task scheduler with configurable concurrency
// limits.
package goscheduler

import (
	"time"
)

type Scheduler struct {
    newItems chan *job
}

// Init initializes the task scheduler with a specified number of concurrent
// workers.
func (s *Scheduler) Init(concurrency int) {
    s.newItems = make(chan *job)
    go s.dispatcher(concurrency, s.newItems)
}

// RepeatEvery adds a task to the scheduler that is invoked every interval.
//
//    interval time.Duration
//      - The interval between the previous execution finishing and the next one starting.
//
//    action func() bool
//      - The action to execute when the job is ready to run. Return false to stop future invocations.
func (s *Scheduler) RepeatEvery(   
        interval time.Duration,
        action func() bool) {
	j := &job{ Action:      action,
               Count:       -1,
		       Interval:    interval }
    j.UpdateRunAt()
    s.newItems <- j
}

// RepeatForCount adds a task to the scheduler that is invoked at intervals for
// a fixed number of times.
//
//    count int
//      - The number of times the job should be run.
// 
//    interval time.Duration
//      - The interval between the previous execution finishing and the next one starting.
//
//    action func() bool
//      - The action to execute when the job is ready to run. Return false to stop future invocations.
func (s *Scheduler) RepeatForCount( 
        count int,
        interval time.Duration,
        action func() bool) {
    j := &job{ Action:      action,
               Count:       count,
               Interval:    interval }
    j.UpdateRunAt()
    s.newItems <- j
}

// Close stops the scheduler from running.
// Currently running jobs will not be cancelled by this call but no new jobs
// will be started.
func (s *Scheduler) Close() {
    close(s.newItems)
}

func (s *Scheduler) dispatcher(concurrency int, newItems chan *job){
    var pendingJobs pendingQueue
    pendingJobs.Init()

    doneJobs := make(chan *job)
    readyJobs := make(chan *job, concurrency)

    for i := 0; i<concurrency; i++ {
        go s.worker(readyJobs, doneJobs)
    }

    next := time.Hour

    for {
        var active *job
        select {
            case j, ok := <- newItems:
                if ok {
                    pendingJobs.Add(j)
                    active, next = s.reschedule(&pendingJobs)
                } else {
                    close(readyJobs)
                    return
                }

            case j := <- doneJobs:
                if j.Count != 0 {
                    j.UpdateRunAt()
                    pendingJobs.Add(j)    
                }
                active, next = s.reschedule(&pendingJobs)

            case <- time.After(next):
                active, next = s.reschedule(&pendingJobs)                
        }

        if active != nil {
            readyJobs <- active
        }
    }
}

func (s *Scheduler) reschedule(pendingJobs * pendingQueue) (*job, time.Duration) {
    now := time.Now()

    if error, j := pendingJobs.Peek(); error == nil {
        if j.RunAt.After(now) {
            return nil, j.RunAt.Sub(now)
        } else {
            pendingJobs.RemoveHead()
            return j, 0
        }
    } else {
        return nil, time.Hour
    }
}

func (s *Scheduler) worker(readyJobs chan *job, doneJobs chan *job){
    for j := range readyJobs {
        shouldContinue := j.Action()

        // Set repeat count to zero for jobs that want to stop.
        if ! shouldContinue {
            j.Count = 0
        }

        // Deccrement repeat count for 'RepeatForCount' jobs.
        if j.Count > 0 {
            j.Count--
        }

        doneJobs <- j
    }
}
