// 可执行程序包，以main开头，就是main函数
package main
// 相当于#include头文件，引入fmt包，用于打印输出
import "fmt"
// Go中首字母大写的名字可以被包外访问
func main() {
	fmt.Println("hello world")
}