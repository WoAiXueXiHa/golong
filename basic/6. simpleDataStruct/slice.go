package main

import "fmt"

// 切片是对底层数组某一段的描述，不是数组本身
// 可以把切片理解成一个结构体
// type slice struct {
// 	ptr *T		// 指向底层数组
// 	lent int	// 当前长度
// 	cap int		// 容量
// }

// 切片不存数据，数据在底层数组
// 多个切片可能共享同一个底层数组
// append 可能复用旧数组，也可能创建新数组
// 切片传参会复制切片头，共享底层数组

// 1. 数组和切片关联
func correlation() {
	arr := [5]int{1,2,3,4,5}
	// [1,4)
	s := arr[1:4]

	fmt.Println(s)			// [2 3 4]
	fmt.Println(len(s))		// 3
	fmt.Println(cap(s))		// 4

	// 修改切片会影响底层数组
	nums := [3]int{1,2,3}
	p := nums[:]
	p[0] = 100
	fmt.Println(nums)	// [100 2 3]
}

// 2. 创建切片
func create() {
	// 字面量
	s := []int{1, 2, 3}
	// 要注意区别：[n]int是n个元素的数组 []int是切片
	fmt.Println(s)

	// make
	p := make([]int, 3)
	fmt.Println(s, len(s), cap(s))

	// 指定容量
	x := make([]int, 0, 10)
	fmt.Println(len(s), cap(s))
}

func main() {
	correlation()
}