package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pop2(stack []float64) ([]float64, float64, float64) {
	var ab []float64
	stack, ab = stack[:len(stack)-2], stack[len(stack)-2:]
	return stack, ab[0], ab[1]
}

func ReversePolishNotation(s string) float64 {
	var stack []float64
	var tokens = strings.Fields(s)
	var value float64
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			var a, b float64
			if len(stack) < 2 {
				return 0
			}
			stack, b, a = pop2(stack)
			switch token {
			case "+":
				value = a + b
			case "-":
				value = a - b
			case "*":
				value = a * b
			case "/":
				value = a / b
			}
		default:
			var err error
			value, err = strconv.ParseFloat(token, 64)
			if err != nil {
				return 0
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return 0
	}
	return stack[0]
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	in := bufio.NewReader(os.Stdin)
	polish, _ := in.ReadString('\n')

	fmt.Println(polish)
	polish = strings.ReplaceAll(polish, "(", " ")
	polish = strings.ReplaceAll(polish, ")", " ")
	fmt.Println(ReversePolishNotation(Reverse(polish)))

}
