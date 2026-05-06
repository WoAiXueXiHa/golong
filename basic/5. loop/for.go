package main

import "fmt"

// go只有for循环，一切为了简洁，规范，统一

// 1. 基础写法
func for1() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
}

// 2. while写法
func for2() {
	retry := 0

	for retry < 3 {
		fmt.Println("retry", retry)
		retry++
	}
}

// 3. 死循环
// func for3() {
// 	for {
// 		fmt.Println("running")
// 	}
// }

// 4. range遍历
func for4() {
	// 注意这里，多行最后一个元素后需要逗号，单行不需要
	models := []string{
		"qwen",
		"openai",
		"claude",
		"deepseek",
	}

	for index, model := range models {
		fmt.Println(index, model)
	}

	fmt.Println("--------无索引版本---------")
	for _, model := range models {
		fmt.Println(model)
	}
}

func main() {
	for1()
	fmt.Println("-------------------")
	for2()
	fmt.Println("-------------------")
	//for3()
	for4()
}