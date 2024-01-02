package main

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode // 指向第一个节点
	Len  int
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: NewListNode(0),
		Len:  0,
	}
}

func (l *LinkedList) Insert(p *ListNode, val interface{}) bool {

}
