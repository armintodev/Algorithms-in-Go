package main

import "fmt"

type StackOfAny[T any] struct {
	items []T
}

// Use StackOfAny struct

func (stack *StackOfAny[T]) Push(item T) {
	//item is added to the right-most position in the slice
	stack.items = append(stack.items, item)
}

func (stack *StackOfAny[T]) Pop() T {
	length := len(stack.items)

	returnValue := stack.items[length-1]
	stack.items = stack.items[:(length - 1)]

	return returnValue
}

func (stack StackOfAny[T]) Top() T {
	length := len(stack.items)
	return stack.items[length-1]
}

func (stack StackOfAny[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func useSecondImplementation() {
	//Create a stack of names
	nameStack := StackOfAny[string]{}
	nameStack.Push("Armin")
	nameStack.Push("Parsa")

	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}

	//Popping 4 time
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	intStack := StackOfAny[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0)

	if !intStack.IsEmpty() {
		top := intStack.Top()
		fmt.Printf("\nTop of intStack is %d", top)
	}

	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}
