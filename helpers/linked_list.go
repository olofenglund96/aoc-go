package helpers

type linkedListElement struct {
	next *linkedListElement
	val  interface{}
}

type linkedList struct {
	head *linkedListElement
}

func NewLinkedList(elements []interface{}) {}
