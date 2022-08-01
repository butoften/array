- [javascript 味的 golang 数组](#javascript-味的-golang-数组)
    - [概述](#概述)
        - [初衷](#初衷)
    - [安装](#安装)
        - [requirement](#requirement)
        - [to install](#to-install)
    - [初始化:](#初始化)
        - [普通使用](#普通使用)
        - [返回地址的方式(指针)](#返回地址的方式指针)
        - [从已存在切片初始化 array.New](#从已存在切片初始化-arraynew)
    - [常用方法](#常用方法)
        - [Map](#map)
        - [Push](#push)
        - [Pop](#pop)
        - [Shift](#shift)
        - [Unshift](#unshift)
        - [切片清空](#切片清空)
        - [切片断开式清空](#切片断开式清空)
        - [Find 搜索](#find-搜索)
        - [FindLast](#findlast)
        - [FindIndex](#findindex)
        - [FindLastIndex](#findlastindex)
        - [Filter根据条件过滤](#filter根据条件过滤)
    - [排序](#排序)
        - [Sort  原生排序](#sort--原生排序)
        - [冒泡排序](#冒泡排序)
        - [选择排序](#选择排序)
        - [插入排序](#插入排序)
        - [希尔排序](#希尔排序)
        - [归并排序](#归并排序)
        - [快速排序](#快速排序)
## javascript 味的 golang 数组

#### 概述

###### 初衷

> golang本身并没有提供太多数组相关的操作api，所以诞生了此工具包。开发过前端的朋友对es6的语法并不陌生，所以本工具包模拟es6+与js的常用方法来实现了这个工具包来操作数组。

#### 安装

###### requirement

```
go 1.18
```

###### to install

```
go get github.com/butoften/array
```

#### 初始化:

###### 普通使用

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

###### 返回地址的方式(指针)
> 可以避免开发者使用&符号于取一次地址

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

###### 从已存在切片初始化 array.New

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
```
#### 常用方法

###### Map

> 返回一个新数组，数组中的元素为原始数组元素调用函数处理后的值
>
> 方法按照原始数组元素顺序依次处理元素。
>
> 不会改变原始数组

```go
objArr := array.New[Test](Test{
  id:   1,
  name: "A",
}, Test{
  id:   2,
  name: "C",
})
newArr := objArr.Map(func(item Test, index int) any {
  return item.name
})
fmt.Printf("newArr: %v\n", newArr)//[A C]
fmt.Printf("objArr: %v\n", objArr)//[{1 A} {2 C}]
```



###### Push

> 方法可向数组的末尾添加一个或多个元素，并返回新的长度。

```go
arr := array.New[int](3, 44, 38)
fmt.Printf("arr: %v\n", arr)
newLen := arr.Push(1)
fmt.Printf("newLen: %v\n", newLen)
newLen = arr.Push(2)
fmt.Printf("newLen: %v\n", newLen)
newLen = arr.Push(3, 5, 6, 7)
fmt.Printf("arr: %v\n", arr)
fmt.Printf("newLen: %v\n", newLen)
```

###### Pop

> Pop() 方法用于删除数组的最后一个元素并返回删除的元素。
>
> 注意：此方法改变数组的长度！

```go
arr := array.New[int](3, 44, 8)
fmt.Printf("arr: %v\n", arr)
last, ok := arr.Pop()
fmt.Printf("last: %v-%v\n", last, ok) //8-true
fmt.Printf("arr: %v\n", arr)          //[3 44]

arr = array.New[int]()
fmt.Printf("arr: %v\n", arr)
last, ok = arr.Pop()
fmt.Printf("last: %v-%v\n", last, ok) //0-false
fmt.Printf("arr: %v\n", arr)          //[]
```
###### Shift

> Shift() 方法用于把数组的第一个元素从其中删除，并返回第一个元素的值。
>
> 此方法改变数组的长度！

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
first, ok := arr.Shift()
fmt.Printf("arr: %v\n", arr)//[44 38]
fmt.Printf("first: %v-%v\n", first, ok) //3 true

arr = array.New[int]()
first, ok = arr.Shift()
fmt.Printf("arr: %v\n", arr)//[]
fmt.Printf("first: %v-%v\n", first, ok)//0 false
```

###### Unshift

> UnShift() 方法可向数组的开头添加一个或更多元素，并返回新的长度
>
> 此方法改变数组的长度！

```go
fruits := array.New[string]("Banana", "Orange", "Apple", "Mango")
fmt.Printf("fruits: %v\n", fruits)
length := fruits.UnShift("Lemon", "Pineapple")//[Banana Orange Apple Mango]
fmt.Printf("fruits: %v\n", fruits) //[Lemon Pineapple Banana Orange Apple Mango]
fmt.Printf("length: %v\n", length) //6 
```


###### 切片清空

> 只清len不清cap

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.Empty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-3
```

###### 切片断开式清空
> len与cap同时清空，断开底层数组

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.BrokenEmpty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-0
```

###### Find 搜索
> 返回结果为 res, exist 其中res为目标结果 ，exist 为bool
>
> 从前向后遍历

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

###### FindLast

> 与Find类似，不同的是，从后向前遍历

```go
objArr := array.New[Test]()
objArr.Push(Test{
  id:   1,
  name: "A",
})
objArr.Push(Test{
  id:   2,
  name: "C",
})

res, exist := objArr.FindLast(func(item Test, key int) bool {
  return item.name == "C" && item.id == 2
})
if exist {
  fmt.Printf("res: %v\n", res)//{2 C}
} else {
  fmt.Printf("not found: %v\n", res)
}
```

###### FindIndex

> FindIndex()返回符合传入回调函数条件的第一个元素索引位置
>
> 如果没有符合条件的元素返回 -1
>
> 从前向后遍历

```go
ages := array.New[int](3, 10, 18, 20)
index := ages.FindIndex(func(item, index int) bool {
  return item == 18
})
fmt.Printf("index: %v\n", index)//2
```

###### FindLastIndex

> 与FindIndex类似，不同的是，从后向前遍历
>
> 如果没有符合条件的元素返回 -1

```go
ages := array.New[int](3, 10, 18, 20)
index := ages.FindLastIndex(func(item, index int) bool {
  return item > 10
})
fmt.Printf("index: %v\n", index) //3
```

###### Filter根据条件过滤

> 返回结果依然是一个数组，如果没有匹配项，则返回空数组

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


#### 排序

###### Sort  原生排序

> golang原生sort.Slice排序封装

```go
type Test struct {
	id   float32
	name string
}
func main(){
  tempB := []Test{
		{
			id:   3.2,
			name: "A1",
		},
		{
			id:   4.2,
			name: "A2",
		},
		{
			id:   38.9,
			name: "A3",
		},
		{
			id:   5.4,
			name: "A4",
		},
		{
			id:   38.7,
			name: "A5",
		},
		{
			id:   38.5,
			name: "A6",
		},
	}
	arrTest := array.New[Test](tempB...)
	fmt.Printf("arr: %v\n", arrTest)
	//升序
	arrTest.Sort(func(a, b Test) bool {
		return a.id < b.id
	})
	fmt.Printf("Sort: %v\n", arrTest)
  //降序
	arrTest.Sort(func(a, b Test) bool {
		return a.id > b.id
	})
	fmt.Printf("Sort: %v\n", arrTest)
}
```

###### 冒泡排序


```go
arr := array.New[int](3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48)
fmt.Printf("arr: %v\n", arr)
//升序
arr.BubbleSort(func(a, b int) bool {
  return a < b
})

fmt.Printf("arr: %v\n", arr)
```
###### 选择排序

```go
//升序
arr.SelectSort(func(a, b int) bool {
  return a < b
})
```
###### 插入排序

```go
//升序
arr.InsertSort(func(a, b int) bool {
  return a < b
})
```
###### 希尔排序

```go
//升序
arr.ShellSort(func(a, b int) bool {
  return a < b
})
```
###### 归并排序

```go
//升序
arr.MergeSort(func(a, b int) bool {
  return a < b
})
```
###### 快速排序

```go
//升序
arr.QuickSort(func(a, b int) bool {
  return a < b
})
```
