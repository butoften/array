package array

import (
	"fmt"
	"testing"
)

type Test struct {
	id   int
	name string
}
type TestFloat struct {
	id   float64
	name string
}
type TestNew struct {
	Test
	age int
}
type TestIII struct {
	TestNew
	haha string
}

func TestReverse(t *testing.T) {
	fmt.Printf("----- Reverse Test Start -----\n")
	originArr := []Test{
		{
			id:   1,
			name: "A",
		},
		{
			id:   2,
			name: "C",
		},
	}
	Reverse(&originArr)
	fmt.Printf("originArr: %v\n", originArr) //[{2 C} {1 A}]
	fmt.Printf("----- Reverse Test End -----\n")
}

func TestToReversed(t *testing.T) {
	fmt.Printf("----- ToReversed Test Start -----\n")
	originArr := []Test{
		{
			id:   1,
			name: "A",
		},
		{
			id:   2,
			name: "C",
		},
	}
	newArr := ToReversed(originArr)
	fmt.Printf("originArr: %v\n", originArr) //[{1 A} {2 C}]
	fmt.Printf("newArr: %v\n", newArr)       //[{2 C} {1 A}]
	fmt.Printf("----- ToReversed Test End -----\n")
}

func TestMap(t *testing.T) {
	fmt.Printf("----- Map Test Start -----\n")
	originArr := []Test{
		Test{
			id:   1,
			name: "A",
		},
		Test{
			id:   2,
			name: "C",
		},
	}
	newArr := Map[Test, TestNew](originArr, func(item Test, index int) TestNew {
		return TestNew{
			Test: item,
			age:  1,
		}
	})
	fmt.Printf("newArr: %v\n", newArr)    //[{{1 A} 1} {{2 C} 1}]
	fmt.Printf("objArr: %v\n", originArr) //[{1 A} {2 C}]
	fmt.Printf("----- Map Test End -----\n\n")
}
func TestForEach(t *testing.T) {
	fmt.Printf("----- ForEach Test Start -----\n")
	objArr := []Test{
		{
			id:   1,
			name: "A",
		},
		{
			id:   2,
			name: "C",
		},
	}
	fmt.Printf("objArr: %v\n", objArr) //[{1 A} {2 C}]
	ForEach(objArr, func(item *Test, index int) {
		item.id += 1
	})
	fmt.Printf("objArr: %v\n", objArr) //[{2 A} {3 C}]
	fmt.Printf("----- ForEach Test End -----\n\n")
}
func TestEvery(t *testing.T) {
	fmt.Printf("----- Every Test Start -----\n")
	originArr := []int{1, 2, 4, 5}
	res := Every[int](originArr, func(item, index int) bool {
		return item > 2
	})
	fmt.Printf("res: %v\n", res) //false
	res = Every[int](originArr, func(item, index int) bool {
		return item > 0
	})
	fmt.Printf("res: %v\n", res) //true
	fmt.Printf("----- Every Test End -----\n\n")
}
func TestSome(t *testing.T) {
	fmt.Printf("----- Some Test Start -----\n")
	originArr := []int{1, 2, 4, 5}
	res := Some(originArr, func(item, index int) bool {
		return item > 2
	})
	fmt.Printf("res: %v\n", res) //true
	res = Some(originArr, func(item, index int) bool {
		return item > 5
	})
	fmt.Printf("res: %v\n", res) //false
	fmt.Printf("----- Some Test End -----\n\n")
}
func TestPush(t *testing.T) {
	fmt.Printf("----- Push Test Start -----\n")
	originArr := []int{3, 44, 38}
	fmt.Printf("originArr: %v\n", originArr) //originArr: [3 44 38]
	newLen := Push(&originArr, 1)
	fmt.Printf("newLen: %v\n", newLen) //newLen: 4
	newLen = Push(&originArr, 2)
	fmt.Printf("newLen: %v\n", newLen) //newLen: 5
	newLen = Push(&originArr, 3, 5, 6, 7)
	fmt.Printf("originArr: %v\n", originArr) //originArr: [3 44 38 1 2 3 5 6 7]
	fmt.Printf("newLen: %v\n", newLen)       //newLen: 9
	fmt.Printf("----- Push Test End -----\n\n")
}
func TestPop(t *testing.T) {
	fmt.Printf("----- Pop Test Start -----\n")
	originArr := []int{3, 44, 8}
	fmt.Printf("originArr: %v\n", originArr)
	last, ok := Pop(&originArr)
	fmt.Printf("last: %v-%v\n", last, ok)    //8-true
	fmt.Printf("originArr: %v\n", originArr) //[3 44]

	originArr = []int{} //空数组 pop会失败
	fmt.Printf("originArr: %v\n", originArr)
	last, ok = Pop(&originArr)
	fmt.Printf("last: %v-%v\n", last, ok)    //0-false
	fmt.Printf("originArr: %v\n", originArr) //[]
	fmt.Printf("----- Pop Test End -----\n\n")
}
func TestShift(t *testing.T) {
	fmt.Printf("----- Shift Test Start -----\n")
	originArr := []int{3, 44, 38}
	fmt.Printf("originArr: %v\n", originArr) //[3 44 38]
	first, ok := Shift(&originArr)
	fmt.Printf("originArr: %v\n", originArr) //[44 38]
	fmt.Printf("first: %v-%v\n", first, ok)  //3 true
	originArr = []int{}
	fmt.Printf("originArr: %v\n", originArr) //[]
	first, ok = Shift(&originArr)
	fmt.Printf("originArr: %v\n", originArr) //[]
	fmt.Printf("first: %v-%v\n", first, ok)  //0 false
	fmt.Printf("----- Shift Test End -----\n\n")
}
func TestUnShift(t *testing.T) {
	fmt.Printf("----- UnShift Test Start -----\n")
	originArr := []string{"Banana", "Orange", "Apple", "Mango"}
	fmt.Printf("originArr: %v\n", originArr)
	length := UnShift(&originArr, "Lemon", "Pineapple") //[Banana Orange Apple Mango]
	fmt.Printf("originArr: %v\n", originArr)            //[Lemon Pineapple Banana Orange Apple Mango]
	fmt.Printf("length: %v\n", length)                  //6
	fmt.Printf("----- UnShift Test End -----\n\n")
}
func TestSlice(t *testing.T) {
	fmt.Printf("----- Slice Test Start -----\n")
	a := []int{1, 2, 3, 4, 5}
	b := a[0:1]
	fmt.Printf("a: %v\n", a) //[1 2 3 4 5]
	fmt.Printf("b: %v\n", b) //[1]
	b = append(b, 6, 7, 8, 9)
	fmt.Printf("a: %v\n", a) //[1 6 7 8 9] 因为整体cap不变，没有扩容，所以b切片在append时，a切片也受到了影响
	fmt.Printf("b: %v\n", b) //[1 6 7 8 9]

	arr := []int{1, 2, 3, 4, 5}
	newArr := Slice(arr, 0, 1)
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5]
	fmt.Printf("newArr: %v\n", newArr) //[1]
	newArr = append(newArr, 6, 7)
	Push(&newArr, 8, 9)
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5] array.Slice方法由于重新初始切片，所以实现了彼此互不影响
	fmt.Printf("newArr: %v\n", newArr) //[1 6 7 8 9]

	c := newArr[0:1] //使用原生切片截取方式 但是要注意扩容规则
	c = append(c, 66, 77, 88, 99)
	fmt.Printf("原生[:] 截取并append后 newArr: %v\n", newArr) //[1 66 77 88 99] 受扩容规则影响
	fmt.Printf("原生[:] 截取并append后 c: %v\n", c)           //[1 66 77 88 99]

	//以下是非常规操作案例
	d := []int{1, 2, 3, 4, 5}
	e := Slice(d, 0, -1)
	fmt.Printf("e: %v\n", e) //[1 2 3 4]
	e = Slice(d, 0, 20)
	fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
	e = Slice(d, -20, 20)
	fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
	e = Slice(d, -1, 20)
	fmt.Printf("e: %v\n", e) //[5]
	e = Slice(d, -5, 20)
	fmt.Printf("e: %v\n", e) //[1 2 3 4 5]
	e = Slice(d, -1, -2)
	fmt.Printf("e: %v\n", e) //[]
	e = Slice(d, -2, -1)
	fmt.Printf("e: %v\n", e) //[4]

	e = Slice(d, -20, -1)
	fmt.Printf("e: %v\n", e) //[1 2 3 4]
	fmt.Printf("----- Slice Test End -----\n\n")
}
func TestSplice(t *testing.T) {
	fmt.Printf("----- Splice Test Start -----\n")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr := Splice(&arr, 3, 0, 11, 12, 13)
	fmt.Printf("delArr: %v\n", delArr) //[]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 13 4 5 6 7 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 3, 2, 11, 12, 13)
	fmt.Printf("delArr: %v\n", delArr) //[4 5]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 13 6 7 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 3, 3, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[4 5 6]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 7 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 3, 4, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[4 5 6 7]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 11 12 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, -3, 5, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[8 9 10]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 11 12]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, -3, -5, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 11 12 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 0, 5, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[1 2 3 4 5]
	fmt.Printf("arr: %v\n", arr)       //[11 12 6 7 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 0, 50, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("arr: %v\n", arr)       //[11 12]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 0, 0)
	fmt.Printf("delArr: %v\n", delArr) //[]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 8 9 10]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 50, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[]
	fmt.Printf("arr: %v\n", arr)       //[1 2 3 4 5 6 7 8 9 10 12]

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delArr = Splice(&arr, 10, 50, 11, 12)
	fmt.Printf("delArr: %v\n", delArr) //[]
	fmt.Printf("arr: %v\n", arr)       //[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
	fmt.Printf("----- Splice Test End -----\n\n")
}
func TestEmpty(t *testing.T) {
	fmt.Printf("----- Empty Test Start -----\n")
	arr := []int{3, 44, 38}
	fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr)) //arr: [3 44 38]-3-3
	Empty(&arr)
	fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr)) //arr: []-0-3
	fmt.Printf("----- Empty Test End -----\n\n")
}
func TestBrokenEmpty(t *testing.T) {
	fmt.Printf("----- BrokenEmpty Test Start -----\n")
	arr := []int{3, 44, 38}
	fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr)) //arr: [3 44 38]-3-3
	BrokenEmpty(&arr)
	fmt.Printf("arr: %v-%v-%v\n", arr, len(arr), cap(arr)) //arr: []-0-0
	fmt.Printf("----- BrokenEmpty Test End -----\n\n")
}
func TestFind(t *testing.T) {
	fmt.Printf("----- Find Test Start -----\n")
	var arr []Test
	Push(&arr, Test{
		id:   1,
		name: "A",
	})
	Push(&arr, Test{
		id:   2,
		name: "B",
	})

	res, exist := Find(arr, func(item Test, key int) bool {
		return item.name == "B" && item.id == 2
	})
	if exist {
		fmt.Printf("res: %v\n", res) //res: {2 B}
	} else {
		fmt.Printf("not found: %v\n", res)
	}
	fmt.Printf("----- Find Test End -----\n\n")
}
func TestFindLast(t *testing.T) {
	fmt.Printf("----- FindLast Test Start -----\n")
	var objArr []Test
	Push(&objArr, Test{
		id:   1,
		name: "A",
	})
	Push(&objArr, Test{
		id:   2,
		name: "C",
	})
	res, exist := FindLast(objArr, func(item Test, key int) bool {
		return item.name == "C" && item.id == 2
	})
	if exist {
		fmt.Printf("res: %v\n", res) //res: {2 C}
	} else {
		fmt.Printf("not found: %v\n", res)
	}
	fmt.Printf("----- FindLast Test End -----\n\n")
}
func TestFindIndex(t *testing.T) {
	fmt.Printf("----- FindIndex Test Start -----\n")
	ages := []int{3, 10, 18, 20}
	index := FindIndex(ages, func(item, index int) bool {
		return item == 18
	})
	fmt.Printf("index: %v\n", index) //index: 2
	fmt.Printf("----- FindIndex Test End -----\n\n")
}
func TestFindLastIndex(t *testing.T) {
	fmt.Printf("----- FindLastIndex Test Start -----\n")
	ages := []int{3, 10, 18, 20}
	index := FindLastIndex(ages, func(item, index int) bool {
		return item > 10
	})
	fmt.Printf("index: %v\n", index) //index: 3
	fmt.Printf("----- FindLastIndex Test End -----\n\n")
}
func TestFilter(t *testing.T) {
	fmt.Printf("----- Filter Test Start -----\n")
	var arr []Test
	Push(&arr, Test{
		id:   1,
		name: "A",
	})
	Push(&arr, Test{
		id:   2,
		name: "B",
	})

	resFilter := Filter(arr, func(item Test, key int) bool {
		return item.name == "A"
	})
	fmt.Printf("resFilter: %v\n", resFilter) //resFilter: [{1 A}]
	fmt.Printf("----- Filter Test End -----\n\n")
}

func TestSort(t *testing.T) {
	fmt.Printf("----- Sort Test Start -----\n")
	arrTest := []TestFloat{
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
	fmt.Printf("arr: %v\n", arrTest)
	//升序
	Sort(&arrTest, func(a TestFloat, b TestFloat) bool {
		return a.id < b.id
	})
	fmt.Printf("升序 Sort: %v\n", arrTest)
	//降序
	Sort(&arrTest, func(a TestFloat, b TestFloat) bool {
		return a.id > b.id
	})
	fmt.Printf("降序 Sort: %v\n", arrTest)
	fmt.Printf("----- Sort Test End -----\n\n")
}
func TestBubbleSort(t *testing.T) {
	fmt.Printf("----- BubbleSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	BubbleSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- BubbleSort Test End -----\n\n")
}
func TestSelectSort(t *testing.T) {
	fmt.Printf("----- SelectSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	SelectSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- SelectSort Test End -----\n\n")
}
func TestInsertSort(t *testing.T) {
	fmt.Printf("----- InsertSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	InsertSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- InsertSort Test End -----\n\n")
}
func TestShellSort(t *testing.T) {
	fmt.Printf("----- ShellSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	ShellSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- ShellSort Test End -----\n\n")
}
func TestMergeSort(t *testing.T) {
	fmt.Printf("----- MergeSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	MergeSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- MergeSort Test End -----\n\n")
}
func TestQuickSort(t *testing.T) {
	fmt.Printf("----- QuickSort Test Start -----\n")
	arr := []int{3, 44, 38, 5, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Printf("arr: %v\n", arr)
	//升序
	QuickSort(&arr, func(a int, b int) bool {
		return a < b
	})
	fmt.Printf("升序 arr: %v\n", arr)
	fmt.Printf("----- QuickSort Test End -----\n\n")
}
