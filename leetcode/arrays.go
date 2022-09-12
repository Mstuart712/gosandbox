package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func StartProblem() {
	arr2 := []int{1,2,3,4,5}
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

  for i := 1; i < len(prices); i++ {
    if prices[i] > prices[i - 1] {
      a += prices[i] - prices[i - 1]	
    }
  }

  return a
}
