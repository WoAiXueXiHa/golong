package main

import "fmt"

// go中的数组是 长度固定、元素类型相同、连续存储的一组值

// 1. 声明
func decl() {
	// 默认零值填充
	fmt.Println("nums:")
	var nums [3]int
	fmt.Println(nums)
	// 声明并初始化
	arr := [5]string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}
	fmt.Println("arr:")
	fmt.Println(arr)
	// 自动推导长度
	num := [...]int{10,20}
	fmt.Println("num:")
	fmt.Println(num)

	// 指定索引初始化
	array := [5]int{0: 10, 3: 40}
	fmt.Println("array: ")
	fmt.Println(array)

	// 注意：长度是数组的一部分
	// a := [3]int
	// b := [4]int
	// fmt.Println(a, b)
	// 编译错误，[3]int [4]int 是不同类型
	// go数组不会退化为指针
}

// 2. 遍历数组
func traverse() {
	nums := [3]int{10, 20, 30}
	// 2.1. 普通for循环
	fmt.Println("for遍历: ")
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	// 2.2. range遍历
	fmt.Println("range遍历: ")
	for i, v := range nums {
		fmt.Println(i, v)
	} 

	fmt.Println("不要索引，只要值: ")
	// 2.3. 不要索引，只要值
	for _, v := range nums {
		fmt.Println(v)
	}

	// 2.4. 不要值，只要索引
	fmt.Println("不要值，只要索引: ")
	for i,_ := range nums {
		fmt.Println(i)
	}

	// 2.5. range遍历数组的本质：
	// 对数组使用range时，go会复制数组再遍历
	fmt.Println("对数组使用range时，go会复制数组再遍历: ")
	arr := [3]int{0, 1, 2}
	// 遍历还是来自开始时的数组拷贝
	for i, v := range arr {
		arr[0] = 100
		fmt.Println(i, v)
	}
	// 遍历结束，索引0处被修改
	fmt.Println(arr)

	// 如果数组很大，使用range拷贝代价太大了，遍历数组指针或者切片
	// 1> 指针
	fmt.Println("指针: ")
	for i, v := range &nums{
		fmt.Println(i, v)
	}
	// 2> 切片，复制切片头，不复制底层数组
	s := arr[:]
	fmt.Println("切片: ")
	for i, v := range s {
		fmt.Println(i, v)
	}
}

// 3. 数组可以比较
func comp() {
	// 必须要长度相等才能比较
	a := [3]int{1,2,3}
	b := [3]int{1,2,3}
	c := [3]int{1,2,4}

	fmt.Println(a == b)
	fmt.Println(a == c)

	// 长度不同不能比较，因为类型不同

	// 前提：数组元素类型可以比较，数组可以作为 map 的 key
	m := map[[2]int]string{
		[2]int{1,2} : "point A",
	}
	fmt.Println(m[[2]int{1,2}])
}



func main() {
	// decl()
	// traverse()
	// comp()
}