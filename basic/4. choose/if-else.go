package main 

import "fmt"

// 1. 最基础的if
func if1() {
	num := 100 

	if num <= 80 {
		fmt.Println("low")
	} else {
		fmt.Println("high")
	}
}

// 2. else if
func ageJudge() {
	age := 20

	if age >= 0 && age <= 18 {
		fmt.Println("未成年")
	} else if age > 18 && age <= 35 {
		fmt.Println("青年")
	} else if age > 35 && age <= 55 {
		fmt.Println("壮年")
	} else if age >= 55 && age <= 120 {
		fmt.Println("老年")
	} else {
		fmt.Println("年龄不正常")
	}
}

// 3. if前面可以带短句
// score 只在if-else的内部有效
func if2() {
	if score := 92; score >= 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("平常")
	}
}

func main() {
	if1()
	fmt.Println("-----------------")
	ageJudge()
	fmt.Println("-----------------")
	if2()
}