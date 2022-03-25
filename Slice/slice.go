package main

import "fmt"

func main() {
	// 创建一个容量为2的切片
	array := make([]int, 0, 2)
	fmt.Println("cap:", cap(array), "len:", len(array), "array:", array)

	// 添加，虽然 append 但是没有赋予原来的变量 array
	_ = append(array, 1)
	fmt.Println("cap:", cap(array), "len:", len(array), "array:", array)

	fmt.Println("___________")
	// 添加后赋予原来的变量
	array = append(array, 1)
	fmt.Println("cap:", cap(array), "len:", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1)
	fmt.Println("cap:", cap(array), "len:", len(array), "array:", array)
	// 我们可以看到 Golang 的切片无法原地 append，每次添加元素时返回新的引用地址，
	//必须把该引用重新赋予之前的切片变量。并且，当容量不够时，会自动倍数递增扩容。
	//事实上，Golang 在切片长度大于 1024 后，会以接近于 1.25 倍进行容量扩容
}
