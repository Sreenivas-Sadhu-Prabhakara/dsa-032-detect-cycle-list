package main

import "testing"

func list(vals ...int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func TestHasCycle(t *testing.T) {
	if HasCycle(nil) {
		t.Error("nil list has no cycle")
	}
	if HasCycle(list(1, 2, 3)) {
		t.Error("straight list has no cycle")
	}

	// Cycle back to the middle.
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	a.Next, b.Next, c.Next = b, c, b
	if !HasCycle(a) {
		t.Error("expected cycle (tail -> middle)")
	}

	// Single node pointing at itself.
	s := &Node{Val: 7}
	s.Next = s
	if !HasCycle(s) {
		t.Error("expected cycle (self-loop)")
	}

	if HasCycle(list(1)) {
		t.Error("single node without loop has no cycle")
	}
}
