package main

import "fmt"

type Ordered interface {
	~float64 | ~int | ~string
}

type Stack[T Ordered] struct {
	items []T
}

func getZero[T Ordered]() T {
	var result T
	return result
}

//Methods

func (stack *Stack[T]) Push(item T) {
	//item is added to the right-most position in the slice
	if item != getZero[T]() { //We exclude item if it is getZero[T]()
		stack.items = append(stack.items, item)
	}
}

func (stack *Stack[T]) Pop() T {
	length := len(stack.items)

	if length > 0 {
		returnValue := stack.items[length-1]
		stack.items = stack.items[:(length - 1)]

		return returnValue
	} else {
		return getZero[T]()
	}
}

func (stack Stack[T]) Top() T {
	length := len(stack.items)

	if length > 0 {
		return stack.items[length-1]
	} else {
		return getZero[T]()
	}
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func useFirstImplementation() {
	//Create a stack of names
	nameStack := Stack[string]{}
	nameStack.Push("Armin")
	nameStack.Push("Parsa")

	topOfStack := nameStack.Top()
	if topOfStack != getZero[string]() {
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}

	//Popping 4 time
	poppedFromStack := nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0) //Problem since 0 is the zero value for int

	top := intStack.Top()
	if top != getZero[int]() {
		fmt.Printf("\nTop of intStack is %d", top)
	}

	popFromStack := intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}
