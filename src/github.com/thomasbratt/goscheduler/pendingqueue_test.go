package goscheduler

import (
	"testing"
	"time"
)

func TestPendingQueueSmoke(t *testing.T) {
    var queue pendingQueue
    queue.Init()

    now := time.Now()

	queue.Add(&job{
            RunAt: now.Add(time.Second*4),
        })
    queue.Add(&job{
            RunAt: now.Add(time.Second*1),
        })
    queue.Add(&job{
            RunAt: now.Add(time.Second*3),
        })
    queue.Add(&job{
            RunAt: now.Add(time.Second*2),
        })
    queue.Add(&job{
            RunAt: now.Add(time.Second*5),
        })

    testNext(t, &queue, now.Add(time.Second * 1))
    testNext(t, &queue, now.Add(time.Second * 2))
    testNext(t, &queue, now.Add(time.Second * 3))
    testNext(t, &queue, now.Add(time.Second * 4))
    testNext(t, &queue, now.Add(time.Second * 5))
}

func testNext(t *testing.T, queue * pendingQueue, expected time.Time) {
    _, j := queue.Peek()
    queue.RemoveHead()

    if j.RunAt != expected {
        t.Logf("actual:%s, expected:%s", j.Interval, expected)
        t.Log(queue.String())
        t.Fatal("Pending queue not ordered")
    }
}
