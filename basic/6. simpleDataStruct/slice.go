package main

import "fmt"

func slice1() {
	// 变长数组
	// []里面没有数字，长度不固定了！
	// [qwen, openai, ...] 变长数组
	models := []string{"qwen", "openai"}
	models = append(models, "llama")

	fmt.Println(models)
	fmt.Println("len ", len(models))
}

func slice2() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) 
	fmt.Println("emp:", s, "len: ", len(s), "cap:", cap(s))
}

func slice3() {
	t := make([][]int, 3)

	for i := 0; i < 3; i++ {
		rowLen := i + 1
		t[i] = make([]int, rowLen)

		for j := 0; j < rowLen; j++ {
			t[i][j] = i + j
		}
	}

	fmt.Println("二维切片: ", t)
}

func slice4() {
	// 切片操作
	s := "hello, golong"
	str := make([]byte, len(s))
	copy(str, s)
	fmt.Println("copy: ", str)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)
}

func main() {
	slice1()
	fmt.Println("-------------------")
	slice2()
	fmt.Println("-------------------")
	slice3()
	fmt.Println("-------------------")
	slice4()
}