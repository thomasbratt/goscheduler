package goscheduler

import (
	"container/heap"
    "errors"	
	"strings"
)

// pendingQueue is a collection of jobs that are scheduled to run at some
// point in the future.
type pendingQueue []*job

func (pq * pendingQueue) Init(){
	heap.Init(pq)
}

func (pq * pendingQueue) Add(j *job){
    heap.Push(pq, j)
    heap.Fix(pq, len(*pq)-1)
}

func (pq * pendingQueue) RemoveHead() {
	heap.Pop(pq)
}

func (pq * pendingQueue) Peek() (error, *job) {
    if len(*pq) > 0 {
        return nil, (*pq)[0]
    } else {
        err := errors.New("No items in queue.")
        return err, nil
    }
}

func (pq * pendingQueue) String() string {
	results := make([]string, len(*pq))

	for i, j := range *pq {
		results[i] = j.Interval.String()
	}

	return strings.Join(results, "|")
}

// Heap.interface implementation.
func (pq *pendingQueue) Push(j interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*pq = append(*pq, j.(*job))
}

// Heap.interface implementation.
func (pq *pendingQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	j := old[n-1]
	*pq = old[0 : n-1]
	return j
}

// sort.Interface implementation.
func (pq pendingQueue) Len() int {
	return len(pq)
}

// sort.Interface implementation.
func (pq pendingQueue) Less(i, j int) bool {
	return pq[i].RunAt.Before(pq[j].RunAt)
}

// sort.Interface implementation.
func (pq pendingQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
