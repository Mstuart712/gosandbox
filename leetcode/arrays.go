package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func StartProblem() {
	arr2 := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(arr2))
}

//remove duplicates of sorted int array
func removeDuplicates(nums []int) int {
    answer := 0
    max := math.MaxInt32
    for i := len(nums) - 1; i > 0; i-- {
        if nums[i] == nums[i-1] {
            nums[i] = max
            answer++
        }
    }
    sort.Ints(nums)
    return len(nums) - answer
}

func maxProfit(prices []int) int {
	a := 0
    for i := 0; i < len(prices); i++ {
		if i + 1 < len(prices) && prices[i] < prices[i + 1] {
			for j := i; j < len(prices); j++ {
				if prices[j + 1] <  prices[j] {
					a = a + (prices[j] - prices[i])
					i = j
					break
				} 	
			}
		}
	}
	return a
}
