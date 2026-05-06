package main

import "fmt"

// 1. 声明并打印变量
func print() {
	model := "qwen"
	port := 3030
	enabled := true
	fmt.Println(model, port, enabled)
}

// 2. 打印类型
func typePrint() {
	name := "hello"
	char := 'a'
	value := 20

	fmt.Printf("name=%T char=%c value=%d\n", name, char, value)
}

// 输出字符串长度
func strLen() {
	// unicode 一个汉字占3字节
	s := "Go语言"
	fmt.Println(len(s))
}

func main() {
	print()
	fmt.Println("-------------------")
	typePrint()
	fmt.Println("-------------------")
	strLen()
}