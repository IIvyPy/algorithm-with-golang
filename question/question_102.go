package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type queueList struct{
	list []*TreeNode
}

func (q *queueList) Pop() *TreeNode{
	r := q.list[0]
	q.list = q.list[1:]
	return r
}

func (q *queueList) Push(val *TreeNode){
	q.list = append(q.list, val)
}

func (q *queueList) Length() int{
	return len(q.list)
}

func NewQueueList() *queueList{
	l := make([]*TreeNode, 0)
	return &queueList{list: l}
}

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil{
		return res
	}
	q := NewQueueList()
	q.Push(root)

	for q.Length() != 0{
		size := q.Length()
		// 先取出孩子节点
		nodes := make([]int, 0)
		for size != 0{
			val := q.Pop()
			if val != nil{
				nodes = append(nodes, val.Val)
				if val.Left != nil{
					q.Push(val.Left)
				}
				if val.Right != nil{
					q.Push(val.Right)
				}
			}
			size--
		}
		res = append(res, nodes)
	}
	return res
}
