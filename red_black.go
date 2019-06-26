package main

import "fmt"

const (
	// RED 红树设为true
	RED bool = true
	// BLACK 黑树设为false
	BLACK bool = false

	// 左旋
	LEFTROTATE bool = true
	// 右旋
	RIGHTROTATE bool = false
)


// RBNode 红黑树
type RBNode struct {
	value               int64
	color               bool
	left, right, parent *RBNode
}

// getGrandParent() 获取父级节点的父级节点
func (rbnode *RBNode) getGrandParent() *RBNode {
	var a *RBNode
	return  a
}

// getSibling() 获取兄弟节点
func (rbnode *RBNode) getSibling() *RBNode {
	var a *RBNode
	return  a
}

// GetUncle() 父节点的兄弟节点
func (rbnode *RBNode) getUncle() *RBNode {
	var a *RBNode
	return  a
}

// 树的结构只包含一个根节点Root
type RBTree struct {
	root *RBNode
}





func main(){
	fmt.Println(123)
}

