package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//给结构体定义方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//go语言函数属于值传递
func (node *Node) SetValue(value int) {
	node.Value = value
}

// 工厂方法创建对象
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
