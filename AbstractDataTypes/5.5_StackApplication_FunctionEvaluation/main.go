package main

import (
	"GenericDataStructuresAndAlgorithms/AbstractDataTypes/5.4_CompareEfficiencyOfNodeAndSliceStacks/NodeStack"
	"fmt"
)

var values map[string]float64

func evaluate(postfix string) float64 {
	operandStack := NodeStack.Stack[float64]{}

	for index := 0; index < len(postfix); index++ {
		ch := string(postfix[index])
		if ch >= "a" && ch <= "z" {
			operandStack.Push(values[ch])
		} else { // 'ch' is an operator
			operand1 := operandStack.Pop()
			operand2 := operandStack.Pop()
			if ch == "+" {
				operandStack.Push(operand1 + operand2)
			} else if ch == "-" {
				operandStack.Push(operand2 - operand1)
			} else if ch == "*" {
				operandStack.Push(operand1 * operand2)
			} else if ch == "/" {
				operandStack.Push(operand2 / operand1)
			}
		}
	}

	return operandStack.Top()
}

func precedence(symbol1, symbol2 string) bool {
	//Returns true if 'symbol1' has higher precedence than 'symbol2'
	if (symbol1 == "+" || symbol1 == "-") &&
		(symbol2 == "(" || symbol2 == "/") {
		return false
	} else if (symbol1 == "(" && symbol2 != ")") || symbol2 == "(" {
		return false
	} else {
		return true
	}
}

func isPresent(symbol string, operators []string) bool {
	for i := 0; i < len(operators); i++ {
		if symbol == string(operators[i]) {
			return true
		}
	}

	return false
}

func infixpostfix(infix string) (postfix string) {
	operators := []string{"+", "-", "*", "/", "(", ")"}
	postfix = ""
	nodeStack := NodeStack.Stack[string]{}

	for index := 0; index < len(infix); index++ {
		newSymbol := string(infix[index])
		if newSymbol == " " || newSymbol == "\n" {
			continue
		}
		if newSymbol >= "a" && newSymbol <= "z" {
			postfix += newSymbol
		}

		if isPresent(newSymbol, operators) {
			if !nodeStack.IsEmpty() {
				topSymbol := nodeStack.Top()
				if precedence(topSymbol, newSymbol) == true {
					if topSymbol != "(" {
						postfix += topSymbol
					}
					nodeStack.Pop()
				}
			}

			if newSymbol != ")" {
				nodeStack.Push(newSymbol)
			} else { // Pop nodeStack down to first left parenthesis
				for {
					if nodeStack.IsEmpty() == true {
						break
					}

					ch := nodeStack.Top()
					if ch != "(" {
						postfix += ch
						nodeStack.Pop()
					} else {
						nodeStack.Pop()
						break
					}
				}
			}
		}
	}

	for {
		if nodeStack.IsEmpty() == true {
			break
		}

		if nodeStack.Top() != "(" {
			postfix += nodeStack.Top()
			nodeStack.Pop()
		}
	}

	return postfix
}

func main() {
	//Step 1
	postFix := infixpostfix("a + (b - c) / (d * e)")
	fmt.Println(postFix)

	//Step 2
	values = make(map[string]float64)
	values["a"] = 10
	values["b"] = 5
	values["c"] = 2
	values["d"] = 4
	values["e"] = 3

	result := evaluate(postFix)
	fmt.Println("function evaluates to: ", result)
}
