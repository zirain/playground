package edf

import (
	"container/heap"
	"math/rand"
	"time"
)

// Entry is an item for load balance
type Entry struct {
	deadline float64
	index    int64
	Value    string
	Weight   float64
}

// priorityQueue is a queue that always pop the highest priority item
type priorityQueue []*Entry

// Len implements heap.Interface/sort.Interface
func (pq priorityQueue) Len() int { return len(pq) }

// Less implements heap.Interface/sort.Interface
func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].deadline == pq[j].deadline {
		return pq[i].index < pq[j].index
	}
	return pq[i].deadline < pq[j].deadline
}

// Swap implements heap.Interface/sort.Interface
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push implements heap.Interface for pushing an item into the heap
func (pq *priorityQueue) Push(x interface{}) {
	entry := x.(*Entry)
	*pq = append(*pq, entry)
}

// Pop implements heap.Interface for poping an item from the heap
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	*pq = old[0 : n-1]
	return entry
}

// EDF implements the Earliest Deadline First scheduling algorithm
type EDF struct {
	pq       *priorityQueue
	curIndex int64
	curDDL   float64
}

// Add a new entry for load balance
func (e *EDF) Add(entry *Entry) {
	entry.deadline = e.curDDL + 1/entry.Weight
	e.curIndex++
	entry.index = e.curIndex
	heap.Push(e.pq, entry)
}

// AddRaw add a new entry for load balance without sort
func (e *EDF) AddRaw(entry *Entry) {
	entry.deadline = e.curDDL + 1/entry.Weight
	e.curIndex++
	entry.index = e.curIndex
	*e.pq = append(*e.pq, entry)
}

// Delete an entry
func (e *EDF) Delete(entry *Entry) {
	entry.Weight = -1
}

// Pick an available entry
func (e *EDF) Pick() *Entry {
	// if no available entry, return nil
	if len(*e.pq) <= 0 {
		return nil
	}
	entry := heap.Pop(e.pq).(*Entry)
	if entry.Weight <= 0 {
		// if Weight isn't positive, try another entry
		return e.Pick()
	}
	// curDDL should be entry's deadline so that new added entry would have a fair
	// competition environment with the old ones
	e.curDDL = entry.deadline
	entry.deadline = entry.deadline + 1/entry.Weight
	e.curIndex++
	entry.index = e.curIndex
	heap.Push(e.pq, entry)
	return entry
}

// NewEDF create a new edf scheduler
func NewEDF(entries []*Entry) *EDF {
	// make a new edf
	priorityQueue := make(priorityQueue, 0)
	edf := &EDF{
		pq:       &priorityQueue,
		curIndex: 0,
	}

	// put entries into priority queue
	// TODO(maziang): use O(N) heap.Init instead of O(NlogN) Add.
	for _, entry := range entries {
		edf.AddRaw(entry)
	}
	heap.Init(edf.pq)

	// avoid instance flood pressure for the first entry
	// start from a random one via pick random times
	rand.Seed(time.Now().UnixNano())
	randomPick := rand.Intn(len(entries))
	for i := 0; i < randomPick; i++ {
		edf.Pick()
	}
	return edf
}
