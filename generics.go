package main

import "fmt"

func Mapkeys[K comparable, V any](m map[K]V) []K {
	r:= make([]K, 0, len(m))

	for k:= range m {
		r = append(r,k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (list *List[T]) Push(value T) {
	if list.tail == nil {
		list.head = &element[T]{val: value}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: value}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elements []T
	for e:=list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func main() {
	var m = map[int]string{1: "1", 2: "2", 3:"3"}
	fmt.Println("keys", Mapkeys(m))

	_ = Mapkeys[int, string](m)


	list := List[int]{}
	list.Push(10)
	list.Push(23)
	list.Push(32)
	fmt.Println("list:", list.GetAll())
}