package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// HasCycle reports whether the list loops back on itself, using Floyd's
// tortoise-and-hare: a fast pointer laps a slow one inside any cycle.
func HasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	a.Next, b.Next, c.Next = b, c, b // 1 -> 2 -> 3 -> back to 2
	fmt.Println(HasCycle(a))
}
