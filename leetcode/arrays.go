package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func StartProblem() {
	arr2 := []int{98, 99, 99, 100}
	fmt.Println(removeDuplicates(arr2))
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
