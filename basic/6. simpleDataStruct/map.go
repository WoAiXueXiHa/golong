package main

import "fmt"

func test1() {
	// map[keyType]valueType
	// 1. make 初始化
	m1 := make(map[string]int)

	// 2. 字面量 初始化就有数据
	m2 := map[string]int{
		"apple": 10,
		"banana": 2,
	}

	// 3. 空 map 不能直接复制
	var m3 map[string]int

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}

func test2() {
	m := make(map[string]int)

	// 赋值
	m["vect"] = 666
	m["kunkun"] = 2

	// 读取，如果key不存在，返回0值
	fmt.Println(m["apple"])
	fmt.Println(m["grape"])

	// 删除，如果key不存在，删除操作无效，不会报错
	delete(m, "grape")

	// 检查key是否存在
	// 双返回值 (value,ok)判断存在
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["grape"]
	fmt.Println(v2, ok2)
}

func main() {
	test1()
	fmt.Println("--------------------")
	test2()
}