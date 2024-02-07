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
func (nestedLinkedList *NestedLinkedList) Add(requestedLinkedList *LinkedList) {
	tempNestedLinkedList := nestedLinkedList

	if tempNestedLinkedList != nil {
		if tempNestedLinkedList.linkedList != nil {
			for tempNestedLinkedList.nextNestedLinkedList != nil {
				tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
			}

			tempNestedLinkedList.nextNestedLinkedList = &NestedLinkedList{
				linkedList:           requestedLinkedList,
				nextNestedLinkedList: nil,
			}

			return
		}

		tempNestedLinkedList.linkedList = requestedLinkedList
	}
}

func (nestedLinkedList *NestedLinkedList) FindAll() {
	tempNestedLinkedList := nestedLinkedList

	for tempNestedLinkedList.linkedList != nil {

		tempNestedLinkedList.linkedList.FindAll()

		if tempNestedLinkedList.nextNestedLinkedList != nil {

			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

			continue
		}

		return

	}
}
