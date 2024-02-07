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

	tempNestedLinkedList := service.nestedLinkedList

	for tempNestedLinkedList != nil {

		tempLinkedList := tempNestedLinkedList.linkedList

		if tempLinkedList != nil {
			if tempLinkedList.headNode.nextNode != nil {
				if tempLinkedList.headNode.nextNode.data == *value {

					fmt.Printf("Found a category for value : %v \n", *value)

					tempLinkedList.Add(value)

					return
				}
			} else {
				tempLinkedList.Add(value)

				return
			}

			if tempNestedLinkedList.nextNestedLinkedList != nil {

				tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

				continue
			}

		}

		tempLinkedList = NewLinkedList()

		tempLinkedList.Add(value)

		tempNestedLinkedList.Add(tempLinkedList)

		return
	}

	tempLinkedList := NewLinkedList()

	tempLinkedList.Add(value)

	tempNestedLinkedList.Add(tempLinkedList)

}

func (service *NestedLinkedListService) FindAll() {

	if service.nestedLinkedList != nil {
		fmt.Print("Finding all value ... \n")
		service.nestedLinkedList.FindAll()
	}

}
