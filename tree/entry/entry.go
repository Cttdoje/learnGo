package main

import (
	"fmt"
	"learnGo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	// new 返回的是地址
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Traverse()
	fmt.Println()
	myNode := myTreeNode{&root}
	myNode.postOrder()

	//使用结构体方法
	/*root.print()
	/*nodes := []treeNode{
		{3,nil,&root},
		{value:5},
	}
	fmt.Println(nodes)*/
	//fmt.Println(root)
	//root.right.left.setValue(4)
	//root.right.left.print()

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(100)
	//pRoot.print()

	//切片中定义可以省略
}
