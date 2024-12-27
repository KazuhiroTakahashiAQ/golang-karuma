package main

import (
	"strconv"
)

func main() {
	print("Hello, World!")
}

func FizzBuzz(input int) string {
	if input%15 == 0 {
		return "FizzBuzz"
	} else if input%5 == 0 {
		return "Buzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(input)
	}
}
