package main

import (
	"fmt"
	"leetcode/base/tool"
)

func main() {
	t := New(1)
	t.Left = New(3)
	t.Right = New(6)
	t.Left.Left = New(4)
	t.Left.Right = New(5)
	t.Left.Left.Left = New(7)

	t.PostOrderTest()
}

//BTree 二叉树
type BTree struct {
	Data  interface{}
	Left  *BTree
	Right *BTree
}

//New 创建新tree
func New(data interface{}) *BTree {
	return &BTree{Data: data}
}

//PreOrderRecursion 前序递归调用
func (root *BTree) PreOrderRecursion() {
	if root != nil {
		root.Left.PreOrderRecursion()
		fmt.Printf("%v  ", root.Data)
		root.Right.PreOrderRecursion()
	}
}

//PreOrderNoRecursion 前序,非递归
func (root *BTree) PreOrderNoRecursion() {
	bt := root
	stack := tool.NewStack()
	for bt != nil || stack.Size > 0 {
		for bt != nil {
			fmt.Printf("%v  ", bt.Data)
			stack.Push(bt)
			bt = bt.Left
		}
		if stack.Size > 0 {
			pop := stack.Pop()
			bt = pop.(*BTree).Right
		}
	}
}

//MidOrderNoRecursion 中序,非递归
func (root *BTree) MidOrderNoRecursion() {
	bt := root
	stack := tool.NewStack()
	for bt != nil || stack.Size > 0 {
		for bt != nil {
			stack.Push(bt)
			bt = bt.Left
		}
		if stack.Size > 0 {
			pop := stack.Pop()
			fmt.Printf("%v  ", pop.(*BTree).Data)
			bt = pop.(*BTree).Right
		}
	}
}

//PostOrderNoRecursion 后序,非递归
func (root *BTree) PostOrderNoRecursion() {
	bt := root
	stack := tool.NewStack()

	var lastVisitNode *BTree

	for bt != nil || stack.Size > 0 {
		for bt != nil {
			stack.Push(bt)
			bt = bt.Left
		}

		if stack.Size > 0 {
			top := stack.Top()
			if top.(*BTree).Right != nil && top.(*BTree).Right != lastVisitNode {
				bt = top.(*BTree).Right
			} else {
				fmt.Printf("%v  ", top.(*BTree).Data)
				lastVisitNode = top.(*BTree)
				stack.Pop()
			}
		}
	}
}

//LevelOrder 层次遍历
func (root *BTree) LevelOrder() {
	queue := make([]interface{}, 0)
	bt := root
	if bt != nil {
		queue = append(queue, bt)
	}

	for len(queue) != 0 {
		first := queue[0]
		fmt.Printf("%v  ", first.(*BTree).Data)
		if first.(*BTree).Left != nil {
			queue = append(queue, first.(*BTree).Left)
		}
		if first.(*BTree).Right != nil {
			queue = append(queue, first.(*BTree).Right)
		}

		queue = queue[1:]
	}
}

//PreOrderTest Pre
func (root *BTree) PreOrderTest() {
	stack := tool.NewStack()
	bt := root
	for bt != nil || stack.Size > 0 {
		for bt != nil {
			fmt.Printf("%v  ", bt.Data)
			stack.Push(bt)
			bt = bt.Left
		}
		if stack.Size > 0 {
			pop := stack.Pop()
			bt = pop.(*BTree).Right
		}
	}
}

//MidOrderTest Pre
func (root *BTree) MidOrderTest() {
	stack := tool.NewStack()
	bt := root
	for bt != nil || stack.Size > 0 {
		for bt != nil {
			stack.Push(bt)
			bt = bt.Left
		}
		if stack.Size > 0 {
			pop := stack.Pop()
			fmt.Printf("%v  ", pop.(*BTree).Data)
			bt = pop.(*BTree).Right
		}
	}
}

//PostOrderTest Post
func (root *BTree) PostOrderTest() {
	stack := tool.NewStack()
	bt := root
	var lastVisitNode *BTree

	for bt != nil || stack.Size > 0 {
		for bt != nil {
			stack.Push(bt)
			bt = bt.Left
		}

		if stack.Size > 0 {
			top := stack.Top().(*BTree)
			if top.Right != nil && top.Right != lastVisitNode {
				stack.Push(top.Right)
			} else {
				fmt.Printf("%v  ", top.Data)
				lastVisitNode = top
				stack.Pop()
			}
		}
	}
}
