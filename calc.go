package main

import (
	"fmt"
	"bufio"
	"os"
	"stack"
	"strings"
	"bytes"
	"strconv"
	"reflect"
)

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var result = compute(line)
	switch result.(type) {
	case int64:
		fmt.Print(result.(int64))
	case float64:
		fmt.Print(result.(float64))
	}
}

func compute(cmd string) interface{} {
	var operatorStack = stack.NewStack()
	var operandStack = stack.NewStack()
	var last_char byte

	for i := 0; i < len(cmd); {
		if isOperand(cmd[i]) {
			last_char = cmd[i]
			var buffer bytes.Buffer
			buffer.WriteString(string([]byte{cmd[i]}))
			i++
			for {
				if (i < len(cmd) && isOperand(cmd[i])) {
					buffer.WriteString(string([]byte{cmd[i]}))
					i++
				} else {
					break;
				}
			}

			var num_str = buffer.String()

			if (strings.Contains(num_str, ".")) {
				num, err := strconv.ParseFloat(num_str, 64)
				if err == nil {
					operandStack.Push(num)
				} else {
					panic("Error pushing float64 val '" + num_str + "' to stack")
				}
			} else {
				num, err := strconv.ParseInt(num_str, 10, 64)
				if err == nil {
					operandStack.Push(num)
				} else {
					panic("Error pushing int64 val '" + num_str + "' to stack")
				}
			}
		} else if isOperator(cmd[i]) {
			last_char = cmd[i]
			for !operatorStack.IsEmpty() && precedence(operatorStack.Top().(byte)) >= precedence(cmd[i]) {
				operatorStack, operandStack = apply(operatorStack, operandStack)
			}
			operatorStack.Push(cmd[i])
			i++
		} else if cmd[i] == ' ' {
			i++
			if isOperand(cmd[i]) && isOperand(last_char) {
				panic("Syntax error - operands cannot have spaces in between them")
			}
		} else if cmd[i] == '(' {
			if isOperator(last_char) || i == 0 {
				i++
				var paren_count = 1
				var paren string
				var j int
				for j = i; j < len(cmd); j++ {
					if cmd[j] == '(' {
						paren_count++
					} else if cmd[j] == ')' {
						paren_count--
						if paren_count == 0 {
							break;
						}
					}

					paren += string([]byte{cmd[j]})
				}
				operandStack.Push(compute(paren))
				i = j + 1
			} else {
				panic("Operator must come before '(' : last_char = " + string([]byte{last_char}))
			}
		} else {
			panic("Illegal Character '" + string([]byte{cmd[i]}) + "'")
		}
	}

	for !operatorStack.IsEmpty() {
		operatorStack, operandStack = apply(operatorStack, operandStack)
	}

	r := operandStack.Pop()
	return r
}

func isOperand(c byte) bool {
	if (c >= '0' && c <= '9') || c == '.' {
		return true
	} else {
		return false
	}
}

func isOperator(c byte) bool {
	return strings.Contains("+-*/", string(c))
}

func precedence(op byte) uint8 {
	switch op {
	case '+', '-': return 0
	case '*', '/': return 1
	default: panic("illegal operator")
	}
}

func apply(operatorStack stack.Stack, operandStack stack.Stack) (stack.Stack, stack.Stack) {
	op := operatorStack.Pop().(byte)

	right := reflect.ValueOf(operandStack.Pop())
	left := reflect.ValueOf(operandStack.Pop())

	if right.Kind() == left.Kind() && right.Kind() == reflect.Int64 {
		switch op {
		case '+': operandStack.Push(left.Int() + right.Int())
		case '-': operandStack.Push(left.Int() - right.Int())
		case '*': operandStack.Push(left.Int() * right.Int())
		case '/': operandStack.Push(left.Int() / right.Int())
		default: panic("Illegal Operator '" + string([]byte{op}) + "'")
		}
	} else if right.Kind() == left.Kind() && right.Kind() == reflect.Float64 {
		switch op {
		case '+': operandStack.Push(left.Float() + right.Float())
		case '-': operandStack.Push(left.Float() - right.Float())
		case '*': operandStack.Push(left.Float() * right.Float())
		case '/': operandStack.Push(left.Float() / right.Float())
		default: panic("Illegal Operator '" + string([]byte{op}) + "'")
		}
	} else {
		if left.Kind() != reflect.Float64 {
			switch op {
			case '+': operandStack.Push(float64(left.Int()) + right.Float())
			case '-': operandStack.Push(float64(left.Int()) - right.Float())
			case '*': operandStack.Push(float64(left.Int()) * right.Float())
			case '/': operandStack.Push(float64(left.Int()) / right.Float())
			default: panic("Illegal Operator '" + string([]byte{op}) + "'")
			}
		} else if right.Kind() != reflect.Float64 {
			switch op {
			case '+': operandStack.Push(left.Float() + float64(right.Int()))
			case '-': operandStack.Push(left.Float() - float64(right.Int()))
			case '*': operandStack.Push(left.Float() * float64(right.Int()))
			case '/': operandStack.Push(left.Float() / float64(right.Int()))
			default: panic("Illegal Operator '" + string([]byte{op}) + "'")
			}
		} else {
			fmt.Println(left.Type())
			panic("Reflection not found (Critical Error)")
			fmt.Println(right.Kind())
			fmt.Println(left.Kind())
		}
	}
	return operatorStack, operandStack
}

