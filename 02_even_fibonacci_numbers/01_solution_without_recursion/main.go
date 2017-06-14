package main

import "fmt"

func main() {
	fibonacci := fibonacci()
	fmt.Println("Fibonaci bellow to 4 million:", fibonacci)

	evenFibonacci := evenNumbers(fibonacci)
	fmt.Println("Values even of Fibonacci bellow to 4 million:", evenFibonacci)

	fmt.Println("Sum of that values:", sumFibonacci(evenFibonacci))
}

func fibonacci() []int {
	primaryValue := 1
	secondValue := 2
	max := 4000000

	result := []int{1, 2}
	for {
		sumValues := primaryValue + secondValue

		if sumValues > max {
			break
		}

		result = append(result, sumValues)
		primaryValue = secondValue
		secondValue = sumValues
	}

	return result
}

func evenNumbers(numbers []int) []int {
	var even []int

	for _, number := range numbers {
		if number%2 == 0 {
			even = append(even, number)
		}
	}

	return even
}

func sumFibonacci(fibonacci []int) int {
	var sum int
	for _, number := range fibonacci {
		sum += number
	}

	return sum
}
