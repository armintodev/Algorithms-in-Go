package SliceStack

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
