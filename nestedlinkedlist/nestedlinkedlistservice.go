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

			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList

			continue
		}

		tempLinkedList = NewLinkedList()

		tempLinkedList.Add(value)

		tempNestedLinkedList.Add(tempLinkedList)

		return
	}

	tempLinkedList := NewLinkedList()

	tempLinkedList.Add(value)

	tempNestedLinkedList.Add(tempLinkedList)

	// fmt.Print("Nested linked list is empty. \n")

	// tempLinkedList := tempNestedLinkedList.linkedList

	// if tempLinkedList != nil {
	// 	if tempLinkedList.headNode.nextNode != nil {
	// 		if tempLinkedList.headNode.nextNode.data == *value {

	// 			fmt.Printf("Found a category for value : %v \n", *value)

	// 			tempLinkedList.Add(value)

	// 			return
	// 		}
	// 	}
	// } else {

	// 	fmt.Printf("Does not found a category for value : %v \n", *value)

	// 	tempLinkedList = NewLinkedList()
	// 	tempLinkedList.Add(value)

	// 	return
	// }

}

func (service *NestedLinkedListService) FindAll() {

	if service.nestedLinkedList != nil {
		fmt.Print("Finding all value ...")
		service.nestedLinkedList.FindAll()
	}

}
