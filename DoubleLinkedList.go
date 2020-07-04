package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (L *LinkedList) insert(inData int) {
	fmt.Println("Inserted: ", inData)
	newNode := &Node{Value: inData}
	if L.Head == nil {
		L.Head = newNode
	} else {
		currentNode := L.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		newNode.Prev = currentNode
		currentNode.Next = newNode
		L.Tail = newNode
	}
}

func (L *LinkedList) PrintList() {
	currentNode := L.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (L *LinkedList) reverse() {
	fmt.Println("Reversed-")
	currentNode := L.Head
	tail := L.Tail
	L.Tail = L.Head
	for currentNode != nil {
		temp := currentNode.Next
		currentNode.Next = currentNode.Prev
		currentNode.Prev = temp
		currentNode = temp
	}
	L.Head = tail
}

func (list *LinkedList)bubblesort() {
	fmt.Println("Sorted-")
	swapped := true
	first := list.Head
	next := &Node{}
	if first == nil {
		return
	}
	for swapped {
		swapped = false
		first = list.Head
		for first.Next != next && first.Next != nil {

			if first.Value > first.Next.Value {
				swap(first, first.Next)
				swapped = true
			}
			first = first.Next
		}
		next = first
	}
}

func swap(a *Node, b *Node) {
	temp := a.Value
	a.Value = b.Value
	b.Value = temp
}

func newList() *LinkedList {
	return &LinkedList{}
}

func main() {
	list := newList()
	list.insert(5)
	list.insert(7)
	list.insert(2)
	list.insert(4)
	list.insert(13)
	list.insert(9)
	list.insert(3)
	list.reverse()
	list.PrintList()
	list.bubblesort()
	list.PrintList()
}
