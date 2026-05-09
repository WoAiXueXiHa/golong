# Go Slice 切片详解 —— 直击本质、日常开发与面试

## 1. 一句话理解切片

切片 `slice` 是 **对底层数组某一段的描述**，不是数组本身。

可以把切片理解成一个小结构体：

```go
type slice struct {
    ptr *T // 指向底层数组
    len int // 当前长度
    cap int // 容量
}
```

核心本质：

- 切片本身不直接存数据，数据在底层数组里
- 多个切片可能共享同一个底层数组
- `append` 可能复用旧数组，也可能创建新数组
- 切片传参会复制切片头，但共享底层数组

---

## 2. 数组和切片的关系

数组是实体，切片是视图。

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]

fmt.Println(s)      // [2 3 4]
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 4
```

`s := arr[1:4]` 表示从 `arr[1]` 开始，看到 `arr[1]` 到 `arr[3]`，左闭右开。

```text
arr: [1 2 3 4 5]
        ↑
        s.ptr
s.len = 3 -> [2 3 4]
s.cap = 4 -> [2 3 4 5]
```

修改切片会影响底层数组：

```go
arr := [3]int{1, 2, 3}
s := arr[:]
s[0] = 100
fmt.Println(arr) // [100 2 3]
```

---

## 3. 创建切片

### 3.1 字面量

```go
s := []int{1, 2, 3}
```

注意区别：

```go
a := [3]int{1, 2, 3} // 数组
s := []int{1, 2, 3}  // 切片
```

### 3.2 make

```go
s := make([]int, 3)
fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 3
```

指定容量：

```go
s := make([]int, 0, 10)
fmt.Println(len(s), cap(s)) // 0 10
```

日常开发中，如果要不断 `append`，推荐预估容量：

```go
result := make([]int, 0, len(nums))
for _, n := range nums {
    if n > 0 {
        result = append(result, n)
    }
}
```

### 3.3 nil 切片

```go
var s []int
fmt.Println(s == nil) // true
fmt.Println(len(s))   // 0
fmt.Println(cap(s))   // 0

s = append(s, 1)      // 可以直接 append
```

---

## 4. len 和 cap

`len` 是当前能访问的元素个数。

`cap` 是从切片起点到底层数组末尾的容量。

```go
s := []int{1, 2, 3, 4, 5}
sub := s[1:3]

fmt.Println(sub)      // [2 3]
fmt.Println(len(sub)) // 2
fmt.Println(cap(sub)) // 4
```

因为 `sub` 从原切片下标 1 开始，底层后面还有 4 个位置。

```go
fmt.Println(sub[:4]) // [2 3 4 5]
// fmt.Println(sub[:5]) // panic: 超过 cap
```

---

## 5. 切片表达式

### 5.1 基本形式：`s[low:high]`

```go
s := []int{0, 1, 2, 3, 4, 5}

fmt.Println(s[1:4]) // [1 2 3]
fmt.Println(s[:3])  // [0 1 2]
fmt.Println(s[3:])  // [3 4 5]
fmt.Println(s[:])   // [0 1 2 3 4 5]
```

规则：包含 `low`，不包含 `high`。

### 5.2 完整形式：`s[low:high:max]`

第三个参数用来限制容量。

```go
s := []int{0, 1, 2, 3, 4, 5}
sub := s[1:3:3]

fmt.Println(sub)      // [1 2]
fmt.Println(len(sub)) // 2
fmt.Println(cap(sub)) // 2
```

容量计算：`cap = max - low`。

它常用于避免 `append` 修改原底层数组：

```go
s := []int{1, 2, 3, 4}
sub := s[:2:2]
sub = append(sub, 100)

fmt.Println(sub) // [1 2 100]
fmt.Println(s)   // [1 2 3 4]
```

---

## 6. append 的本质

```go
s := []int{1, 2}
s = append(s, 3)
```

`append` 必须接收返回值，因为它可能返回新的切片。

容量够时，复用原底层数组：

```go
s := make([]int, 2, 4)
s[0], s[1] = 1, 2
s2 := append(s, 3)

fmt.Println(s)  // [1 2]
fmt.Println(s2) // [1 2 3]
```

容量不够时，分配新数组并复制旧元素：

```go
s := make([]int, 2, 2)
s[0], s[1] = 1, 2
s2 := append(s, 3)

fmt.Println(s2) // [1 2 3]
```

不要死记扩容倍数，掌握本质即可：

- 容量够：共享底层数组
- 容量不够：新建底层数组
- 所以 append 后一定使用返回值

追加另一个切片：

```go
a := []int{1, 2}
b := []int{3, 4}
a = append(a, b...)
```

---

## 7. copy 的本质

`copy(dst, src)` 返回实际复制的元素个数。

```go
src := []int{1, 2, 3}
dst := make([]int, 2)

n := copy(dst, src)
fmt.Println(n)   // 2
fmt.Println(dst) // [1 2]
```

复制数量是 `min(len(dst), len(src))`。

复制一个完整切片：

```go
b := make([]int, len(a))
copy(b, a)
```

也可以：

```go
b := append([]int(nil), a...)
```

---

## 8. nil 切片 vs 空切片

| 类型 | 写法 | 是否 nil | len | cap |
|------|------|----------|-----|-----|
| nil 切片 | `var s []int` | true | 0 | 0 |
| 空切片 | `[]int{}` | false | 0 | 0 |
| make 空切片 | `make([]int, 0)` | false | 0 | 0 |

```go
var a []int
b := []int{}

fmt.Println(a == nil) // true
fmt.Println(b == nil) // false
```

它们都能 `append`，都能 `range`。

JSON 中有区别：

```go
var a []int
b := []int{}

json.Marshal(a) // null
json.Marshal(b) // []
```

对外 API 如果要求返回 `[]`，不要返回 nil 切片。

---

## 9. 切片作为函数参数

切片传参会复制切片头，但底层数组共享。

```go
func modify(s []int) {
    s[0] = 100
}

func main() {
    nums := []int{1, 2, 3}
    modify(nums)
    fmt.Println(nums) // [100 2 3]
}
```

但函数里修改切片长度，外部看不到：

```go
func add(s []int) {
    s = append(s, 100)
}

func main() {
    nums := []int{1, 2, 3}
    add(nums)
    fmt.Println(nums) // [1 2 3]
}
```

正确做法是返回新切片：

```go
func add(s []int) []int {
    return append(s, 100)
}
```

---

## 10. 共享底层数组的坑

### 坑1：子切片修改原切片

```go
s := []int{1, 2, 3, 4}
sub := s[1:3]
sub[0] = 100

fmt.Println(s) // [1 100 3 4]
```

### 坑2：子切片 append 覆盖原数据

```go
s := []int{1, 2, 3, 4}
sub := s[:2]
sub = append(sub, 100)

fmt.Println(s) // [1 2 100 4]
```

解决方式：限制容量或复制。

```go
sub := s[:2:2]
sub = append(sub, 100)
```

或者：

```go
sub := append([]int(nil), s[:2]...)
```

### 坑3：小切片引用大数组导致内存不释放

```go
func header(data []byte) []byte {
    return data[:10]
}
```

如果 `data` 很大，返回的小切片仍然引用整个大数组。长期保存时应复制：

```go
func header(data []byte) []byte {
    return append([]byte(nil), data[:10]...)
}
```

这是实际开发中非常重要的内存坑。

---

## 11. 日常常见操作

### 11.1 删除元素

保持顺序：

```go
s = append(s[:i], s[i+1:]...)
```

不保持顺序，性能更好：

```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

如果元素是指针或包含大对象引用，删除后建议清理尾部，帮助 GC：

```go
copy(s[i:], s[i+1:])
s[len(s)-1] = nil
s = s[:len(s)-1]
```

### 11.2 过滤元素

复用底层数组：

```go
result := s[:0]
for _, v := range s {
    if keep(v) {
        result = append(result, v)
    }
}
```

### 11.3 一对一转换

```go
result := make([]int, len(nums))
for i, n := range nums {
    result[i] = n * 2
}
```

### 11.4 收集不确定数量结果

```go
result := make([]User, 0, len(users))
for _, user := range users {
    if user.Active {
        result = append(result, user)
    }
}
```

### 11.5 对外返回副本

如果不希望调用方修改内部数据，返回前 copy。

```go
func Keys(keys []string) []string {
    out := make([]string, len(keys))
    copy(out, keys)
    return out
}
```

---

## 12. range 的坑

`range` 的第二个变量是副本。

```go
type User struct {
    Name string
}

users := []User{{Name: "Tom"}, {Name: "Jerry"}}

for _, user := range users {
    user.Name = "Changed"
}

fmt.Println(users) // [{Tom} {Jerry}]
```

正确修改方式：

```go
for i := range users {
    users[i].Name = "Changed"
}
```

如果元素本身是指针，则可以修改指向的对象：

```go
users := []*User{{Name: "Tom"}}
for _, user := range users {
    user.Name = "Changed"
}
```

---

## 13. 二维切片

每行单独分配：

```go
grid := make([][]int, 3)
for i := range grid {
    grid[i] = make([]int, i+1)
}
```

每行长度可以不同。

如果是固定矩阵，可以一次性分配连续内存：

```go
rows, cols := 3, 4
buf := make([]int, rows*cols)
grid := make([][]int, rows)

for i := range grid {
    grid[i] = buf[i*cols : (i+1)*cols]
}
```

优点：分配次数少，内存连续，缓存友好。

---

## 14. 字符串和切片

字符串可以切片，但按字节切：

```go
s := "hello"
fmt.Println(s[1:4]) // ell
```

中文要小心：

```go
s := "你好世界"
fmt.Println(len(s)) // 12，字节数
```

按字符处理用 `[]rune`：

```go
r := []rune(s)
fmt.Println(len(r))       // 4
fmt.Println(string(r[:2])) // 你好
```

字符串转 `[]byte` 会复制：

```go
b := []byte(s)
b[0] = 'H'
```

---

## 15. 并发注意

多个 goroutine 同时 append 同一个切片是不安全的。

```go
s := []int{}

go func() { s = append(s, 1) }()
go func() { s = append(s, 2) }()
```

需要加锁：

```go
var mu sync.Mutex

mu.Lock()
s = append(s, 1)
mu.Unlock()
```

或者每个 goroutine 使用自己的局部切片，最后汇总。

---

## 16. 切片 vs 数组

| 对比项 | 数组 | 切片 |
|------|------|------|
| 长度 | 固定 | 可变 |
| 类型 | 长度是类型一部分 | 长度不是类型一部分 |
| 赋值 | 复制所有元素 | 复制切片头，共享底层数组 |
| 传参 | 复制整个数组 | 复制切片头 |
| append | 不支持 | 支持 |
| 比较 | 元素可比较时可比较 | 只能和 nil 比较 |
| 使用频率 | 较少 | 极高 |

---

## 17. 面试高频问题

### 17.1 切片底层结构是什么？

切片头包含指针、长度、容量，指针指向底层数组。

### 17.2 len 和 cap 区别？

`len` 是当前可见元素数量，`cap` 是从切片起点到底层数组末尾的容量。

### 17.3 append 什么时候扩容？

追加后的长度超过容量时扩容。容量够则复用底层数组，容量不够则新建数组并复制数据。

### 17.4 切片传参会复制吗？

会复制切片头，不复制底层数组。修改元素会影响外部，修改长度不会直接影响外部。

### 17.5 nil 切片和空切片区别？

`len` 都是 0，都能 append。nil 切片等于 nil，JSON 是 `null`；空切片不是 nil，JSON 是 `[]`。

### 17.6 切片能作为 map key 吗？

不能。切片不可比较，只能和 nil 比较。需要时转成字符串或数组。

### 17.7 如何复制切片？

```go
b := make([]int, len(a))
copy(b, a)
```

或：

```go
b := append([]int(nil), a...)
```

### 17.8 如何避免子切片 append 影响原切片？

使用完整切片表达式限制容量：

```go
sub := s[:2:2]
```

或者复制：

```go
sub := append([]int(nil), s[:2]...)
```

---

## 18. 开发速查

```go
// 创建
s := []int{1, 2, 3}

// 长度为 3
s := make([]int, 3)

// 长度为 0，容量为 10
s := make([]int, 0, 10)

// 追加
s = append(s, 4)

// 合并
s = append(s, other...)

// 截取
sub := s[1:3]

// 限制容量截取
sub := s[1:3:3]

// 复制
dst := make([]int, len(s))
copy(dst, s)

// 删除，保持顺序
s = append(s[:i], s[i+1:]...)

// 删除，不保持顺序
s[i] = s[len(s)-1]
s = s[:len(s)-1]

// 过滤
result := s[:0]
for _, v := range s {
    if keep(v) {
        result = append(result, v)
    }
}
```
