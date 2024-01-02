package main

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{
		Val: val,
		Next: nil,
	}
}