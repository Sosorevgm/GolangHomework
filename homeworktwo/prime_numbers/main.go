package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("Введите положительное число")
	var input string
	fmt.Scanln(&input)
	var tmp, err = strconv.Atoi(input)
	if err == nil {
		var res, getCountErr = getPrimeCount(int64(tmp))
		if getCountErr == nil {
			fmt.Println(res)
		} else {
			fmt.Println(getCountErr)
		}
	}
}

func getPrimeCount(num int64) (string, error) {
	var counter int64
	if num < 0 {
		return "", errors.New("число должно быть положительным")
		//return "", new(NegativeNumberError)
	} else if num == 0 || num == 1 {
		//return "", new(NegativeNumberError)
		return "", errors.New("число должно быть положительным")
	}
	var i int64 = 2
	for ; i < num; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			counter++
		}
	}
	return "Найдено простых чисел: " + strconv.Itoa(int(counter)), nil
}

//type NegativeNumberError struct {
//	err string
//}
//
//func (e *NegativeNumberError) Error() string {
//	return "число должно быть положительным"
//}
