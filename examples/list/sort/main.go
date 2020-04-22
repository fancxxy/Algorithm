package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func quickSort(head *ListNode, tail *ListNode) *ListNode {
	if head == tail {
		return head
	}

	pivot := head.Val
	store := head
	curr := head.Next
	for curr != tail {
		if curr.Val < pivot {
			store = store.Next
			store.Val, curr.Val = curr.Val, store.Val
		}
		curr = curr.Next
	}

	store.Val, head.Val = head.Val, store.Val
	quickSort(head, store)
	quickSort(store.Next, tail)

	return store
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	var prev, curr *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	node1 := mergeSort(head)
	node2 := mergeSort(slow)

	if node1.Val > node2.Val {
		head = node2
		node2 = node2.Next
	} else {
		head = node1
		node1 = node1.Next
	}

	curr = head
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			curr.Next = node2
			node2 = node2.Next
		} else {
			curr.Next = node1
			node1 = node1.Next
		}
		curr = curr.Next
	}
	if node1 != nil {
		curr.Next = node1
	}
	if node2 != nil {
		curr.Next = node2
	}

	return head
}

func main() {
	list := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	// quickSort(list, nil)

	for node := mergeSort(list); node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
}
