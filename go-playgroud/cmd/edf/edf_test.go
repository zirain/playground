package edf

import (
	"testing"
)

func TestEdfSchedule(t *testing.T) {
	// Some items and their priorities.
	A := &Entry{Value: "A", Weight: 3}
	B := &Entry{Value: "B", Weight: 2}
	C := &Entry{Value: "C", Weight: 1}
	D := &Entry{Value: "D", Weight: 2}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	e := NewEDF([]*Entry{A, B, C})
	t.Logf("=== add entry A, B, C\n")

	// Take the items out; they arrive in increasing deadline order.
	for i := 0; i < 12; i++ {
		item := e.Pick()
		t.Logf("%s ", item.Value)
	}

	// Add new entry of weight 2
	e.Add(D)
	t.Logf("=== add entry D\n")

	for i := 0; i < 8; i++ {
		item := e.Pick()
		t.Logf("%s ", item.Value)
	}

	// Delete entry B
	e.Delete(B)
	t.Logf("=== del entry B\n")

	for i := 0; i < 6; i++ {
		item := e.Pick()
		t.Logf("%s ", item.Value)
	}
}
