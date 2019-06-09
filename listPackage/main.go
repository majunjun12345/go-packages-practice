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

// 格式化显示链表
func (s *SymplyLinkList) String() string {
	current := s.Front
	str := ""
	for current != nil {
		str += current.Value + ">>"
		current = current.next
	}
	return str
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
		s.Front = n
	}
	s.Length++
}

// 在指定节点前添加节点
func (s *SymplyLinkList) InsertBefore(n *Node, before *Node) {
	if s.Front == before {
		n.next = s.Front
		s.Front = n.next
		s.Length++
	} else {
		current := s.Front
		for current != nil && current.next != nil && current.next != before {
			current = current.next
		}

		if current.next == before {
			n.next = before
			current.next = n
			s.Length++
		} else {
			fmt.Printf("no such node:%+v\n", before)
		}
	}
}

// 在指定节点后添加节点
func (s *SymplyLinkList) InsertAfter(n *Node, after *Node) {
	n.next = after.next
	after.next = n
}

// 删除指定节点
func (s *SymplyLinkList) Remove(n *Node) {
	if s.Front == n {
		s.Front = n.next
		s.Length--
	} else {
		current := s.Front
		for current.next != n && current != nil && current.next != nil {
			current = current.next
		}
		if current.next == n {
			current.next = current.next.next
			s.Length--
		}
	}
}

func main() {
	// BasicOperation()
	TestSympleLinkList()
}

func TestSympleLinkList() {
	l := New()
	n1 := &Node{Value: "1"}
	n2 := &Node{Value: "2"}
	n3 := &Node{Value: "3"}
	n4 := &Node{Value: "4"}
	n5 := &Node{Value: "5"}
	n6 := &Node{Value: "6"}
	n7 := &Node{Value: "7"}
	n8 := &Node{Value: "8"}
	n9 := &Node{Value: "9"}
	n10 := &Node{Value: "10"}
	l.AppendLink(n1)
	l.AppendLink(n2)
	l.AppendLink(n3)
	l.AppendLink(n4)
	l.AppendLink(n5)
	l.AppendLink(n6)
	l.AppendLink(n7)
	l.AppendLink(n8)
	l.AppendLink(n9)
	l.AppendLink(n10)
	// l.Front = n1
	// n1.next = n2
	// n2.next = n3
	// n3.next = n4
	// n4.next = n5
	// n5.next = n6
	// n6.next = n7
	// n7.next = n8
	// n8.next = n9
	// n9.next = n10

	fmt.Println(l)
	fmt.Println(l.Length)

	l.Remove(n3)
	l.AppendLink(&Node{Value: "11"})
	l.PrependLink(&Node{Value: "15"})
	fmt.Println(l)
	fmt.Println(l.Length)
	fmt.Println(l.FrontLink().Value)
	fmt.Println(l.BackLink().Value)
	l.Remove(n10)
	fmt.Println(l)
	fmt.Println(l.Length)
	fmt.Println(l.FrontLink().Value)
	fmt.Println(l.BackLink().Value)
	l.Remove(n1)
	l.InsertBefore(n3, n5)
	fmt.Println(l)
	/*
		如果还想在前面添加 3，则必须重新定义 Value 为 3 的节点
		因为 Node 采用的是传址
	*/
	nn := &Node{
		Value: "3",
	}
	l.InsertBefore(nn, n7)

	fmt.Println(l)
	fmt.Println(l.Length)
	fmt.Println(l.FrontLink().Value)
	fmt.Println(l.BackLink().Value)

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
