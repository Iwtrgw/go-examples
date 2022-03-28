package main

// Node 定义链表节点
type Node struct {
	Value int
	Next  *Node
}

// 初始化栈结构（空栈）
var size = 0
var stack = new(Node)

// Push 进栈
func Push(v int) bool {
	// 空栈的话直接将值放入头节点
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}
	// 否则将插入节点作为栈的头节点
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func main() {

}
