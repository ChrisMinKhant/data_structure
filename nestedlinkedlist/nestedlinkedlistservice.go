package nestedlinkedlist

import "fmt"

type NestedLinkedListService struct {
	nestedLinkedList *NestedLinkedList
}

func NewNestedLinkedListService() *NestedLinkedListService {
	return &NestedLinkedListService{
		nestedLinkedList: NewNestedLinkedList(),
	}
}

func (service *NestedLinkedListService) AddData(value *int) {

	if service.nestedLinkedList != nil {

		fmt.Print("Non-nil nested linked list. \n")

		tempNestedLinkedList := service.nestedLinkedList

		tempLinkedList := tempNestedLinkedList.linkedList

		if tempLinkedList != nil {
			// find the right category to add the new node
			for tempLinkedList.headNode.nextNode.data != *value {

				tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
				tempLinkedList = tempNestedLinkedList.linkedList
			}
			tempLinkedList.Add(value)
		}

		tempLinkedList.Add(value)

		return
	}

	// create new category for requested value

	fmt.Print("Nil nested linked list. \n")

	tempLinkedList := NewLinkedList()

	tempLinkedList.Add(value)

	fmt.Print("After adding linked list to nested linked list.")
}

func (service *NestedLinkedListService) FindAll() {

	if service.nestedLinkedList != nil {
		service.nestedLinkedList.FindAll()
	}

}
