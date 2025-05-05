package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(d int) *Node {
	return &Node{
		data: d,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
}

func (n *LinkedList) Append(d int) {
	newNode := NewNode(d)
	if n.head == nil {
		n.head = newNode
		return
	}

	cur := n.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newNode
}

func (n *LinkedList) PrePend(d int) {
	newNode := &Node{data: d, next: n.head}
	n.head = newNode
}

func (n *LinkedList) Delete(d int) {
	cur := n.head

	if cur == nil {
		return
	}

	if cur.data == d {
		cur = cur.next
	}

	for cur.next != nil {
		if cur.next.data == d {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (n *LinkedList) Display() {
	cur := n.head

	for cur != nil {
		fmt.Printf("%v, ", cur.data)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	root := LinkedList{}
	for _, d := range data {
		root.Append(d)
	}
	root.Display()

	root.PrePend(35)

	root.Display()

	root.Delete(5)

	root.Display()
}
