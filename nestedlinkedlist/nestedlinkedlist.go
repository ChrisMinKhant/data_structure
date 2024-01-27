package nestedlinkedlist

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

	for tempNestedLinkedList.nextNestedLinkedList != nil {
		tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
	}

	tempNestedLinkedList.nextNestedLinkedList = &NestedLinkedList{
		linkedList:           linkedList,
		nextNestedLinkedList: nil,
	}
}
