package binaryTree

import (
	"sync"
	"fmt"
)

type treesFun interface {
	Insert(int)
}

type treeNode struct{
	value     	int
	leftNode    *treeNode
	rightNode   *treeNode
}

type myTree struct{
	root    *treeNode
	lock    sync.Mutex
	treesFun
}

func InsertTree(treeNode *treeNode, newTree *treeNode){
	if newTree.value > treeNode.value{
		if treeNode.rightNode != nil{
			InsertTree(treeNode.rightNode, newTree)
		}else{
			treeNode.rightNode = newTree
		}
	}else{
		if treeNode.leftNode != nil{
			InsertTree(treeNode.leftNode, newTree)
		}else{
			treeNode.leftNode = newTree
		}
	}
}

func (node *myTree)Insert(value int){
	node.lock.Lock()
	defer node.lock.Unlock()

	newNode := new(treeNode)
	newNode.value = value

	if node.root == nil{
		node.root = newNode
	}else{
		InsertTree(node.root, newNode)
	}
}

func BinaSearch(tree *treeNode, value int) (bool){
	fmt.Println(tree.value, value)
	if tree.value > value {
		if tree.rightNode != nil {
			return BinaSearch(tree.rightNode, value)
		} else {
			return false
		}
	}else if tree.value < value{
		if tree.rightNode != nil{
			return BinaSearch(tree.rightNode, value)
		}else{
			return false
		}
	}else{
		return true
	}
}

func PreorderTraveser(node *treeNode){
	if node != nil {
		fmt.Println(node.value)
		PreorderTraveser(node.leftNode)
		PreorderTraveser(node.rightNode)
	}
}

func InorderTraverser(node *treeNode){
	fmt.Println("xxxx", node)
	if node != nil{
		InorderTraverser(node.leftNode)
		fmt.Println(node.value)
		InorderTraverser(node.rightNode)
	}
}
