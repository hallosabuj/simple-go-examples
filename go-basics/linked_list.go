package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func (list *LinkedList) AddItemInTheBegining(val int) *LinkedList {
	item := LinkedList{
		Val:  val,
		Next: list,
	}
	item.Next = list
	return &item
}
func (list *LinkedList) AddItemAtTheEnd(val int) *LinkedList {
	if list == nil {
		fmt.Println("List is empty")
		list = &LinkedList{
			Val:  val,
			Next: nil,
		}
		return list
	} else {
		fmt.Println("List is not empty")
		item := list
		for item.Next != nil {
			item = item.Next
		}
		item.Next = &LinkedList{
			Val:  val,
			Next: nil,
		}
		return list
	}
}
func (list *LinkedList) Traverse() {
	item := list
	for item.Next != nil {
		fmt.Println("Value:=", item.Val)
		item = item.Next
	}
	if item.Next == nil {
		fmt.Println("Value:=", item.Val)
	}
}
func (list *LinkedList) IsPresent(value int) (*LinkedList, bool) {
	item := list
	for item != nil {
		if item.Val == value {
			return item, true
		}
		item = item.Next
	}
	return nil, false
}

func main() {
	fmt.Println("Hello world")
	var head *LinkedList = nil
	head = head.AddItemAtTheEnd(1)
	head = head.AddItemAtTheEnd(2)
	head = head.AddItemAtTheEnd(3)
	head = head.AddItemInTheBegining(0)
	head.Traverse()
	node, ok := head.IsPresent(1)
	if ok {
		fmt.Println("Item is present : ", node.Val)
	} else {
		fmt.Println("Item is not present")
	}
}
