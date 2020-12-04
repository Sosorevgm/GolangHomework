package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, res int
	var op string

	fmt.Print("Введите первое число: ")
	a = getNumber()

	fmt.Print("Введите второе число: ")
	b = getNumber()

	fmt.Print("Введите арифметическую операцию (+, -, *, /). Для отмены введите: exit. ")
	op = getOperation()

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	fmt.Println("Ответ: " + strconv.Itoa(res))
}

func getNumber() int {
	var number string
	for {
		fmt.Scanln(&number)
		var res, err = strconv.Atoi(number)
		if err != nil {
			fmt.Print("Введите целое число: ")
			continue
		}
		return res
	}
}

func getOperation() string {
	var operator string
	for {
		var input string
		fmt.Scanln(&input)
		if input == "+" || input == "-" || input == "*" || input == "/" {
			operator = input
			break
		} else if input == "exit" {
			os.Exit(1)
		}
		fmt.Print("Неверная операция, повторите ввод или наберите exit для отмены. ")
	}
	return operator
}
