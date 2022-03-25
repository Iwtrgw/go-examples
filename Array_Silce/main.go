package main

import "sync"

// Array 可变长数组
type Array struct {
	array []int      // 固定大小数组，用满容量和满大小的切片来代替
	cap   int        // 真正容量
	len   int        // 长度
	lock  sync.Mutex // 并发安全锁
}

// Make 数组初始化
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large then cap")
	}
	// 把切片当数组用
	array := make([]int, cap, cap)
	// 元数据
	s.array = array
	s.len = 0
	s.cap = cap
	return s
}

// Append 数组添加一个元素
func (a *Array) Append(element int) {
	// 迸发锁
	a.lock.Lock()
	defer a.lock.Unlock()
	// 如果大小等于容量，则需要扩容
	if a.len == a.cap {
		// 没容量，数组扩容2倍
		newCap := a.cap * 2
		// 如果之前容量为0，则新容量为1
		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]int, newCap, newCap)
		for k, v := range a.array { // 把扩容前的数据放入新扩容的数组
			newArray[k] = v
		}
		// 替换数组
		a.array = newArray
		a.cap = newCap
	}
	// 把元素放入数组
	a.array[a.len] = element
	// 真实长度+1
	a.len += 1
}

// AppendMany 添加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

func main() {

}
