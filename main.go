package main

import "fmt"

func superAdd(numbers ...int) (cnt int) {
	// 방법 1
	for idx, number := range numbers {
		fmt.Println("numbers[", idx, "] =", number)
		cnt += number
	}

	// 방법 2
	/*
		defer fmt.Println()

		for _, number := range numbers {
			fmt.Print(number, " ")
			cnt += number
		}
	*/

	// 방법 3
	/*
		for i := 0; i < len(numbers); i++ {
			fmt.Println("numbers[", i, "] =", numbers[i])
			cnt += numbers[i]
		}
	*/
	return
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println("result is", total)
}
