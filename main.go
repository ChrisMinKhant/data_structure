package main

import "ds/rnd/nestedlinkedlist"

func main() {

	nllService := nestedlinkedlist.NewNestedLinkedListService()

	dataOne := 6
	dataTwo := 10
	dataThree := 6
	dataFour := 10

	nllService.AddData(&dataOne)
	nllService.AddData(&dataTwo)
	nllService.AddData(&dataThree)
	nllService.AddData(&dataFour)

	nllService.FindAll()

	nllService.DeleteCategory(&dataOne)

	nllService.FindAll()
}
