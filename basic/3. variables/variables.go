package main

import "fmt"

// func main() {
// 	// // 显式声明类型
// 	// var age int = 20
// 	// // 自动推导类型
// 	// var name = "Vect"
// 	// // 短声明，只能在函数内部使用
// 	// hobby := "hiking"

// 	// 一次声明多个变量，挨个赋值
// 	var a1, a2, a3 int = 1, 2, 3
// 	fmt.Println(a1, a2, a3)

// 	// int8 int16 int32 int64
// 	// 这个和cpp的int32_t一样，很丝滑

// 	// 零值问题，只声明不赋值，会有零值
// 	var cnt int16
// 	var name string
// 	var ok bool
// 	fmt.Printf("count=%d\n", cnt)
// 	fmt.Printf("name=%q\n", name)  // %s对于空字符串啥也不显示 %q会显示""
// 	fmt.Printf("ok=%v\n", ok)
// }

// // 字符串相关
// func str() {
// 	s := "golong"

// 	fmt.Println("len: ", len(s))		// len是字节数，而不是字符数
// 	fmt.Println("first byte: ", s[0])	// 提取的是ASCII码

// 	for index, r := range s {
// 		fmt.Printf("index=%d rune=%c\n", index, r)
// 	}
// }



// func main() {
// 	str()
// }

func main() {
	// 显式类型转换，go不喜欢隐式类型转换
	var a int16 = 10
	var b int32 = 20
	fmt.Println(int32(a) + b)
}