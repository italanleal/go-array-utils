package main

import (
	"fmt"

	"github.com/italanleal/go-array-utils/utils"
)

func main() {
	// Exemplo de uso da função Map
	numbers := []int{1, 2, 3, 4, 5}
	squared := utils.Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared numbers:", squared)

	// Exemplo de uso da função Reduce
	sum := utils.Reduce(numbers, func(acc, cur int) int {
		return acc + cur
	}, 0)
	fmt.Println("Sum of numbers:", sum)

	// Exemplo de uso da função ForEach
	fmt.Println("Printing numbers:")
	utils.ForEach(numbers, func(n int) {
		fmt.Println(n)
	})

	// Exemplo de uso da função Reverse
	reversed := utils.Reverse(numbers)
	fmt.Println("Reversed numbers:", reversed)

	// Exemplo de uso da função Filter
	evenNumbers := utils.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)

}
