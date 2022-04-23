goscheduler
===========

An efficient periodic task scheduler with concurrency limits.

[![CircleCI](https://circleci.com/gh/thomasbratt/goscheduler/tree/master.svg?style=svg)](https://circleci.com/gh/thomasbratt/goscheduler/tree/master)

Features
--------

* Limits the number of concurrently executing tasks. This avoids overloading
  the resource used by the task (a server or network link, for example).
* Efficiently schedules tasks that must be repeated periodically.
* Individual tasks are rescheduled _after_ they have run. This prevents multiple
  invocations of the same task from overlapping.
  
Installation
--------

$ go install github.com/thomasbratt/goscheduler

Example
--------

```
  package main
  
  import (
      "fmt"
      "time"
      "github.com/thomasbratt/goscheduler"
  )
  
  func main(){
      fmt.Println("Starting...")
      
      s := new(goscheduler.Scheduler)
      s.Init(1)
      
      start := time.Now()
      
      s.RepeatEvery(  time.Millisecond * 100,
                      func() bool {
                          fmt.Printf("Scheduled task ran at: %s\n", time.Since(start))
                          return true
                      })
      
      time.Sleep(time.Second * 1)
      
      s.Close()
```

Output from Example
--------

```
>  Starting...
>  Scheduled task ran at: 100.245068ms
>  Scheduled task ran at: 200.570528ms
>  Scheduled task ran at: 300.895522ms
>  Scheduled task ran at: 401.277121ms
>  Scheduled task ran at: 501.588894ms
>  Scheduled task ran at: 601.925013ms
>  Scheduled task ran at: 702.263688ms
>  Scheduled task ran at: 802.606226ms
>  Scheduled task ran at: 902.941508ms
```
