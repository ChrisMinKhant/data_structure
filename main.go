package main

import (
	"ds/rnd/nestedlinkedlist"
)

var createdList nestedlinkedlist.LinkedList = nestedlinkedlist.NewLinkedList()

func main() {
	// temporaryNode := &nestedlinkedlist.Node{
	// 	OrderNumber: 25,
	// 	NextNode:    nil,
	// }

	// createdList.Add(*temporaryNode)

	// temporaryNode = &nestedlinkedlist.Node{
	// 	OrderNumber: 55,
	// 	NextNode:    nil,
	// }

	// createdList.Add(*temporaryNode)

	// temporaryNode = &nestedlinkedlist.Node{
	// 	OrderNumber: 66,
	// 	NextNode:    nil,
	// }

	// createdList.Add(*temporaryNode)

	createdList.ShowAll()
}
