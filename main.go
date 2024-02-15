package main

import "ds/rnd/nestedlinkedlist"

func main() {

	nllService := nestedlinkedlist.NewNestedLinkedListService()

	dataOne := 6
	dataTwo := 5
	dataThree := 6

	nllService.AddData(&dataOne)
	nllService.AddData(&dataTwo)
	nllService.AddData(&dataThree)

	nllService.FindAll()

	nllService.DeleteCategory(&dataOne)

	nllService.FindAll()
}
