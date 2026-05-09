# Go Map 详解 —— C++ 转型指南

## 1. 什么是 Map？

Map 是 **键值对（key-value）** 的无序集合，类似 C++ 的 `std::unordered_map`（哈希表实现）。

**C++ 对比：**

```cpp
// C++: unordered_map
std::unordered_map<std::string, int> m;
m["apple"] = 1;
```

```go
// Go: map
m := make(map[string]int)
m["apple"] = 1
```

---

## 2. 声明和初始化

### 方式一：make（最常用）

```go
m := make(map[string]int)
```

### 方式二：字面量（初始化时就有数据）

```go
m := map[string]int{
    "apple":  1,
    "banana": 2,
}
```

### 方式三：空 Map

```go
var m map[string]int  // nil map，不能直接赋值
```

**C++ 对比：**

```cpp
std::map<std::string, int> m1;           // 空 map
std::map<std::string, int> m2 = {{"apple", 1}};  // 字面量
```

---

## 3. 基本操作

```go
package main

import "fmt"

func main() {
    m := make(map[string]int)

    // 赋值
    m["apple"] = 1
    m["banana"] = 2

    // 读取（如果 key 不存在，返回零值）
    fmt.Println(m["apple"])           // 1
    fmt.Println(m["grape"])           // 0（零值）

    // 删除（key 不存在时，删除操作无效，不会报错）
    delete(m, "banana")

    // 检查 key 是否存在
    v, ok := m["apple"]   // ok == true 表示存在
    fmt.Println(v, ok)

    v2, ok2 := m["grape"] // ok2 == false，v2 是零值
    fmt.Println(v2, ok2)
}
```

**C++ 对比：**

```cpp
m.erase("banana");

// C++ 检查 key 是否存在
auto it = m.find("apple");
if (it != m.end()) {
    std::cout << it->second;
}
```

**关键区别：**
- C++ 用 `find()` 返回迭代器判断存在
- Go 用 **双返回值** `(value, ok)` 判断存在，更简洁

---

## 4. 遍历

```go
m := map[string]int{
    "apple":  1,
    "banana": 2,
    "grape":  3,
}

for key, value := range m {
    fmt.Println(key, value)
}
```

**注意：** Map 遍历是 **无序的**，每次遍历顺序可能不同。

**C++ 对比：**

```cpp
for (const auto& [key, value] : m) {
    std::cout << key << value;
}
```

---

## 5. 获取 Map 长度

```go
m := map[string]int{"a": 1, "b": 2}
fmt.Println(len(m))  // 2
```

**C++ 对比：** `m.size()`

---

## 6. nil Map vs 空 Map

| 类型 | 描述 | 能赋值吗？ |
|------|------|----------|
| `nil map` | `var m map[string]int` 声明但未初始化 | ❌ 会 panic |
| `empty map` | `make(map[string]int)` 或 `map[string]int{}` | ✅ 可以 |

```go
var nilMap map[string]int
// nilMap["a"] = 1  // panic: assignment to entry in nil map

emptyMap := make(map[string]int)
emptyMap["a"] = 1  // ✅ 正常
```

**C++ 对比：** C++ 的 `std::map<int, int>* p = nullptr` 类似 nil map，解引用会崩溃。

---

## 7. 完整示例

```go
package main

import "fmt"

func main() {
    // 创建
    m := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }

    // 添加
    m["grape"] = 12

    // 读取
    fmt.Printf("apple: %d\n", m["apple"])

    // 检查存在
    if val, ok := m["watermelon"]; ok {
        fmt.Println("watermelon exists:", val)
    } else {
        fmt.Println("watermelon not found, val is zero:", val)
    }

    // 删除
    delete(m, "banana")

    // 遍历
    fmt.Println("\nAll fruits:")
    for key, val := range m {
        fmt.Printf("  %s: %d\n", key, val)
    }

    // 长度
    fmt.Printf("\nTotal: %d types of fruit\n", len(m))
}
```

---

## 8. Go Map vs C++ unordered_map 关键区别

| 特性 | Go | C++ |
|------|-----|-----|
| 底层实现 | 哈希表 | 哈希表（unordered_map）或红黑树（map） |
| 声明方式 | `make(map[K]V)` | `unordered_map<K, V>` |
| key 存在检查 | `v, ok := m[k]` 双返回值 | `m.find(k) != m.end()` |
| 遍历顺序 | 无序 | 无序（unordered_map）或有序（map） |
| 并发安全 | ❌ 不安全，需要 sync.RWMutex | ❌ 不安全 |
| nil 状态 | 有 nil map | 有 nullptr |
| 字面量语法 | `map[K]V{...}` | `unordered_map<K, V>{...}` |

---

## 9. 需要注意的坑

### 坑1：读取不存在的 key 返回零值，无法区分"不存在"和"值为零"

```go
m := map[string]int{}
v := m["not_exist"]  // v == 0，但不知道是不存在还是值就是 0

// ✅ 正确做法：用 ok 判断
if v, ok := m["not_exist"]; ok {
    fmt.Println("exists:", v)
} else {
    fmt.Println("not exists")
}
```

### 坑2：Map 是引用类型

```go
m1 := map[string]int{"a": 1}
m2 := m1      // m2 和 m1 底层指向同一个 Map
m2["a"] = 999
fmt.Println(m1["a"])  // 999，m1 也被改了
```

### 坑3：Map 不能用 `==` 比较

```go
// 编译错误！Map 不能直接比较
// if m1 == m2 { }
```

### 坑4：遍历顺序不固定，不能依赖输出顺序

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

for k, v := range m {
    fmt.Println(k, v)
}
```

上面代码每次运行的输出顺序都可能不同。日常开发中，如果要稳定输出，比如生成日志、签名、测试快照、配置文件，就需要先排序 key。

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
```

### 坑5：Map 不是并发安全的

多个 goroutine 同时读写同一个 map，可能直接 panic。

```go
m := make(map[string]int)

go func() {
    m["a"] = 1
}()

go func() {
    fmt.Println(m["a"])
}()
```

可能报错：

```text
fatal error: concurrent map read and map write
```

日常开发中有两种常见解决方式：

```go
type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (s *SafeMap) Get(key string) (int, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    v, ok := s.m[key]
    return v, ok
}

func (s *SafeMap) Set(key string, value int) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.m[key] = value
}
```

或者使用标准库的 `sync.Map`，它更适合读多写少、key 比较稳定的场景。

```go
var m sync.Map

m.Store("apple", 1)
value, ok := m.Load("apple")
if ok {
    fmt.Println(value)
}
```

### 坑6：删除遍历中的 key 是安全的，但新增 key 不要依赖是否会被遍历到

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

for k := range m {
    if k == "b" {
        delete(m, k) // 安全
    }
}
```

但遍历过程中新增元素时，新元素可能被遍历到，也可能不会，不要依赖这个行为。

```go
for k := range m {
    m["new"] = 100 // 不推荐依赖 new 是否会出现在本轮遍历中
    fmt.Println(k)
}
```

### 坑7：value 是结构体时，不能直接修改字段

```go
type User struct {
    Name string
    Age  int
}

m := map[string]User{
    "tom": {Name: "Tom", Age: 18},
}

// 编译错误：不能直接修改 map 中结构体值的字段
// m["tom"].Age = 20
```

原因是 `m["tom"]` 取出来的是一个临时副本，不能直接对它的字段赋值。

正确做法一：取出、修改、再放回去。

```go
user := m["tom"]
user.Age = 20
m["tom"] = user
```

正确做法二：map 的 value 使用指针。

```go
m2 := map[string]*User{
    "tom": {Name: "Tom", Age: 18},
}

m2["tom"].Age = 20
```

使用指针时要注意 nil 判断。

### 坑8：nil map 可以读、可以 len、可以 delete，但不能写

```go
var m map[string]int

fmt.Println(m["a"]) // 0
fmt.Println(len(m))  // 0
delete(m, "a")      // 不报错

// m["a"] = 1       // panic
```

所以函数里如果要写入 map，要确保它已经初始化。

```go
func addScore(scores map[string]int, name string, score int) {
    scores[name] = score
}

func main() {
    scores := make(map[string]int)
    addScore(scores, "tom", 90)
}
```

### 坑9：map 的 key 必须是可比较类型

Go 的 map key 必须支持 `==` 比较。

可以作为 key 的类型：

- `string`
- 整数、浮点数、布尔值
- 指针
- 数组（元素也必须可比较）
- 结构体（所有字段都必须可比较）

不能作为 key 的类型：

- slice
- map
- function

```go
// 编译错误：slice 不能作为 map key
// m := map[[]int]string{}
```

如果确实要用 slice 表示 key，常见做法是转换成字符串。

```go
key := strings.Join([]string{"go", "map"}, ":")
m := map[string]int{key: 1}
```

### 坑10：map 扩容后，不能获取元素地址

Go 不允许直接对 map 元素取地址。

```go
m := map[string]int{"a": 1}
// p := &m["a"] // 编译错误
```

因为 map 可能扩容搬迁元素，元素地址不稳定。如果需要长期持有地址，可以让 value 本身就是指针。

```go
m := map[string]*int{}
x := 1
m["a"] = &x
```

### 坑11：预估容量可以减少扩容成本

如果大概知道要放多少数据，可以给 `make` 传入容量提示。

```go
m := make(map[string]int, 1000)
```

第二个参数不是固定容量，只是容量提示。map 仍然可以继续增长。

### 坑12：清空 map 的几种方式

方式一：重新创建一个新 map。

```go
m = make(map[string]int)
```

方式二：逐个删除，保留原 map 对象。

```go
for k := range m {
    delete(m, k)
}
```

如果还有其他变量引用同一个 map，重新 `make` 只会让当前变量指向新 map，其他引用仍然指向旧 map。

```go
m1 := map[string]int{"a": 1}
m2 := m1

m1 = make(map[string]int)
fmt.Println(m2) // map[a:1]
```

---

## 10. 日常开发常见用法

### 10.1 用 map 做集合 Set

Go 没有内置 Set，常用 `map[T]struct{}` 模拟。

```go
set := make(map[string]struct{})

set["go"] = struct{}{}
set["cpp"] = struct{}{}

if _, ok := set["go"]; ok {
    fmt.Println("go exists")
}
```

也可以用 `map[string]bool`，但 `struct{}` 不占额外空间，更常见。

```go
set := map[string]bool{}
set["go"] = true
```

### 10.2 用 map 统计次数

```go
words := []string{"go", "cpp", "go", "java", "go"}
count := make(map[string]int)

for _, word := range words {
    count[word]++
}

fmt.Println(count) // map[cpp:1 go:3 java:1]
```

即使 key 不存在，`count[word]` 也会返回 int 零值 `0`，所以可以直接 `++`。

### 10.3 用 map 分组

```go
type User struct {
    Name string
    City string
}

users := []User{
    {Name: "Tom", City: "Beijing"},
    {Name: "Jerry", City: "Shanghai"},
    {Name: "Alice", City: "Beijing"},
}

groups := make(map[string][]User)
for _, user := range users {
    groups[user.City] = append(groups[user.City], user)
}
```

这里 `groups[user.City]` 不存在时会返回 `nil slice`，而 `append(nil, user)` 是合法的。

### 10.4 用 map 做查表替代大量 switch

```go
func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }

ops := map[string]func(int, int) int{
    "+": add,
    "-": sub,
}

if op, ok := ops["+"]; ok {
    fmt.Println(op(1, 2))
}
```

读取函数类型 value 时，要判断是否存在，避免调用 nil 函数。

---

## 11. 面试/开发常用操作速查

```go
// 创建
m := make(map[string]int)

// 添加/修改
m["key"] = value

// 读取
v := m["key"]

// 检查存在
v, ok := m["key"]

// 删除（安全，key 不存在也不报错）
delete(m, "key")

// 长度
n := len(m)

// 遍历
for k, v := range m { }

// 清空（重新 make）
m = make(map[string]int)
```
