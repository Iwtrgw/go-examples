package main

import (
	"fmt"
	"sync"
)

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

// Get 获取某个下标的元素
func (a *Array) Get(index int) int {
	// 越界
	if a.len == 0 || index > a.len {
		panic("index over len")
	}
	return a.array[index]
}

// Len 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// Print 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result += "]"
	return
}

// 测试
func main() {
	// 创建一个容量为3的数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 添加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加一个元素
	a.Append(11)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 添加多个元素
	a.AppendMany(4, 5)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
