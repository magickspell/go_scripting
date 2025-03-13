package main

import (
	"fmt"
)

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Append adds a new element with the given value to the end of the list.
func (l *List[T]) Append(value T) {
	newNode := &List[T]{val: value}
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Print prints all the elements in the list.
func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Print(current.val, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Contains checks if the list contains a specific value.
func (l *List[T]) Contains(value T, equals func(T, T) bool) bool {
	current := l
	for current != nil {
		if equals(current.val, value) {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	// Create a new list with an initial value.
	list := &List[int]{val: 1}

	// Append some values to the list.
	list.Append(2)
	list.Append(3)
	list.Append(4)

	// Print the list.
	list.Print() // Output: 1 -> 2 -> 3 -> 4 -> nil

	// Get the length of the list.
	fmt.Println("Length:", list.Length()) // Output: Length: 4

	// Check if the list contains a specific value.
	equals := func(a, b int) bool { return a == b }
	fmt.Println("Contains 3:", list.Contains(3, equals)) // Output: Contains 3: true
	fmt.Println("Contains 5:", list.Contains(5, equals)) // Output: Contains 5: false

	stringList := &List[string]{val: "one"}
	stringList.Append("two")
	stringList.Append("three")
	stringList.Append("four")
	equalsString := func(a, b string) bool { return a == b }
	fmt.Println("Contains 'one':", stringList.Contains("one", equalsString))
	fmt.Println("Contains 'five':", stringList.Contains("five", equalsString))
}
