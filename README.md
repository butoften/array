- [javascript 味的 golang 数组](#javascript-味的-golang-数组)
    - [概述](#概述)
        - [初衷](#初衷)
    - [安装](#安装)
        - [requirement](#requirement)
        - [to install](#to-install)
    - [初始化:](#初始化)
        - [普通初始化array.New](#普通初始化arraynew)
        - [返回地址的方式(指针)初始化 array.PNew](#返回地址的方式指针初始化-arraypnew)
        - [从已存在切片初始化](#从已存在切片初始化)
    - [常用方法](#常用方法)
        - [Map](#map)
        - [Every](#every)
        - [Some](#some)
        - [Push](#push)
        - [Pop](#pop)
        - [Shift](#shift)
        - [Unshift](#unshift)
        - [Slice](#slice)
        - [Splice](#splice)
        - [Empty](#empty)
        - [BrokenEmpty](#brokenempty)
        - [Find](#find)
        - [FindLast](#findlast)
        - [FindIndex](#findindex)
        - [FindLastIndex](#findlastindex)
        - [Filter](#filter)
    - [排序](#排序)
        - [Sort  原生排序](#sort--原生排序)
        - [BubbleSort 冒泡排序](#bubblesort-冒泡排序)
        - [SelectSort 选择排序](#selectsort-选择排序)
        - [InsertSort 插入排序](#insertsort-插入排序)
        - [ShellSort 希尔排序](#shellsort-希尔排序)
        - [MergeSort 归并排序](#mergesort-归并排序)
        - [QuickSort 快速排序](#quicksort-快速排序)
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

###### 普通初始化array.New

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

###### 返回地址的方式(指针)初始化 array.PNew
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

###### 从已存在切片初始化

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
```
#### 常用方法

###### Map

> * 返回一个新数组，数组中的元素为原始数组元素调用函数处理后的值
>
> * 方法按照原始数组元素顺序依次处理元素。
>
> * 不会改变原始数组

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

###### Every

> * 用于检测数组所有元素是否都符合指定条件（通过函数提供）
>
> * 如果数组中检测到有一个元素不满足，则整个表达式返回 false，都满足时，返回true
>
> * 注：如果是空数组，直接返回false ，这里与js里不一样。
>
> * 不会改变原始数组

```go
arr := array.New[int](1, 2, 4, 5)
res := arr.Every(func(item, index int) bool {
  return item > 2
})
fmt.Printf("res: %v\n", res) //false
res = arr.Every(func(item, index int) bool {
  return item > 0
})
fmt.Printf("res: %v\n", res) //true

arr = array.New[int]()
res = arr.Every(func(item, index int) bool {
  return item > 2
})
fmt.Printf("res: %v\n", res) //false
```

###### Some

> * 用于检测数组中的元素是否满足指定条件（函数提供），只要有一个满足条件，就返回true
>
> * 如果没有满足条件的元素，则返回false
>
> * 如果是空数组，直接返回false
>
> * 不会改变原始数组

```go
arr := array.New[int](1, 2, 4, 5)
res := arr.Some(func(item, index int) bool {
  return item > 2
})
fmt.Printf("res: %v\n", res) //true

res = arr.Some(func(item, index int) bool {
  return item > 5
})
fmt.Printf("res: %v\n", res) //false

arr = array.New[int]()
res = arr.Some(func(item, index int) bool {
  return item > 2
})
fmt.Printf("res: %v\n", res) //false
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

> * last, ok = arr.Pop()
> * Pop() 方法用于删除数组的最后一个元素并返回删除的元素。
> * 注意：此方法改变数组的长度！
> * 空数组 Pop会失败 ok为false

```go
arr := array.New[int](3, 44, 8)
fmt.Printf("arr: %v\n", arr)
last, ok := arr.Pop()
fmt.Printf("last: %v-%v\n", last, ok) //8-true
fmt.Printf("arr: %v\n", arr)          //[3 44]

arr = array.New[int]()//空数组 pop会失败
fmt.Printf("arr: %v\n", arr)
last, ok = arr.Pop()
fmt.Printf("last: %v-%v\n", last, ok) //0-false
fmt.Printf("arr: %v\n", arr)          //[]
```
###### Shift

> * first, ok := arr.Shift()
> * Shift() 方法用于把数组的第一个元素从其中删除，并返回第一个元素的值。
> * 此方法改变数组的长度！
> * 空数组Shift会失败，ok为false

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
first, ok := arr.Shift()
fmt.Printf("arr: %v\n", arr)//[44 38]
fmt.Printf("first: %v-%v\n", first, ok) //3 true

arr = array.New[int]()//空数组
first, ok = arr.Shift()
fmt.Printf("arr: %v\n", arr)//[]
fmt.Printf("first: %v-%v\n", first, ok)//0 false
```

###### Unshift

> * UnShift() 方法可向数组的开头添加一个或更多元素，并返回新的长度
>
> * 此方法改变数组的长度！

```go
fruits := array.New[string]("Banana", "Orange", "Apple", "Mango")
fmt.Printf("fruits: %v\n", fruits)
length := fruits.UnShift("Lemon", "Pineapple")//[Banana Orange Apple Mango]
fmt.Printf("fruits: %v\n", fruits) //[Lemon Pineapple Banana Orange Apple Mango]
fmt.Printf("length: %v\n", length) //6 
```
###### Slice

> * 从已有的数组中返回选定区间的新元素数组，返回类型为array.Array，可以继续使用此工具包的各种方法
>
> * 此方法不会对源数组产生影响（原生切片因扩容规则：不扩容的情况下，会对源切片产生影响）
>
> * 如果你不喜欢此方法你依然可以使用原生切片截取方式[:]来操作，但要注意扩容规则
>
> * 因为js里slice方法支持起始与结束值为负，所以本工具包也实现了相同的算法，请查看下面非常规案例

```go
a := []int{1, 2, 3, 4, 5}
b := a[0:1]
fmt.Printf("a: %v\n", a) //[1 2 3 4 5]
fmt.Printf("b: %v\n", b) //[1]
b = append(b, 6, 7, 8, 9)
fmt.Printf("a: %v\n", a) //[1 6 7 8 9] 因为整体cap不变，没有扩容，所以b切片在append时，a切片也受到了影响
fmt.Printf("b: %v\n", b) //[1 6 7 8 9]

arr := array.New[int](1, 2, 3, 4, 5)
newArr := arr.Slice(0, 1)
fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5]
fmt.Printf("newArr: %v\n", newArr) //[1]
newArr = append(newArr, 6, 7)
newArr.Push(8, 9)
fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5] array.Slice方法由于重新初始切片，所以实现了彼此互不影响
fmt.Printf("newArr: %v\n", newArr) //[1 6 7 8 9]

c := newArr[0:1] //使用原生切片截取方式 但是要注意扩容规则
c = append(c, 66, 77, 88, 99)
fmt.Printf("原生[:] 截取并append后 newArr: %v\n", newArr) //[1 66 77 88 99] 受扩容规则影响
fmt.Printf("原生[:] 截取并append后 c: %v\n", c)           //[1 66 77 88 99]

//以下是非常规操作案例
d := array.New[int](1, 2, 3, 4, 5)
e := d.Slice(0, -1)
fmt.Printf("e: %v\n", e) //[1 2 3 4]
e = d.Slice(0, 20)
fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
e = d.Slice(-20, 20)
fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
e = d.Slice(-1, 20)
fmt.Printf("e: %v\n", e) //[5]
e = d.Slice(-5, 20)
fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
e = d.Slice(-1, -2)
fmt.Printf("e: %v\n", e) //[]
e = d.Slice(-2, -1)
fmt.Printf("e: %v\n", e) //[4]

e = d.Slice(-20, -1)
fmt.Printf("e: %v\n", e) //[1 2 3 4]
```

###### Splice

> * Splice(index int, howMany int, args ...T)
> * @param index 规定从何处添加或删除元素，该参数是插入元素或删除元素的起始下标，必须是整数
> * @param howMany 规定应该删除多少元素
>     howMany 为0时，表示不删除，如果 args有值则在index处插入args
>     howMany 为负数时，相当于0，不删除
>     howMany 为正数时，表示删除数量，删除后，如果 args有值则在index处插入args
> * @param args 要在index处添加的多个元素
> * 用于添加或删除数组中的元素
> * 会改变原始数组
> * 返回的是含有被删除的元素的数组

```go
arr := array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr := arr.Splice(3, 0, 11, 12, 13)
fmt.Printf("delArr: %v\n", delArr) //[]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 13 4 5 6 7 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(3, 2, 11, 12, 13)
fmt.Printf("delArr: %v\n", delArr) //[4 5]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 13 6 7 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(3, 3, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[4 5 6]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 7 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(3, 4, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[4 5 6 7]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(-3, 5, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[8 9 10]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 11 12]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(-3, -5, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[]
fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 11 12 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(0, 5, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[1 2 3 4 5]
fmt.Printf("arr: %v\n", arr)       //[11 12 6 7 8 9 10]

arr = array.New[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
delArr = arr.Splice(0, 50, 11, 12)
fmt.Printf("delArr: %v\n", delArr) //[1 2 3 4 5 6 7 8 9 10]
fmt.Printf("arr: %v\n", arr)       //[11 12]
```

###### Empty

> * 切片清空
>
> * 只清len不清cap

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.Empty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-3
```

###### BrokenEmpty
> * 切片断开式清空
> * len与cap同时清空，断开底层数组

```go
tempA := []int{3, 44, 38}
arr := array.New[int](tempA...)
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: [3 44 38]-3-3
arr.BrokenEmpty()
fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr))//arr: []-0-0
```

###### Find
> * 根据回调函数进行搜索
>* 返回结果为 res, exist 其中res为目标结果 ，exist 为bool
> * 从前向后遍历

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

> * FindIndex()返回符合传入回调函数条件的第一个元素索引位置
>
> * 如果没有符合条件的元素返回 -1
>
> * 从前向后遍历

```go
ages := array.New[int](3, 10, 18, 20)
index := ages.FindIndex(func(item, index int) bool {
  return item == 18
})
fmt.Printf("index: %v\n", index)//2
```

###### FindLastIndex

> * 与FindIndex类似，不同的是，从后向前遍历
>
> * 如果没有符合条件的元素返回 -1

```go
ages := array.New[int](3, 10, 18, 20)
index := ages.FindLastIndex(func(item, index int) bool {
  return item > 10
})
fmt.Printf("index: %v\n", index) //3
```

###### Filter

> * 根据回调函数里的条件进行过滤
> * 返回结果依然是一个数组，如果没有匹配项，则返回空数组

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

###### BubbleSort 冒泡排序


```go
arr := array.New[int](3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48)
fmt.Printf("arr: %v\n", arr)
//升序
arr.BubbleSort(func(a, b int) bool {
  return a < b
})

fmt.Printf("arr: %v\n", arr)
```
###### SelectSort 选择排序

```go
//升序
arr.SelectSort(func(a, b int) bool {
  return a < b
})
```
###### InsertSort 插入排序

```go
//升序
arr.InsertSort(func(a, b int) bool {
  return a < b
})
```
###### ShellSort 希尔排序

```go
//升序
arr.ShellSort(func(a, b int) bool {
  return a < b
})
```
###### MergeSort 归并排序

```go
//升序
arr.MergeSort(func(a, b int) bool {
  return a < b
})
```
###### QuickSort 快速排序

```go
//升序
arr.QuickSort(func(a, b int) bool {
  return a < b
})
```
