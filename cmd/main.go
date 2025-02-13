package main

import (
	"fmt"

	"github.com/italanleal/go-array-utils/utils"
)

func main() {
	nums := []int{1, 2, 4, 5}

	fmt.Println(nums)
	fmt.Println(utils.Reverse(nums))
	fmt.Println(utils.Map(nums, func(value int) int {
		return (value * 2)
	}))
	fmt.Println(utils.Reduce(nums, func(acm int, cur int) int {
		return (acm + cur)
	}, 0))
}
