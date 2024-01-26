package nestedlinkedlist

import "fmt"

/*
*	This is node struct which is the basic component of the linked list
 */
type Node struct {
	OrderNumber int
	NextNode    *Node
}

/*
*	This is the main struct for linked list.
 */
type LinkedList struct {
	HeadNode *Node
}

// Constructor function
func NewLinkedList() LinkedList {
	return LinkedList{
		HeadNode: &Node{
			OrderNumber: 0,
			NextNode:    nil,
		},
	}
}

// Add function to add new node next to the last node
func (linkedList *LinkedList) Add(node Node) {

	if linkedList.HeadNode.NextNode != nil {
		tempNode := linkedList.HeadNode.NextNode

		for tempNode.NextNode != nil {
			tempNode = tempNode.NextNode
		}

		tempNode.NextNode = &node

		return
	}

	linkedList.HeadNode.NextNode = &node
}

func (linkedList *LinkedList) ShowAll() {

	if linkedList.HeadNode.NextNode != nil {
		tempNode := linkedList.HeadNode.NextNode

		for tempNode.NextNode != nil {
			fmt.Printf(" OrderNumber : %v \n", tempNode.OrderNumber)

			tempNode = tempNode.NextNode
		}

		fmt.Printf(" OrderNumber : %v \n", tempNode.OrderNumber)

		return
	}

	fmt.Printf(" There is only head node. ")
}
