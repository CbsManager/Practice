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

func main() {
	fmt.Println(fizzBuzzNormal(20))
	fmt.Println(fizbuzz1if(3))
}
