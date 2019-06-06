package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value string
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type SymplyLinkList struct {
	Front  *Node
	Length int
}

func (s *SymplyLinkList) Init() *SymplyLinkList {
	s.Length = 0
	return s
}

// New 一个链表
func New() *SymplyLinkList {
	return new(SymplyLinkList).Init()
}

// 返回头节点
func (s *SymplyLinkList) FrontLink() *Node {
	return s.Front
}

// 返回尾节点
func (s *SymplyLinkList) BackLink() *Node {
	currentNode := s.Front
	for currentNode != nil && currentNode.next != nil { // 这种 for 循环也很有意思
		currentNode = currentNode.next
	}

	// for {
	// 	if currentNode.next != nil {
	// 		currentNode = currentNode.next
	// 	} else {
	// 		break
	// 	}
	// }
	return currentNode
}

// 添加尾节点
func (s *SymplyLinkList) AppendLink(n *Node) {
	if s.Front == nil {
		s.Front = n
	} else {
		currentNode := s.Front
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
	s.Length++
}

// 添加头节点
func (s *SymplyLinkList) PrependLink(n *Node) {
	if s.Front == nil {
		s.Front = n
	} else {
		n.next = s.Front
		s.Front = n.next
	}
	s.Length++
}

// 在指定节点前添加节点

// 在指定节点后添加节点

// 删除指定节点

func main() {
	// BasicOperation()

}

func BasicOperation() {
	alist := list.New()
	alist.PushBack("a")
	alist.PushBack("b")
	fmt.Println(alist.Len())

	for e := alist.Front(); e != nil; e = e.Next() { // 这种 for 循环
		fmt.Println(e.Value)
	}

	alist.Remove(alist.Back())

	for e := alist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
