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

/*
*	Declares nested linked list and return memory address of it.
 */
func NewNestedLinkedList() *NestedLinkedList {
	return &NestedLinkedList{
		linkedList:           nil,
		nextNestedLinkedList: nil,
	}
}

/*
*	Adds new linked list to the nested linked list
*	Time complexity of O(n)
 */
func (nestedLinkedList *NestedLinkedList) Add(requestedLinkedList *LinkedList) {

	// Declares temporary nested linke list to operate on.
	tempNestedLinkedList := nestedLinkedList

	// Check if the nested linked list to operate on is null.
	if tempNestedLinkedList != nil {

		// If the linked list of the nested linked list is not null,
		// check if the linked list inside the nested linked list is null.
		if tempNestedLinkedList.linkedList != nil {

			// If the linked list inside the nested linked list is not null,
			// iterate through to the end of the nested linked list, where
			// there is nil in the next nested linked list.
			for tempNestedLinkedList.nextNestedLinkedList != nil {
				tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
			}

			// Add new nested linked list with the linked list value
			// to the next nested linked list of the nested linked list,
			// once the iteration reaches the end.
			tempNestedLinkedList.nextNestedLinkedList = &NestedLinkedList{
				linkedList:           requestedLinkedList,
				nextNestedLinkedList: nil,
			}

			return
		}

		// If the linked list of the nested linked list is null,
		// add the linked list to it.
		tempNestedLinkedList.linkedList = requestedLinkedList
	}
}

/*
*	Iterates all the linked list of the nested linked list and
* 	prints out the values in the linked list by calling the FindAll() function
* 	of the linked list.
 */
func (nestedLinkedList *NestedLinkedList) FindAll() {

	// declares tempoaray nested linked list to operate on.
	tempNestedLinkedList := nestedLinkedList

	// Iterates through all the linked list of the nested linked list.
	for tempNestedLinkedList.linkedList != nil {

		// Then calls the FindAll() function of each linked list.
		tempNestedLinkedList.linkedList.FindAll()

		if tempNestedLinkedList.nextNestedLinkedList != nil {

			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

			continue
		}

		return

	}
}
