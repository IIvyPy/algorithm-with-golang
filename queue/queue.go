package queue

//实现一个队列，用链表和数组都可以实现队列
type Queue interface {
	// 链表的长度
	Len() uint32
	// 链表的容量
	Cap() uint32
	////插入某一个值
	//InsertValue() 	error
	////删除某一个值
	//DeleteValue()   error
	// 首部
	Front() *Node
	// 尾部
	Rear() *Node
	// 入队列
	Enqueue(value interface{}) bool
	// 出队列
	Dequeue() bool
}

type Node struct {
	value     interface{}
	frontNode *Node
	rearNode  *Node
}

type myQueue struct {
	len   uint32
	cap   uint32
	front *Node
	rear  *Node
}

func (node *myQueue) Len() uint32 {
	return node.len
}

func (node *myQueue) Cap() uint32 {
	return node.cap
}

func (node *myQueue) Enqueue(value interface{}) bool {
	newNode := new(Node)
	node.rear.rearNode = newNode
	newNode.frontNode = node.rear
	node.rear = newNode
	return true
}
