# Go 函数详解 —— C++ 转型指南

## 1. 什么是函数？

函数是一段 **完成特定任务的代码块**，可以被多次调用。

**C++ 对比：** Go 的函数和 C++ 函数本质上一样，但语法更简洁。

---

## 2. 基本函数定义

```go
package main

import "fmt"

// 标准函数定义
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("World")  // Hello, World
}
```

**C++ 对比：**

```cpp
#include <iostream>
#include <string>
void greet(const std::string& name) {
    std::cout << "Hello, " << name << std::endl;
}
```

**关键区别：**
- Go 参数类型写在参数名 **后面**
- Go 函数定义不需要分号结尾
- Go 没有 `const` 关键字（通过类型约束）

---

## 3. 函数返回值

### 3.1 单返回值

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(1, 2)
    fmt.Println(result)  // 3
}
```

### 3.2 多返回值（Go 特有！）

这是 Go 最强大的特性之一，C++ 需要用结构体或输出参数才能实现。

```go
func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    x, y := swap(1, 2)
    fmt.Println(x, y)  // 2 1
}
```

**C++ 对比（模拟多返回值）：**

```cpp
// C++ 方式1：用输出参数
void swap(int a, int b, int& outA, int& outB) {
    outA = b;
    outB = a;
}

// C++ 方式2：用 std::pair
std::pair<int, int> swap(int a, int b) {
    return {b, a};
}

// C++ 方式3：用 std::tuple
std::tuple<int, int> swap(int a, int b) {
    return std::make_tuple(b, a);
}
```

Go 的多返回值直接返回，语法更清晰。

---

## 4. 命名返回值（Named Return Values）

Go 允许给返回值起名字，函数内可以直接使用。

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 裸 return，返回 x 和 y
}

func main() {
    fmt.Println(split(10))  // 4 6
}
```

**注意：** 命名返回值使得 `return` 可以不写返回值（裸 return），但一般不建议在长函数中使用，因为可读性差。

---

## 5. 错误处理：多返回值的经典用法

Go 没有异常机制，错误通过 **多返回值** 传递，这是 Go 的核心设计哲学。

```go
import "strconv"

func parseNumber(s string) (int, error) {
    num, err := strconv.Atoi(s)  // 字符串转整数
    if err != nil {
        return 0, err  // 返回零值和错误
    }
    return num, nil    // 返回结果和 nil（无错误）
}

func main() {
    if num, err := parseNumber("42"); err == nil {
        fmt.Println("Number:", num)
    } else {
        fmt.Println("Error:", err)
    }
}
```

**C++ 对比：** C++ 通常用 `try-catch` 异常机制或返回错误码。

---

## 6. 可变参数函数（Variadic Functions）

可变参数函数可以接受 **零个或多个** 同类型的参数。

### 6.1 基本用法

```go
// nums 是 slice []int
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
    fmt.Println(sum())              // 0
}
```

### 6.2 传递 slice

```go
func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(nums...))  // 15，展开 slice
}
```

**C++ 对比：**

```cpp
// C++ 可变参数模板
template<typename... Args>
int sum(Args... args) {
    return (args + ...);  // C++17 折叠表达式
}

// 或者用 initializer_list
int sum(std::initializer_list<int> nums) {
    int total = 0;
    for (int n : nums) total += n;
    return total;
}
```

---

## 7. 闭包（Closure）

函数可以定义在其他函数内部，捕获外部变量。

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos := adder()
    fmt.Println(pos(1))  // 1
    fmt.Println(pos(2))  // 3
    fmt.Println(pos(3))  // 6
}
```

**C++ 对比：** C++ 用 lambda 表达式实现类似功能。

```cpp
auto adder() {
    int sum = 0;
    return [sum](int x) mutable {
        sum += x;
        return sum;
    };
}
```

---

## 8. 递归（Recursion）

函数可以调用自己。

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5))  // 120
}
```

---

## 9. 函数类型与变量

在 Go 中，函数是一等公民（first-class citizen），可以像变量一样传递。

```go
func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    // 函数可以赋值给变量
    var op func(int, int) int
    op = add
    fmt.Println(op(3, 4))  // 7

    op = multiply
    fmt.Println(op(3, 4))  // 12
}
```

### 9.1 作为参数传递

```go
func apply(op func(int, int) int, a, b int) int {
    return op(a, b)
}

func main() {
    result := apply(add, 3, 4)
    fmt.Println(result)  // 7
}
```

### 9.2 作为返回值

```go
func getOperator(op string) func(int, int) int {
    switch op {
    case "+":
        return add
    case "*":
        return multiply
    default:
        return nil
    }
}
```

---

## 10. defer 延迟执行

`defer` 关键字用于延迟函数调用，常用于资源清理。

```go
func readFile(filename string) {
    // defer 确保函数退出时执行
    defer fmt.Println("Cleanup: done")

    fmt.Println("Reading file:", filename)
    return  // defer 会在 return 之前执行
}

func main() {
    readFile("test.txt")
    // 输出:
    // Reading file: test.txt
    // Cleanup: done
}
```

### defer 多个调用（栈结构，后进先出）

```go
func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")

    fmt.Println("main")
    // 输出:
    // main
    // third
    // second
    // first
}
```

**C++ 对比：** 类似 RAII（Resource Acquisition Is Initialization）模式，但更简单。

---

## 11. 日常开发中的函数设计建议

### 11.1 函数尽量短小，职责单一

Go 社区非常重视代码可读性。一个函数最好只做一件事，名字直接表达意图。

```go
// 不推荐：函数里同时做解析、校验、保存、打印
func handleUser(input string) error {
    // ...
    return nil
}

// 推荐：拆成多个职责清晰的小函数
func parseUser(input string) (User, error) { ... }
func validateUser(user User) error { ... }
func saveUser(user User) error { ... }
```

**日常经验：** 如果一个函数需要频繁滚动屏幕才能看完，通常就该考虑拆分。

### 11.2 错误要尽早返回，减少嵌套

Go 常见风格是 **提前返回错误**，避免多层 `if-else` 嵌套。

```go
func process(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    return handleData(data)
}
```

不要写成：

```go
func process(path string) error {
    file, err := os.Open(path)
    if err == nil {
        defer file.Close()
        data, err := io.ReadAll(file)
        if err == nil {
            return handleData(data)
        }
    }
    return err
}
```

### 11.3 返回值顺序约定：结果在前，error 最后

Go 约定俗成的写法是：

```go
func findUser(id int64) (User, error) {
    // ...
}
```

不要反过来写：

```go
func findUser(id int64) (error, User) { // 不符合 Go 习惯
    // ...
}
```

### 11.4 不要滥用命名返回值和裸 return

命名返回值适合非常短的函数，或者需要在 `defer` 中修改返回值的场景。

```go
func area(width, height int) (result int) {
    result = width * height
    return
}
```

但在业务代码里，长函数使用裸 `return` 容易让人不知道到底返回了什么。

```go
// 不推荐：函数很长时，裸 return 可读性差
func buildConfig() (cfg Config, err error) {
    // ... 很多逻辑
    return
}

// 推荐：明确返回
func buildConfig() (Config, error) {
    cfg := Config{}
    // ... 很多逻辑
    return cfg, nil
}
```

### 11.5 函数参数过多时，考虑使用结构体

如果函数参数超过 3～4 个，并且经常一起传递，建议封装成结构体。

```go
// 不推荐：参数太多，调用时难以理解每个值的含义
func createUser(name string, age int, email string, city string, active bool) error {
    // ...
}

// 推荐：用结构体表达参数含义
type CreateUserRequest struct {
    Name   string
    Age    int
    Email  string
    City   string
    Active bool
}

func createUser(req CreateUserRequest) error {
    // ...
}
```

这样后续新增字段也更方便，不容易破坏调用方代码。

### 11.6 参数传值还是传指针？

Go 默认是值传递。是否使用指针，主要看三个因素：

| 场景 | 建议 |
|------|------|
| 需要修改调用方对象 | 使用指针 |
| 结构体很大，拷贝成本高 | 使用指针 |
| 表示可选值或可能为空 | 可以使用指针 |
| 小结构体、基本类型、只读数据 | 通常直接传值 |

```go
type User struct {
    Name string
    Age  int
}

func rename(user *User, name string) {
    user.Name = name
}

func display(user User) {
    fmt.Println(user.Name, user.Age)
}
```

**注意：** 不要为了“看起来高效”到处用指针。指针会增加共享状态和空指针风险。

### 11.7 defer 的参数会立即求值

`defer` 延迟的是函数调用的执行，但参数会在执行到 `defer` 这一行时立即计算。

```go
func main() {
    x := 1
    defer fmt.Println(x)
    x = 2
}
```

输出：

```text
1
```

如果想在最后读取最新值，可以使用闭包：

```go
func main() {
    x := 1
    defer func() {
        fmt.Println(x)
    }()
    x = 2
}
```

输出：

```text
2
```

### 11.8 defer 常用于资源释放，但循环里要小心

```go
func readFiles(paths []string) error {
    for _, path := range paths {
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close() // 注意：会等 readFiles 结束才关闭

        // ...
    }
    return nil
}
```

如果循环很多次，文件会一直不关闭，可能耗尽文件句柄。可以拆成小函数：

```go
func readFiles(paths []string) error {
    for _, path := range paths {
        if err := readOneFile(path); err != nil {
            return err
        }
    }
    return nil
}

func readOneFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    // ...
    return nil
}
```

### 11.9 闭包捕获循环变量要注意

在并发或保存函数时，闭包捕获循环变量容易写出错误代码。

```go
func main() {
    names := []string{"a", "b", "c"}

    for _, name := range names {
        go func() {
            fmt.Println(name)
        }()
    }
}
```

在较新的 Go 版本中，`range` 循环变量的行为已经改进，但日常开发中仍建议显式传参，兼容老代码也更清晰：

```go
for _, name := range names {
    go func(name string) {
        fmt.Println(name)
    }(name)
}
```

### 11.10 匿名函数立即执行

Go 中可以定义匿名函数并立即调用，常用于局部封装一段逻辑。

```go
result := func(a, b int) int {
    return a + b
}(1, 2)

fmt.Println(result) // 3
```

但不要为了“炫技”滥用。普通逻辑直接写更清楚。

### 11.11 函数没有默认参数，也不支持重载

Go 不支持 C++ 那样的默认参数和函数重载。

```cpp
// C++ 支持
void connect(std::string host, int port = 80);
void print(int x);
void print(std::string s);
```

Go 通常有几种替代方案：

```go
// 方式1：提供不同名字的函数
func Connect(host string, port int) error { ... }
func ConnectDefault(host string) error {
    return Connect(host, 80)
}

// 方式2：使用配置结构体
type ConnectOptions struct {
    Host string
    Port int
}

func ConnectWithOptions(opts ConnectOptions) error {
    if opts.Port == 0 {
        opts.Port = 80
    }
    // ...
    return nil
}
```

### 11.12 方法也是函数，只是多了接收者

Go 没有 C++ 的 class，但可以给类型定义方法。

```go
type Counter struct {
    n int
}

func (c *Counter) Add(delta int) {
    c.n += delta
}

func (c Counter) Value() int {
    return c.n
}
```

这里：

- `func (c *Counter) Add(...)` 是指针接收者，可以修改原对象
- `func (c Counter) Value()` 是值接收者，拿到的是副本

日常建议：如果一个类型的方法里有指针接收者，其他方法通常也统一使用指针接收者，减少混乱。

---

## 12. 常见函数踩坑

### 坑1：忽略 error

```go
// 不推荐
value, _ := strconv.Atoi("abc")
fmt.Println(value)
```

除非你非常确定错误可以忽略，否则应该处理：

```go
value, err := strconv.Atoi("abc")
if err != nil {
    return err
}
fmt.Println(value)
```

### 坑2：返回局部变量指针是安全的，但不要滥用

Go 允许返回局部变量地址，编译器会做逃逸分析。

```go
func newUser(name string) *User {
    user := User{Name: name}
    return &user
}
```

这在 Go 中是安全的，不像 C/C++ 返回栈上局部变量指针那样危险。但如果不需要共享或修改，直接返回值也很好。

### 坑3：函数变量可能是 nil

```go
var fn func()
// fn() // panic: runtime error

if fn != nil {
    fn()
}
```

从 map、配置、回调中取函数时尤其要注意。

### 坑4：defer 中修改返回值要谨慎

```go
func f() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("wrap: %w", err)
        }
    }()

    return doSomething()
}
```

这种写法可以用，但依赖命名返回值。团队协作中要确保大家都能读懂，否则宁可显式处理错误。

---

## 13. 完整示例：多返回值 + 错误处理

```go
package main

import (
    "fmt"
    "strconv"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func parseAndDivide(a, b string) {
    x, err := strconv.ParseFloat(a, 64)
    if err != nil {
        fmt.Printf("Invalid number: %s\n", a)
        return
    }

    y, err := strconv.ParseFloat(b, 64)
    if err != nil {
        fmt.Printf("Invalid number: %s\n", b)
        return
    }

    result, err := divide(x, y)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("%s / %s = %.2f\n", a, b, result)
}

func main() {
    parseAndDivide("10", "2")   // 10 / 2 = 5.00
    parseAndDivide("10", "0")  // Error: division by zero
    parseAndDivide("abc", "2")  // Invalid number: abc
}
```

---

## 14. Go 函数 vs C++ 函数关键区别

| 特性 | Go | C++ |
|------|-----|-----|
| 多返回值 | ✅ 原生支持 | ❌ 需要 pair/tuple/输出参数 |
| 语法 | 类型后置 `func f(a int)` | 类型前置 `void f(int a)` |
| 默认参数 | ❌ 不支持 | ✅ 支持 |
| 函数重载 | ❌ 不支持 | ✅ 支持 |
| 闭包 | ✅ 匿名函数捕获外部变量 | ✅ lambda 表达式 |
| defer | ✅ Go 独有 | ❌ 无 |
| 尾递归优化 | ❌ 不保证 | ✅ 通常支持 |

---

## 15. 面试/开发常用速查

```go
// 基本函数
func add(a, b int) int { return a + b }

// 多返回值
func divide(a, b int) (int, int) { return a / b, a % b }

// 多返回值 + 错误
func parse(s string) (int, error) { ... }

// 可变参数
func sum(nums ...int) int { ... }

// defer 清理
func read() {
    defer close()  // 退出前执行
    ...
}
```
