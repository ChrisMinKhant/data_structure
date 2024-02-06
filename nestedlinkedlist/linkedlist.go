package nestedlinkedlist

import (
	"fmt"
)

type LinkedList struct {
	headNode *Node
}

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

	if linkedList.headNode.nextNode != nil {
		tempNode := linkedList.headNode.nextNode

		for tempNode.nextNode != nil {
			tempNode = tempNode.nextNode
		}

		fmt.Print(" Line -> 33 | Node is added. \n")
		tempNode.nextNode = &Node{
			data:     *data,
			nextNode: nil,
		}

		return
	}

	fmt.Print(" Line -> 42 | Node is added. \n")
	linkedList.headNode.nextNode = &Node{
		data:     *data,
		nextNode: nil,
	}

	fmt.Printf(" Added value : %v \n", linkedList.headNode.nextNode.data)
}

func (linkedList *LinkedList) FindAll() {
	if linkedList.headNode.nextNode != nil {
		tempNode := linkedList.headNode.nextNode

		fmt.Printf("Fetched node value : %v \n", tempNode.data)

		for tempNode.nextNode != nil {
			fmt.Printf("Fetched node value : %v \n", tempNode.nextNode.data)

			tempNode = tempNode.nextNode
		}
	}
}
