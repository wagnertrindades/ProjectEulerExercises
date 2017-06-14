package main

import "fmt"

func main() {
	startFibonacci := []int{1, 2}
	fibonacci := fibonacciRecursive(startFibonacci)
	fmt.Println("Fibonaci bellow to 4 million:", fibonacci)

	evenFibonacci := evenNumbers(fibonacci)
	fmt.Println("Values even of Fibonacci bellow to 4 million:", evenFibonacci)

	fmt.Println("Sum of that values:", sumFibonacci(evenFibonacci))
}

func fibonacciRecursive(fibonacci []int) []int {
	lastValue := fibonacci[len(fibonacci)-1]
	penultimateValue := fibonacci[len(fibonacci)-2]

	sum2LastValues := lastValue + penultimateValue
	if sum2LastValues > 4000000 {
		return fibonacci
	}

	fibonacci = append(fibonacci, sum2LastValues)
	return fibonacciRecursive(fibonacci)
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
