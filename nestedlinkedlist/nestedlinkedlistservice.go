package nestedlinkedlist

import "fmt"

/*
*	This is nested linked list service file, which mainly includes
*	basic operations that can be done on nested linked list.
*	All the operations that can be done on nested linked list have to be called
*	only from this service.
 */

type NestedLinkedListService struct {
	nestedLinkedList *NestedLinkedList
}

/*
*	Declares new nested linked list service	variable
*	and return the memory address of it.
 */
func NewNestedLinkedListService() *NestedLinkedListService {
	return &NestedLinkedListService{
		nestedLinkedList: NewNestedLinkedList(),
	}
}

/*
*	Adds single value to the node of the nested linked list,
*	according to their category.
*	The result nested linked list will be categorized by default.
 */
func (service *NestedLinkedListService) AddData(value *int) {

	// Declares temporary nested linked list to operate on.
	tempNestedLinkedList := service.nestedLinkedList

	// Iterates through to the end of the nested linked list.
	// By then, we will get separate nested linked list to concern.
	for tempNestedLinkedList != nil {

		tempLinkedList := tempNestedLinkedList.linkedList

		// Check if the linked list of the nested linked list is nil.
		// This if-statement is to ensure that there is intialized linked list inside
		// the nested linked list, to avoid nil pointer exception.
		if tempLinkedList != nil {

			// Check if the next node right after the head node the linked list is nil.
			if tempLinkedList.headNode.nextNode != nil {

				// Check if the value of the next node is equal to the value to be added.
				if tempLinkedList.headNode.nextNode.data == *value {

					fmt.Printf("Found a category for value : %v \n", *value)

					tempLinkedList.Add(value)

					return
				}

				// We will checking only the value of the first node right after the head node.
				// Because, we can easily know that the linked list is the right category for the
				// value to be added, just by checking the first node.

			} else {

				tempLinkedList.Add(value)

				return
			}

			// Swap the temporary nested linked list with the next nested linked list,
			// only if there is one.
			if tempNestedLinkedList.nextNestedLinkedList != nil {

				tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

				continue
			}

		}

		// If there is no right category for the valued to be added,
		// we will added new nested linked with the value already in it.
		tempLinkedList = NewLinkedList()

		tempLinkedList.Add(value)

		tempNestedLinkedList.Add(tempLinkedList)

		return
	}

	// This is the case where, there is no single nested linked list.
	tempLinkedList := NewLinkedList()

	tempLinkedList.Add(value)

	tempNestedLinkedList.Add(tempLinkedList)

}

/*
*	Prints out all the value of the linked list of the nested linked list.
 */
func (service *NestedLinkedListService) FindAll() {

	// Check the nested linked list is not nil.
	// This is to avoid nil pointer exception.
	if service.nestedLinkedList != nil {
		fmt.Print("Finding all value ... \n")
		service.nestedLinkedList.FindAll()
	}

}
