package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l LinkedList) printHead() {
	fmt.Println(l.Head)
}

func (l LinkedList) printAll() {
	fmt.Println("LINKEDLIST:")
	pointer := l.Head

	for pointer != nil {
		fmt.Println(pointer)
		pointer = pointer.Next
	}
}

func (l LinkedList) append(nuevo interface{}) {

	nuevoNodo := Node{nuevo, nil}

	pointer := l.Head

	for pointer.Next != nil {
		pointer = pointer.Next
	}
	pointer.Next = &nuevoNodo
}

func (l *LinkedList) deleteFirst(value interface{}) {
	//pointer := l.Head

	if l.Head.Value == value {
		//delete head
		l.Head = l.Head.Next
	} else if l.Head.Next != nil {
		pointer := l.Head.Next
		before := l.Head

		for pointer.Next != nil {
			if pointer.Value == value {
				//saltarse este
				before.Next = pointer.Next
				return
			}
			pointer = pointer.Next
			before = before.Next
		}
		//si el ultimo es el valor
		if pointer.Value == value {
			before.Next = nil
		}

	}
}

func main() {
	first := Node{"primero", nil}
	linkedList := LinkedList{&first}
	//linkedList.append(22)
	//linkedList.append("trauma")
	//linkedList.append("ultimo")

	//delete
	linkedList.deleteFirst("primero")
	//linkedList.deleteFirst("trauma")

	linkedList.printAll()

}
