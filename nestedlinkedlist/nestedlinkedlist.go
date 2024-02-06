package nestedlinkedlist

import "fmt"

/*
*	Nested linked list struct which is the basic component of
*	the structure.
*	In this nested linked list, there is no head.
 */
type NestedLinkedList struct {
	linkedList           *LinkedList
	nextNestedLinkedList *NestedLinkedList
}

func NewNestedLinkedList() *NestedLinkedList {
	return &NestedLinkedList{
		linkedList:           nil,
		nextNestedLinkedList: nil,
	}
}

/*
*	Add new linked list to the nested linked list
*	Time complexity of O(n)
 */
func (nestedLinkedList *NestedLinkedList) Add(linkedList *LinkedList) {
	tempNestedLinkedList := nestedLinkedList

	if tempNestedLinkedList.linkedList != nil {
		for tempNestedLinkedList.nextNestedLinkedList != nil {
			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
		}

		tempNestedLinkedList.nextNestedLinkedList = &NestedLinkedList{
			linkedList:           linkedList,
			nextNestedLinkedList: nil,
		}

		return
	}

	tempNestedLinkedList.linkedList = linkedList

}

func (nestedLinkedList *NestedLinkedList) FindAll() {
	tempNestedLinkedList := nestedLinkedList

	loopCount := 1

	for tempNestedLinkedList.linkedList != nil {

		fmt.Printf(" Loop Count : %v \n", loopCount)
		loopCount++

		tempNestedLinkedList.linkedList.FindAll()

		if tempNestedLinkedList.nextNestedLinkedList != nil {

			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

			continue
		}

		return

	}
}
