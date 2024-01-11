package main

// 单链表节点
type ListNode struct {
	Val  interface{}
	Next *ListNode // 下一个节点指针
}

// 单链表结构
type LinkedList struct {
	Head *ListNode // 头节点
	Len  int       // 单链表长度
}

// 创建一个Node节点
func NewListNode(val interface{}) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 创建一个空的链表
func NewLinkedList() *LinkedList {
	// Head头节点，不存储数据
	return &LinkedList{
		Head: NewListNode(0),
		Len:  0,
	}
}

// 链表是否为空
// 判断链表头节点 Head 为 nil
func (l *LinkedList) IsEmpty() bool {
	if l.Head == nil {
		return true
	}
	return false
}

// 插入值为 val 的节点 newNode 到 preNode 后
// preNode -> newNode -> Node
func (l *LinkedList) Insert(preNode *ListNode, val interface{}) bool {
	if preNode == nil {
		return false
	}

	newNode := NewListNode(val)
	oldNext := preNode.Next
	preNode.Next = newNode
	newNode.Next = oldNext
	l.Len++
	return true
}

// 删除节点
func (l *LinkedList) Remove(n *ListNode) bool {
	if n == nil {
		return false
	}

}
