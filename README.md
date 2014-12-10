goscheduler
===========

An efficient periodic task scheduler with concurrency limits.

Features
--------

* Limits the nunmber of concurrently executing tasks. This avoids overloading
  the resource used by the task (a server or network link, for example).
* Efficiently schedules tasks that must be repeated periodically.
* Individual tasks are rescheduled _after_ they have run. This prevents multiple
  invocations of the same task from overlapping.
  
Installation
--------

$ go install github.com/thomasbratt/goscheduler

Examples
--------

See the unit test file: scheduler_test.go
