package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func StartProblem() {
	arr2 := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(arr2))
	
}

func singleNumber(nums []int) int {
  arrMap := make(map[int]bool)
  for i := 0; i < len(nums); i++ {
    if _, ok := arrMap[nums[i]]; ok {
      arrMap[nums[i]] = true
    } else {
      arrMap[nums[i]] = false
    }
  }
  for key, element := range arrMap {
    if element == false {
      return key
    }
  }
}

func containsDuplicate(nums []int) bool {
    arrMap := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
      if _, ok := arrMap[nums[i]]; ok {
        return true
      } else {
        arrMap[nums[i]] = true
      }
    }
    return false
}

func rotate(nums []int, k int) {
  if k > len(nums) {
    fmt.Println("we hit this")
    k = k % len(nums)
    fmt.Println(k)
  }
  start := nums[len(nums)-k:]	
	end := nums[:len(nums)-k]
  result := append(start, end...)
  for i := 0; i < len(nums); i++ {
      nums[i] = result[i]
  }
	fmt.Println("nums: ", nums)
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
