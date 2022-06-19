package main

import (
	"fmt"
)

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func (sl *mySingleLinkedList) checkValidate(node node) error {
	current := sl.head
	if current.data == node.data {
		return fmt.Errorf("%v sudah ada, node tidak boleh sama", current.data)
	}
	for current.next != nil {
		if current.next.data == node.data {
			return fmt.Errorf("%v sudah ada, node tidak boleh sama", current.next.data)
		}
		current = current.next
	}
	return nil
}

func (sl *mySingleLinkedList) addToHead(name string) {
	newNode := &node{
		data: name,
	}

	if sl.head == nil {
		sl.head = newNode
	} else {
		err := sl.checkValidate(*newNode)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		newNode.next = sl.head
		sl.head = newNode
	}

	sl.size++
}

func (sl *mySingleLinkedList) addToTail(name string) {
	newNode := &node{
		data: name,
	}

	if sl.head == nil {
		sl.head = newNode
	} else {
		err := sl.checkValidate(*newNode)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	sl.size++
}

func (sl *mySingleLinkedList) iterateList() {
	fmt.Println("\nLinked List")
	for node := sl.head; node != nil; node = node.next {
		fmt.Print(node.data, " -> ")
	}
	fmt.Println()
}

func NewLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	//Create New Linked List
	fmt.Println("Create New Linked List")
	singleList := NewLinkedList()
	singleList.addToHead("28")
	singleList.addToHead("01")
	singleList.addToHead("01")
	singleList.addToTail("23")
	singleList.addToTail("23")
	singleList.addToHead("08")
	singleList.addToTail("08")
	singleList.addToTail("11")
	singleList.iterateList()
	fmt.Println("Size: ", singleList.size)

}
