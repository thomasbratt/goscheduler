package goscheduler

import (
    "fmt"
    "testing"
    "time"
)

func TestRepeatEvery(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long running test.")
    }

    s := new(Scheduler)
    s.Init(1)

    start := time.Now()

    s.RepeatEvery(  time.Millisecond * 80,
            func() bool {
                time.Sleep(time.Millisecond * 20)
                fmt.Printf("smoke at: %s\n", time.Since(start))
                return true
            })

    time.Sleep(time.Second * 1)

    s.Close()
}

func TestRepeatEvery2ConcurrentJobs(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long running test.")
    }

    s := new(Scheduler)
    s.Init(2)

    start := time.Now()

    s.RepeatEvery(  time.Millisecond * 30,
            func() bool {
                time.Sleep(time.Millisecond * 20)
                fmt.Printf("action1 at: %s\n", time.Since(start))
                return true
            })
    s.RepeatEvery(  time.Millisecond * 80,
            func() bool {
                time.Sleep(time.Millisecond * 20)
                fmt.Printf("action2 at: %s\n", time.Since(start))
                return true
            })

    time.Sleep(time.Second * 1)

    s.Close()
}

func TestRepeatEvery2JobsWithConcurrency1(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long running test.")
    }

    s := new(Scheduler)
    s.Init(1)

    start := time.Now()

    s.RepeatEvery(  time.Millisecond * 100,
            func() bool {
                time.Sleep(time.Millisecond * 100)
                fmt.Printf("action 100 at: %s\n", time.Since(start))
                return true
            })
    s.RepeatEvery(  time.Millisecond * 10,
            func() bool {
                time.Sleep(time.Millisecond * 10)
                fmt.Printf("action 10 at: %s\n", time.Since(start))
                return true
            })

    time.Sleep(time.Second * 1)

    s.Close()
}

func TestRepeatForCount(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping long running test.")
    }

    s := new(Scheduler)
    s.Init(1)

    start := time.Now()

    count := 0
    s.RepeatForCount(   2,
                        time.Millisecond * 10,
                        func() bool {
                            count++
                            time.Sleep(time.Millisecond * 50)
                            fmt.Printf("action at: %s\n", time.Since(start))
                            return true
                        })

    time.Sleep(time.Millisecond * 500)

    s.Close()

    if count != 2 {
        t.Fatal("Failed to execute required number of times")
    }
}
