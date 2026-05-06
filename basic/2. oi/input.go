package main 

import (
	"fmt"
	"os"
)

func main() {
	// var name string
	// var age int

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)
	// fmt.Printf("name: %s, age: %d\n", name, age)

	// 一般很少使用输入，而是读配置信息

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("port =", port)
}