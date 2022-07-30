## javascript 味的 golang 数组



### 如果你是前端程序员学习golang，又想使用es6的语法糖来操作数组，可以使用此库

requirement：

```
go 1.18
```

to install：

```
go get github.com/butoften/array
```

切片数组:

```go
type Test struct {
	id   int
	name string
}
func main() {
	arr := array.New[Test]()
	arr.Push(Test{
		id:   1,
		name: "A",
	})
	arr.Push(Test{
		id:   2,
		name: "B",
	})

	fmt.Printf("arr: %v\n", arr)
}
```

切片数组地址(指针)

```go
type Test struct {
	id   int
	name string
}
func main() {
	arr := array.PNew[Test]()
	arr.Push(Test{
		id:   1,
		name: "A",
	})
	arr.Push(Test{
		id:   2,
		name: "B",
	})
	fmt.Printf("arr: %T\n", arr)
	fmt.Printf("arr: %T\n", *arr)
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("arr: %v\n", *arr)
}
```

从已存在切片初始化array.New

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
```

切片清空：只清len不清cap

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.Empty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-3
```

切片清空：len同时cap清空，断开底层数组

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.BrokenEmpty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-0
```



find用法：

```go
type Test struct {
	id   int
	name string
}

func main() {
	arr := array.New[Test]()
	arr.Push(Test{
		id:   1,
		name: "A",
	})
	arr.Push(Test{
		id:   2,
		name: "B",
	})

	res, exist := arr.Find(func(item Test, key int) bool {
		return item.name == "B" && item.id == 2
	})
	if exist {
		fmt.Printf("res: %v\n", res)
	} else {
		fmt.Printf("not found: %v\n", res)
	}
}
```

filter用法：

```go
type Test struct {
	id   int
	name string
}

func main() {
	arr := array.New[Test]()
	arr.Push(Test{
		id:   1,
		name: "A",
	})
	arr.Push(Test{
		id:   2,
		name: "B",
	})

	resFilter := arr.Filter(func(item Test, key int) bool {
		return item.name == "A"
	})

	fmt.Printf("resFilter: %v\n", resFilter)
}
```

排序：

```go
arr := array.New[int]([]int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48})
fmt.Printf("arr: %v\n", arr)

//冒泡排序：升序
arr.BubbleSort(func(a, b int) int {
  return a - b
})
fmt.Printf("arr: %v\n", arr)

//冒泡排序：降序
arr.BubbleSort(func(a, b int) int {
  return b - a
})
fmt.Printf("arr: %v\n", arr)


//选择排序：升序
arr.SelectSort(func(a, b int) int {
  return a - b
})

//插入排序：升序
arr.InsertSort(func(a, b int) int {
  return a - b
})

//希尔排序：升序
arr.ShellSort(func(a, b int) int {
  return a - b
})

//归并排序：升序
arr.MergeSort(func(a, b int) int {
  return a - b
})
```

