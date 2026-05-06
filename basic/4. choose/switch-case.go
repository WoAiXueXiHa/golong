package main

import "fmt"

// go的switch-case比C++的安全，不会向下穿透，每个case默认有break

// 1. 正常语句
func switch1() {
	status := "running"

	switch status {
	case "queued":
		fmt.Println("waiting")
	case "running":
		fmt.Println("processing")
	case "done":
		fmt.Println("finished")
	default:
		fmt.Println("unknown")
	}
}

// 2. 多个case合并
func switch2() {
	status := "created"

	switch status {
	case "created", "queued" :
		fmt.Println("not started")
	case "running" :
		fmt.Println("running")
	case "done", "failed" :
		fmt.Println("finished")
	}
}

func main() {
	switch1()
	fmt.Println("-----------------")
	switch2()
}