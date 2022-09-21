package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func StartLinkedProblem() {
  fourth := &ListNode{4, nil}
  third := &ListNode{3, fourth}
  second := &ListNode{2, third}
  first := &ListNode{1, second}
  test := removeNthFromEnd(first, 2)
  for test != nil {
    fmt.Println(test.Val)
    test = test.Next
  }
}

func deleteNode(node *ListNode) {
    //this is in the interview list. The hardest part of this question could be solved by asking the interviewer what they mean by delete
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  //needs another attempt
    listCopy := &ListNode {
      Val: 0,
      Next: head,
    }
    i := 0
    for listCopy != nil {
      listCopy = listCopy.Next
      if i == n {
        listCopy.Val = listCopy.Next.Val
        listCopy.Next = listCopy.Next.Next
      }
      i++
    }
    return head
}