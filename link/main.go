package main

import (
	"fmt"
)

// LinkNode 单链表
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

// Ring 循环链表
type Ring struct {
	next, prev *Ring
	Value      interface{}
}

// 循环链表初始化，前置节点和后置节点都设为自己
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// New 创建N个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// Next 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// Move 获取第N个节点: 因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// Link 添加节点:往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// Unlink 删除节点后面的N个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len 链表长度获取
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}

// LinkNewTest 循环列表测试
func LinkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接5个新节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()
		// 如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	temp := r.Unlink(3) // 解除了后面两个节点

	// 打印原来的节点
	node := r
	for {
		// 打印节点的值
		fmt.Println(node.Value)
		// 移动到下一个节点
		node = node.Next()
		if node == r {
			break
		}
	}
	fmt.Println("_____________")
	// 打印被切断的节点
	node = temp
	for {
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == temp {
			break
		}
	}
}

func main() {
	// 新节点
	node := new(LinkNode)
	node.Data = 2

	// 新节点2
	node2 := new(LinkNode)
	node2.Data = 3
	node.NextNode = node2 // 将node2 链接到 node 节点上

	// 新节点3
	node3 := new(LinkNode)
	node3.Data = 4
	node2.NextNode = node3 // 将node3 链接到 node2 节点上

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.NextNode
			continue
		}
		// 如果下一个节点为nil,表示链表数据已取完
		break
	}

	//LinkNewTest()
	deleteTest()
}
