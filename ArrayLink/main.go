package main

import "fmt"

// ArrayLink 数组实现链表
func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}
	// 定义一个5个节点的数组
	var array [5]Value
	array[0] = Value{"I", 3}
	array[1] = Value{"Army", 4}
	array[2] = Value{"You", 1}
	array[3] = Value{"Love", 2}
	// -1 表示已经是最后一个节点
	array[4] = Value{"!", -1}
	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}
}

func main() {
	array := [5]int64{}
	fmt.Println(array)

	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])

	ArrayLink()
}
