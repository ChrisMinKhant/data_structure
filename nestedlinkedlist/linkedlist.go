package nestedlinkedlist

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

		tempNode.nextNode = &Node{
			data:     *data,
			nextNode: nil,
		}

		return
	}

	linkedList.headNode.nextNode = &Node{
		data:     *data,
		nextNode: nil,
	}
}
