package leetcode

import "fmt"

func StartStringProblem() {
	fmt.Println("Start string problem")
	var str = "thisisatest"
	byteArray := []byte(str)
	reverseString(byteArray)
}

func reverseString(s []byte)  {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}