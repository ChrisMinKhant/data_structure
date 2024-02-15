package nestedlinkedlist

import (
	"fmt"
)

type LinkedList struct {
	headNode *Node
}

/* Declare the linked list with initiated
* head node and return the memory address of it.
 */
func NewLinkedList() *LinkedList {
	return &LinkedList{
		headNode: &Node{
			data:     0,
			nextNode: nil,
		},
	}
}

/*
*	Add new node to the linked list
*	Time complexity of O(n)
 */
func (linkedList *LinkedList) Add(data *int) {

	// Check is there any next node after the head node
	if linkedList.headNode.nextNode != nil {
		tempNode := linkedList.headNode.nextNode

		// Iterate through till the next node is nil
		for tempNode.nextNode != nil {
			tempNode = tempNode.nextNode
		}

		// Iteration reaches the end of the list.
		// Therefore, add new node to the next node.
		tempNode.nextNode = &Node{
			data:     *data,
			nextNode: nil,
		}

		return
	}

	// There is no next node after head node.
	// Therefore, add new node to that next node.
	linkedList.headNode.nextNode = &Node{
		data:     *data,
		nextNode: nil,
	}

}

/*
*	Print out all the value in the nodes, which
* 	are linked in the list.
 */
func (linkedList *LinkedList) FindAll() {

	// There will be value only if there is a next node
	// after the head node.
	if linkedList.headNode.nextNode != nil {
		tempNode := linkedList.headNode.nextNode

		// Print out the value of the next node that comes
		// right after the head node.
		fmt.Printf("Fetched node value : %v \n", tempNode.data)

		// Print out the values of all the rest of the nodes,
		// by iterating through the linked list.
		for tempNode.nextNode != nil {
			fmt.Printf("Fetched node value : %v \n", tempNode.nextNode.data)

			tempNode = tempNode.nextNode
		}
	}
}
