package main

import (
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
		var res, _ = getPrimeCount(int64(tmp))
		fmt.Println(res)
	}

}

func getPrimeCount(num int64) (string, error) {
	var counter int64
	if num < 0 {
		panic("Число должно быть положительным")
	} else if num == 0 || num == 1 {
		return "Найдено простых чисел: 0", nil
	}
	var i int64 = 2
	for ; i < num; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			counter++
		}
	}
	return "Найдено простых чисел: " + strconv.Itoa(int(counter)), nil
}
