package array

import (
	"math"
	"sort"
)

// As 断言
func As[T any](source []any) (newArr []T) {
	newArr = make([]T, 0, len(source))
	for i := 0; i < len(source); i++ {
		if s, ok := source[i].(T); ok {
			newArr = append(newArr, s)
		}
	}
	return
}

// Push 方法可向数组的末尾添加一个或多个元素，并返回新的长度。
func Push[T any](arr *[]T, args ...T) int {
	*arr = append(*arr, args...)
	return len(*arr)
}

// Pop 方法用于删除数组的最后一个元素并返回删除的元素。
// 注意：此方法改变数组的长度！
func Pop[T any](arr *[]T) (last T, ok bool) {
	if len(*arr) == 0 {
		ok = false
		return
	}
	ok = true
	arrLen := len(*arr)
	last = (*arr)[arrLen-1]
	*arr = (*arr)[:arrLen-1]
	return
}

// Unique 去重
func Unique() {

}

// Splice
/**
* Splice(index int, howMany int, args ...T)
* @param index 规定从何处添加或删除元素，该参数是插入元素或删除元素的起始下标，必须是整数
* @param howMany 规定应该删除多少元素
    howMany 为0时，表示不删除，如果 args有值则在index处插入args
    howMany 为负数时，相当于0，不删除
    howMany 为正数时，表示删除数量，删除后，如果 args有值则在index处插入args
* @param args 要在index处添加的多个元素
* 用于添加或删除数组中的元素
* 会改变原始数组
* 返回的是含有被删除的元素的数组
*/
func Splice[T any](arr *[]T, index int, howMany int, args ...T) (delArr []T) {
	oldArrLen := len(*arr)
	argsLen := len(args)
	delArr = make([]T, 0)
	if index < 0 {
		index = index + oldArrLen
		if index < 0 {
			index = 0
		}
	}
	if howMany < 0 {
		howMany = 0
	}
	behindSegmentLen := oldArrLen - index //原始数组arr 从index开始（包括index）到最后的内容段

	if howMany > behindSegmentLen { //要删除的数量大于可删除的数量
		howMany = behindSegmentLen
	}
	if howMany == 0 && argsLen == 0 {
		*arr = append(*arr, args...)
		return
	}
	if index > oldArrLen-1 {
		*arr = append((*arr), args...)
		return
	}
	/* 1 2 3 4 5 6 7 8 9 10
	Splice(3,0,11,12,13)
	1 2 3     4 5 6 7 8 9 10    add cap[11 12 13] 扩容 3次
	1 2 3     4 5 6    4 5 6 7 8 9 10 移位
	1 2 3     11 12 13     4 5 6 7 8 9 10 */

	if howMany == 0 { //不删除
		for i := 0; i < argsLen; i++ {
			*arr = append(*arr, args[i]) //先实现扩容，用args[i]临时占位，下面通过移位实现交换
		}
		newArrLen := len(*arr)
		for i := 0; i < behindSegmentLen; i++ {
			(*arr)[newArrLen-i-1] = (*arr)[index+behindSegmentLen-i-1] //把前面的往末尾移位
		}
		for i := 0; i < argsLen; i++ {
			(*arr)[i+index] = args[i] //把args里的内容覆盖进来
		}
	} else { //删除
		/* 1 2 3 4 5 6 7 8 9 10
		Splice(3,2,11,12,13)
		1 2 3     to del[4 5] 6 7 8 9 10    删除2个加3个 add cap[11]扩容1个就够了
		1 2 3     to del[4 5] 6 7 8 9    9 10
		1 2 3     11 12 13   9 10 */

		/* 1 2 3 4 5 6 7 8 9 10
		Splice(3,2,11,12)
		1 2 3     to del[4 5] 6 7 8 9 10    删除2个加2个 不用扩容，也不用移位，直接用args覆盖
		1 2 3     11 12     6 7 8 9 10 */

		/* 1 2 3 4 5 6 7 8 9 10
		Splice(3,3,11,12)
		1 2 3     to del[4 5 6] 7 8 9 10    删除3个减2个 不用扩容 空间会多 需要从后向前移位，再把最后几位多的扔掉
		1 2 3     11 12  7 8 9 10 */

		/* 1 2 3 4 5 6 7 8 9 10
		Splice(3,4,11,12)
		1 2 3     to del[4 5 6 7] 8 9 10    删除4个减2个 不用扩容 空间会多 需要从后向前移位，再把最后几位多的扔掉
		1 2 3     11 12  8 9 10 */
		capExLen := howMany - argsLen //扩容长度

		if capExLen < 0 { //需要扩容
			capExLen = capExLen * (-1)
			for i := 0; i < capExLen; i++ {
				*arr = append(*arr, args[i]) //扩容时，用args[i]临时占位，下面通过移位实现交换
			}
			newArrLen := len(*arr)
			for i := 0; i < behindSegmentLen-howMany; i++ {
				(*arr)[newArrLen-i-1] = (*arr)[index+behindSegmentLen-i-1] //把前面的往末尾移位
			}
			for i := 0; i < howMany; i++ {
				delArr = append(delArr, (*arr)[i+index])
			}
			for i := 0; i < argsLen; i++ {
				(*arr)[i+index] = args[i] //把args里的内容覆盖进来
			}
		} else { //不用扩容
			for i := 0; i < howMany; i++ {
				delArr = append(delArr, (*arr)[i+index])
			}
			for i := 0; i < argsLen; i++ {
				(*arr)[i+index] = args[i] //把args里的内容覆盖进来
			}

			if capExLen > 0 { //减多加少

				for i := 0; i < behindSegmentLen-howMany; i++ { //把后面的移动往前移
					(*arr)[i+index+argsLen] = (*arr)[i+index+argsLen+howMany-argsLen] //把后面的移动往前移
				}
				*arr = (*arr)[:oldArrLen-capExLen] //截掉后面
			}
		}
	}
	return
}

// Slice 从已有的数组中返回选定区间的新元素数组
// 此方法不会对源数组产生影响（原生切片因扩容规则：不扩容的情况下，会对源切片产生影响）
// 如果你不喜欢此方法你依然可以使用原生切片截取方式[:]来操作，但要注意扩容规则
func Slice[T any](arr []T, start, end int) (newArr []T) {
	arrLen := len(arr)
	newArr = make([]T, 0)
	if start < 0 {
		start = start + arrLen
		if start < 0 {
			start = 0
		}
	}
	if end > arrLen {
		end = arrLen
	}
	if end < 0 {
		end = end + arrLen
	}
	if end-start > 0 {
		for i := start; i < end; i++ {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// Shift 方法用于把数组的第一个元素从其中删除，并返回第一个元素的值。
// 此方法改变数组的长度！
func Shift[T any](arr *[]T) (first T, ok bool) {
	if len(*arr) == 0 {
		ok = false
		return
	}
	ok = true
	first = (*arr)[0]
	*arr = (*arr)[1:]
	return
}

// UnShift 方法可向数组的开头添加一个或更多元素，并返回新的长度
// 此方法改变数组的长度！
func UnShift[T any](arr *[]T, args ...T) int {
	var argsLen = len(args)
	for i := argsLen - 1; i >= 0; i-- {
		*arr = append([]T{args[i]}, *arr...)
	}
	return len(*arr)
}

//返回一个新数组，数组中的元素为原始数组元素调用函数处理后的值
//方法按照原始数组元素顺序依次处理元素
//不会改变原始数组

func Map[T1 any, T2 any](arr []T1, callback func(item T1, index int) T2) (newArr []T2) {
	newArr = make([]T2, 0)
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		newArr = append(newArr, callback(arr[i], i))
	}
	return
}

// ForEach 方法用于调用数组的每个元素，并将元素传递给回调函数。
// 注意: ForEach() 对于空数组是不会执行回调函数的
func ForEach[T any](arr []T, callback func(item *T, index int)) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		callback(&arr[i], i)
	}
}

// Every 用于检测数组所有元素是否都符合指定条件（通过函数提供）
// 如果数组中检测到有一个元素不满足，则整个表达式返回 false，都满足时，返回true
// 注：如何是空数组，直接返回false 这里与js里不一样。
// 不会改变原始数组
func Every[T any](arr []T, callback func(item T, index int) bool) (res bool) {
	arrLen := len(arr)
	if arrLen == 0 {
		res = false
		return
	}
	res = true
	for i := 0; i < arrLen; i++ {
		if !callback(arr[i], i) {
			res = false
		}
	}
	return
}

// Some 用于检测数组中的元素是否满足指定条件（函数提供），只要有一个满足条件，就返回true
// 如果没有满足条件的元素，则返回false
// 如何是空数组，直接返回false
// 不会改变原始数组
func Some[T any](arr []T, callback func(item T, index int) bool) (res bool) {
	arrLen := len(arr)
	if arrLen == 0 {
		res = false
		return
	}
	res = false
	for i := 0; i < arrLen; i++ {
		if callback(arr[i], i) {
			res = true
			return
		}
	}
	return
}

// Find 从前向后遍历
func Find[T any](arr []T, callback func(item T, index int) bool) (res T, ok bool) {
	ok = false
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if callback(arr[i], i) {
			res = arr[i]
			ok = true
			return
		}
	}
	return
}

// FindLast 从后向前遍历
func FindLast[T any](arr []T, callback func(item T, index int) bool) (res T, ok bool) {
	ok = false
	arrLen := len(arr)
	for i := arrLen - 1; i >= 0; i-- {
		if callback(arr[i], i) {
			res = arr[i]
			ok = true
			return
		}
	}
	return
}

// FindIndex 返回符合传入回调函数条件的第一个元素索引位置
// 如果没有符合条件的元素返回 -1
// 从前向后遍历
func FindIndex[T any](arr []T, callback func(item T, index int) bool) (firstIndex int) {
	firstIndex = -1
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if callback(arr[i], i) {
			firstIndex = i
			return
		}
	}
	return
}

// FindLastIndex
// 与FindIndex不同的是，从后向前遍历
func FindLastIndex[T any](arr []T, callback func(item T, index int) bool) (firstIndex int) {
	firstIndex = -1
	arrLen := len(arr)
	for i := arrLen - 1; i >= 0; i-- {
		if callback(arr[i], i) {
			firstIndex = i
			return
		}
	}
	return
}

// Filter 根据条件过滤 返回结果依然是一个数组，如果没有匹配项，则返回空数组
func Filter[T any](arr []T, callback func(item T, index int) bool) (res []T) {
	res = make([]T, 0)
	for i := 0; i < len(arr); i++ {
		if callback(arr[i], i) {
			res = append(res, arr[i])
		}
	}
	return
}
func Empty[T any](arr *[]T) {
	*arr = (*arr)[0:0]
}

func BrokenEmpty[T any](arr *[]T) {
	*arr = make([]T, 0)
}

// Sort golang原生排序
func Sort[T any](arr *[]T, callback func(a T, b T) bool) {
	sort.Slice(*arr, func(i, j int) bool {
		return callback((*arr)[i], (*arr)[j])
	})
}

// BubbleSort 冒泡排序
func BubbleSort[T any](arr *[]T, callback func(a T, b T) bool) {
	var arrLen = len(*arr)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if !callback((*arr)[j], (*arr)[j+1]) {
				temp := (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort[T any](arr *[]T, callback func(a T, b T) bool) {
	var arrLen = len(*arr)
	for i := 0; i < arrLen-1; i++ {
		minIndex := i
		tempMinIndex := i
		for j := i + 1; j < arrLen; j++ {
			if callback((*arr)[j], (*arr)[tempMinIndex]) {
				tempMinIndex = j
			}
		}
		if tempMinIndex != minIndex { //有变化才交换
			minIndex = tempMinIndex
			temp := (*arr)[minIndex]
			(*arr)[minIndex] = (*arr)[i]
			(*arr)[i] = temp
		}
	}
}

// QuickSort 快速排序
func QuickSort[T any](arr *[]T, callback func(a T, b T) bool) {
	quickSortSelf(arr, 0, len(*arr), callback)
}
func quickSortSelf[T any](arr *[]T, left, right int, callback func(a T, b T) bool) {
	if left < right {
		pivot := (*arr)[left]
		j := left
		for i := left; i < right; i++ {
			if callback((*arr)[i], pivot) {
				j++
				(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			}
		}
		(*arr)[left], (*arr)[j] = (*arr)[j], (*arr)[left]
		quickSortSelf(arr, left, j, callback)
		quickSortSelf(arr, j+1, right, callback)
	}
}

// InsertSort 插入排序
func InsertSort[T any](arr *[]T, callback func(a T, b T) bool) {
	var arrLen = len(*arr)
	for i := 1; i < arrLen; i++ {
		prevIndex := i - 1
		current := (*arr)[i]
		for prevIndex >= 0 && !callback((*arr)[prevIndex], current) {
			(*arr)[prevIndex+1] = (*arr)[prevIndex]
			prevIndex--
		}
		(*arr)[prevIndex+1] = current
	}
}

// ShellSort 希尔排序
func ShellSort[T any](arr *[]T, callback func(a T, b T) bool) {
	var arrLen = len(*arr)
	for gap := math.Floor(float64(arrLen / 2)); gap > 0; gap = math.Floor(gap / 2) {
		for i := gap; int(i) < arrLen; i++ {
			j := i
			current := (*arr)[int(i)]
			for j-gap >= 0 && callback(current, (*arr)[int(j-gap)]) {
				(*arr)[int(j)] = (*arr)[int(j-gap)]
				j = j - gap
			}
			(*arr)[int(j)] = current
		}
	}
}

// MergeSort 归并排序
func MergeSort[T any](arr *[]T, callback func(a T, b T) bool) {
	mergeSortSelf[T](arr, 0, len(*arr)-1, callback)
}
func mergeSortSelf[T any](arr *[]T, start, end int, callback func(a T, b T) bool) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSortSelf[T](arr, start, mid, callback)
	mergeSortSelf[T](arr, mid+1, end, callback)
	merge[T](arr, start, mid, end, callback)
}
func merge[T any](arr *[]T, start, mid, end int, callback func(a T, b T) bool) {
	rightIndex := start
	leftIndex := mid + 1
	tmpIndex := 0
	tmp := make([]T, 1+end-start)
	for rightIndex <= mid && leftIndex <= end {
		if callback((*arr)[rightIndex], (*arr)[leftIndex]) {
			tmp[tmpIndex] = (*arr)[rightIndex]
			tmpIndex++
			rightIndex++
		} else {
			tmp[tmpIndex] = (*arr)[leftIndex]
			tmpIndex++
			leftIndex++
		}
	}
	var appendStart, appendEnd int
	if rightIndex > mid {
		appendStart = leftIndex
		appendEnd = end
	} else {
		appendStart = rightIndex
		appendEnd = mid
	}
	for appendStart <= appendEnd {
		tmp[tmpIndex] = (*arr)[appendStart]
		tmpIndex++
		appendStart++
	}
	var tempLen = len(tmp)
	for i := 0; i < tempLen; i++ {
		(*arr)[start+i] = tmp[i]
	}
}
