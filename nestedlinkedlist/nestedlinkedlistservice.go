package nestedlinkedlist

type NestedLinkedListService struct {
	nestedLinkedList *NestedLinkedList
}

func NewNestedLinkedListService() *NestedLinkedListService {
	return &NestedLinkedListService{
		nestedLinkedList: nil,
	}
}

func (service *NestedLinkedListService) addData(value *int) {

	if service.nestedLinkedList != nil {

		tempNestedLinkedList := service.nestedLinkedList

		tempLinkedList := tempNestedLinkedList.linkedList

		// find the right category to add the new node
		for tempLinkedList.headNode.nextNode.data != *value {

			tempNestedLinkedList = tempNestedLinkedList.nextNestedLinkedList
			tempLinkedList = tempNestedLinkedList.linkedList

		}

		tempLinkedList.Add(value)

		return
	}

	// create new category for requested value

	tempLinkedList := NewLinkedList()

	tempNestedLinkedList := NewNestedLinkedList()

	tempLinkedList.Add(value)

	tempNestedLinkedList.Add(tempLinkedList)
}
