# Go Array 数组详解 —— 直击本质、日常开发与面试

## 1. 一句话理解数组

Go 数组是 **长度固定、元素类型相同、连续存储的一组值**。

```go
var nums [3]int = [3]int{10, 20, 30}
fmt.Println(nums[0]) // 10
```
  
数组的本质重点：

- 长度是类型的一部分
- 数组是值类型
- 赋值和传参都会复制整个数组
- 实际开发中数组用得不多，更多使用切片 `slice`

---

## 2. 数组的声明和初始化

### 2.1 指定长度

```go
var nums [3]int
fmt.Println(nums) // [0 0 0]
```

数组声明后会自动初始化为元素类型的零值。

| 类型 | 零值 |
|------|------|
| `int` | `0` |
| `string` | `""` |
| `bool` | `false` |
| 指针 | `nil` |

### 2.2 声明时赋值

```go
nums := [3]int{10, 20, 30}
fmt.Println(nums) // [10 20 30]
```

### 2.3 让编译器自动推导长度

```go
nums := [...]int{10, 20, 30}
fmt.Println(len(nums)) // 3
```

注意：`[...]int` 仍然是数组，不是切片。

### 2.4 指定下标初始化

```go
nums := [5]int{0: 10, 3: 40}
fmt.Println(nums) // [10 0 0 40 0]
```

这个写法适合稀疏初始化。

---

## 3. 长度是数组类型的一部分

这是 Go 数组最关键的地方。

```go
var a [3]int
var b [4]int

fmt.Println(a, b)
// a = b // 编译错误：[3]int 和 [4]int 是不同类型
```

`[3]int` 和 `[4]int` 是两个完全不同的类型。

**C++ 对比：** C++ 中 `int a[3]` 和 `int b[4]` 也是不同长度的数组，但 C++ 数组经常退化为指针；Go 数组不会自动退化为指针。

---

## 4. 数组是值类型

数组赋值会复制整个数组。

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

这点和切片不同。

```go
s1 := []int{1, 2, 3}
s2 := s1

s2[0] = 100

fmt.Println(s1) // [100 2 3]
fmt.Println(s2) // [100 2 3]
```

数组复制的是所有元素，切片复制的是切片头，底层数组仍然共享。

---

## 5. 数组作为函数参数

### 5.1 直接传数组：会复制

```go
func modify(nums [3]int) {
    nums[0] = 100
}

func main() {
    nums := [3]int{1, 2, 3}
    modify(nums)
    fmt.Println(nums) // [1 2 3]
}
```

函数里的 `nums` 是外部数组的副本。

### 5.2 传数组指针：可以修改原数组

```go
func modify(nums *[3]int) {
    nums[0] = 100
}

func main() {
    nums := [3]int{1, 2, 3}
    modify(&nums)
    fmt.Println(nums) // [100 2 3]
}
```

Go 允许对数组指针直接使用下标：

```go
nums[0] = 100
```

不需要写成：

```go
(*nums)[0] = 100
```

虽然两者本质等价。

### 5.3 日常建议

如果函数要处理一组不固定长度的数据，优先使用切片：

```go
func sum(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

数组参数只适合长度本身有明确意义的场景，例如：

```go
func handleIPv4(ip [4]byte) {}
func handleMD5(hash [16]byte) {}
func handleSHA256(hash [32]byte) {}
```

---

## 6. 遍历数组

### 6.1 普通 for

```go
nums := [3]int{10, 20, 30}

for i := 0; i < len(nums); i++ {
    fmt.Println(i, nums[i])
}
```

### 6.2 range 遍历

```go
nums := [3]int{10, 20, 30}

for i, v := range nums {
    fmt.Println(i, v)
}
```

只要值，不要下标：

```go
for _, v := range nums {
    fmt.Println(v)
}
```

只要下标，不要值：

```go
for i := range nums {
    fmt.Println(i)
}
```

---

## 7. range 遍历数组的本质

对数组使用 `range` 时，Go 会复制数组再遍历。

```go
nums := [3]int{1, 2, 3}

for i, v := range nums {
    nums[0] = 100
    fmt.Println(i, v)
}

fmt.Println(nums)
```

输出中 `v` 仍然来自遍历开始时的数组副本。

如果数组很大，直接 range 数组可能产生复制成本。可以遍历数组指针或切片：

```go
for i, v := range &nums {
    fmt.Println(i, v)
}
```

或者：

```go
s := nums[:]
for i, v := range s {
    fmt.Println(i, v)
}
```

面试常问点：

- `range` 数组：会复制数组
- `range` 数组指针：不会复制数组
- `range` 切片：复制切片头，不复制底层数组

---

## 8. 数组可以比较

只要元素类型可比较，数组就可以使用 `==` 和 `!=`。

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{1, 2, 4}

fmt.Println(a == b) // true
fmt.Println(a == c) // false
```

注意长度不同不能比较，因为类型不同。

```go
a := [3]int{1, 2, 3}
b := [4]int{1, 2, 3, 4}

fmt.Println(a)
fmt.Println(b)
// fmt.Println(a == b) // 编译错误
```

数组可比较，所以数组可以作为 map 的 key。

```go
m := map[[2]int]string{
    [2]int{1, 2}: "point A",
}

fmt.Println(m[[2]int{1, 2}]) // point A
```

但 slice 不可以作为 map key。

---

## 9. 数组和切片的转换

数组可以通过切片表达式得到切片。

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]

fmt.Println(s) // [2 3 4]
```

得到的切片和原数组共享底层数据。

```go
arr := [3]int{1, 2, 3}
s := arr[:]

s[0] = 100
fmt.Println(arr) // [100 2 3]
```

这也是理解切片的入口：**切片本身不是数组，它只是对底层数组某一段的描述。**

---

## 10. 多维数组

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix[0][1]) // 2
```

多维数组的长度也属于类型的一部分。

```go
var a [2][3]int
var b [3][2]int

fmt.Println(a, b)
// a = b // 编译错误
```

日常开发中，二维数据更常用二维切片：

```go
grid := [][]int{
    {1, 2, 3},
    {4, 5},
}
```

二维切片每一行长度可以不同；二维数组每一行长度固定。

---

## 11. 数组的日常使用场景

虽然 Go 日常开发中切片更常见，但数组并不是没用。

### 11.1 固定长度的二进制数据

```go
var ip [4]byte
var md5 [16]byte
var sha256 [32]byte
```

长度本身就是数据含义的一部分。

### 11.2 小型固定查表

```go
weekdays := [...]string{
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
}

fmt.Println(weekdays[1]) // Monday
```

### 11.3 作为 map key

```go
type Point [2]int

visited := map[Point]bool{}
visited[Point{1, 2}] = true
```

这在网格搜索、坐标去重、算法题里很常见。

### 11.4 避免堆分配的小缓冲区

某些性能敏感代码会用固定数组作为临时缓冲区，然后切成 slice 使用。

```go
var buf [1024]byte
s := buf[:0]

s = append(s, 'h', 'i')
fmt.Println(string(s))
```

这类写法常见于底层库或性能优化场景，初学阶段知道即可。

---

## 12. 数组常见坑

### 坑1：把数组误以为是切片

```go
arr := [3]int{1, 2, 3}
// arr = append(arr, 4) // 编译错误
```

`append` 只能用于切片，不能用于数组。

正确写法：

```go
s := []int{1, 2, 3}
s = append(s, 4)
```

### 坑2：数组传参会复制

```go
func f(a [1000000]int) {
    // 会复制一个很大的数组
}
```

大数组不要直接传值，使用切片或数组指针。

```go
func f(a []int) {}
func g(a *[1000000]int) {}
```

### 坑3：`[...]int` 不是切片

```go
a := [...]int{1, 2, 3} // 数组
s := []int{1, 2, 3}    // 切片

fmt.Printf("%T\n", a) // [3]int
fmt.Printf("%T\n", s) // []int
```

区别只在中括号里有没有 `...` 或数字。

### 坑4：range 的值是副本

```go
nums := [3]int{1, 2, 3}

for _, v := range nums {
    v = 100
}

fmt.Println(nums) // [1 2 3]
```

`v` 是元素副本，不是原元素。

如果要修改原数组：

```go
for i := range nums {
    nums[i] = 100
}
```

### 坑5：数组长度不同就是不同类型

```go
func print3(nums [3]int) {}

print3([3]int{1, 2, 3})
// print3([4]int{1, 2, 3, 4}) // 编译错误
```

如果不想限制长度，用切片：

```go
func printAny(nums []int) {}
```

---

## 13. 数组 vs 切片

| 对比项 | 数组 Array | 切片 Slice |
|------|------------|------------|
| 长度 | 固定 | 可变 |
| 类型是否包含长度 | 是，`[3]int` 和 `[4]int` 不同 | 否，都是 `[]int` |
| 赋值 | 复制所有元素 | 复制切片头，共享底层数组 |
| 传参 | 复制整个数组 | 复制切片头 |
| 能否 append | 不能 | 能 |
| 能否比较 | 元素可比较时可以 | 只能和 nil 比较 |
| 日常使用频率 | 较少 | 非常高 |

一句话：

- 数组强调 **固定长度和值语义**
- 切片强调 **动态长度和共享底层数组**

---

## 14. 面试高频问题

### 14.1 `[3]int` 和 `[]int` 有什么区别？

`[3]int` 是长度为 3 的数组，长度是类型的一部分。

`[]int` 是切片，长度可变，底层引用一个数组。

```go
a := [3]int{1, 2, 3}
s := []int{1, 2, 3}

fmt.Printf("%T\n", a) // [3]int
fmt.Printf("%T\n", s) // []int
```

### 14.2 数组传参会发生什么？

会复制整个数组。函数内修改不会影响外部数组。

```go
func f(a [3]int) {
    a[0] = 100
}
```

如果要修改外部数组，传数组指针或切片。

### 14.3 数组能不能作为 map key？

可以，前提是数组元素类型可比较。

```go
m := map[[2]int]string{}
m[[2]int{1, 2}] = "ok"
```

### 14.4 range 数组会不会复制？

会。`range` 数组时会复制数组。大数组应避免直接 range，可以 range 数组指针或切片。

### 14.5 为什么 Go 日常更常用切片？

因为大多数业务数据长度不固定，而切片支持动态增长、传参轻量、使用方便。

数组更适合长度具有业务含义的固定数据。

---

## 15. 开发速查

```go
// 定长数组
arr := [3]int{1, 2, 3}

// 自动推导长度
arr := [...]int{1, 2, 3}

// 访问元素
fmt.Println(arr[0])

// 修改元素
arr[0] = 100

// 长度
fmt.Println(len(arr))

// 遍历
for i, v := range arr {
    fmt.Println(i, v)
}

// 数组转切片
s := arr[:]

// 数组指针传参
func modify(arr *[3]int) {
    arr[0] = 100
}
```
