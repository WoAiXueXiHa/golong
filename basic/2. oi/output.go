package main

import "fmt"

func main() {
	// Println会自动换行，多个输出会带空格
	fmt.Println("hello")
	fmt.Println("go", "backend", 8080)

	// 格式化输出，和C一样写法
	// 占位符：%s字符串，%d整数，%t布尔值， %f浮点数
	// %v通用占位符，会根据值的类型选择合适的占位符
	// %T打印类型
	model := "Openai"
	qps := 200
	enable := true

	fmt.Printf("model: %s\nqps: %d\nenable: %t\n", model, qps, enable)

	// 生成字符串，不直接打印，返回一个字符串
	service := "ai-backend"
	port := 8090

	msg := fmt.Sprintf("service: %s, port: %d", service, port)
	fmt.Println(msg)
}