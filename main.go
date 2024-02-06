package main

import "ds/rnd/nestedlinkedlist"

func main() {

	nllService := nestedlinkedlist.NewNestedLinkedListService()

	dataOne := 6
	dataTwo := 8

	nllService.AddData(&dataOne)
	nllService.AddData(&dataTwo)

	nllService.FindAll()
}
