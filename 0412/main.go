package main

import "strconv"

func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		divisibleBy3 := isDivisibleByX(i, 3)
		divisibleBy5 := isDivisibleByX(i, 5)
		if divisibleBy3 && divisibleBy5 {
			answer = append(answer, "FizzBuzz")
		} else if divisibleBy3 {
			answer = append(answer, "Fizz")
		} else if divisibleBy5 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}
	}
	return answer
}

func isDivisibleByX(n, x int) bool {
	return n%x == 0
}
