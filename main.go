package main

import (
	"fmt"
	"strconv"
)

func fizzBuzzNormal(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}
func fizbuzz1if(n int) string {
	fizz := [3]string{"Fizz", "", ""}
	buzz := [6]string{"Buzz", "", "", "", "", ""}
	if n%3 == 0 || n%5 == 0 {
		return fizz[n%3] + buzz[n%5]
	}
	return strconv.Itoa(n)
}

func fizznoif(n int) string {
	Fizz := map[bool]string{
		true:  "Fizz",
		false: "",
	}

	Buzz := map[bool]string{
		true:  "Buzz",
		false: "",
	}

	num := map[bool]string{
		true:  strconv.Itoa(n),
		false: "",
	}

	return Fizz[n%3 == 0] + Buzz[n%5 == 0] + num[n%3 != 0 && n%5 != 0]
}

func main() {
	fmt.Println(fizzBuzzNormal(20))
	fmt.Println(fizbuzz1if(1))
	fmt.Println(fizznoif(5))
	fmt.Println(RomanToNumber("XXI"))
}
