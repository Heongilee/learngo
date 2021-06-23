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

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func pop(arr []string) []string {
	if len(arr) >= 1 {
		arr = arr[:len(arr)-1]
	}

	return arr
}

func popleft(arr []string) []string {
	if len(arr) >= 1 {
		arr = arr[1:]
	}

	return arr
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	heongi := person{name: "heongi", age: 26, favFood: []string{"kimchi", "ramen"}}
	fmt.Println(heongi)

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		heongi := person{"heongi", 26, []string{"kimchi", "ramen"}}
		fmt.Println(heongi)
		fmt.Println("name:", heongi.name)
		fmt.Println("age:", heongi.age)
		fmt.Println("favFood:", heongi.favFood)
	*/
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		heongi := map[string]string{"name": "heongi", "age": "26"}
		fmt.Println(heongi)

		for key, value := range heongi {
			fmt.Println(key, value)
		}
	*/
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		students1 := []string{"Rod Johnson", "Yan Caroff", "adai"}
		students1 = append(students1, "Juergen Hoeller")
		fmt.Println(students1)
		students1 = pop(students1)
		fmt.Println(students1)
		students1 = popleft(students1)
		fmt.Println(students1)
		// students2 := [5]string{"Rod Johnson", "Yan Caroff", "adai"}
		// fmt.Println(students2)

		names := [5]string{"nico", "lynn", "dal"}
		fmt.Println(names)
		names[3] = "awesome"
		names[4] = "heonjjang"

		fmt.Print("[")
		for i, elem := range names {
			fmt.Print(elem)
			if i < len(names)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	*/
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		total := superAdd(1, 2, 3, 4, 5, 6)
		fmt.Println("result is", total)
	*/

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// fmt.Println(canIDrink(18))
	/*
		a := 2
		b := &a
		fmt.Println("The value of variable a is ", a)
		fmt.Println("the value of the variable pointed to by b is ", *b)
		*b = 20
		fmt.Println("The value of variable a is ", a)
		fmt.Println("the value of the variable pointed to by b is ", *b)
		// c := &b
		// fmt.Println("The value of variable a is ", a)
		// fmt.Println("The address of variable a is ", &a)
		// fmt.Println("The value of variable a is ", b)
		// fmt.Println("The address of variable a is ", &b)
		// fmt.Println("The value of variable a is ", c)
		// fmt.Println("The address of variable a is ", &c)
	*/
}
