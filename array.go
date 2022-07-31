package array

import (
	"math"
)

type Array[T any] []T

func New[T any](args ...T) Array[T] {
	arr := make([]T, 0)
	arr = append(arr, args...)
	return arr
}
func PNew[T any](args ...T) *Array[T] {
	// arr := new(Array[T])
	arr := make(Array[T], 0)
	arr = append(arr, args...)
	return &arr
}

//方法可向数组的末尾添加一个或多个元素，并返回新的长度。
func (arr *Array[T]) Push(args ...T) int {
	*arr = append(*arr, args...)
	return len(*arr)
}

//Pop() 方法用于删除数组的最后一个元素并返回删除的元素。
//注意：此方法改变数组的长度！
func (arr *Array[T]) Pop() (last T, ok bool) {
	if len(*arr) == 0 {
		ok = false
		return
	}
	ok = true
	len := len(*arr)
	last = (*arr)[len-1]
	*arr = (*arr)[:len-1]
	return
}

//Shift() 方法用于把数组的第一个元素从其中删除，并返回第一个元素的值。
//此方法改变数组的长度！
func (arr *Array[T]) Shift() (first T, ok bool) {
	if len(*arr) == 0 {
		ok = false
		return
	}
	ok = true
	first = (*arr)[0]
	*arr = (*arr)[1:]
	return
}

//UnShift() 方法可向数组的开头添加一个或更多元素，并返回新的长度
//此方法改变数组的长度！
func (arr *Array[T]) UnShift(args ...T) int {
	var argsLen = len(args)
	for i := argsLen - 1; i >= 0; i-- {
		*arr = append([]T{args[i]}, *arr...)
	}
	return len(*arr)
}

func (arr *Array[T]) Find(callback func(item T, key int) bool) (res T, ok bool) {
	ok = false
	for i := 0; i < len(*arr); i++ {
		if callback((*arr)[i], i) {
			res = (*arr)[i]
			ok = true
		}
	}
	return
}
func (arr *Array[T]) Filter(callback func(item T, key int) bool) (res Array[T]) {
	res = make(Array[T], 0)
	for i := 0; i < len(*arr); i++ {
		if callback((*arr)[i], i) {
			res = append(res, (*arr)[i])
		}
	}
	return
}
func (arr *Array[T]) Empty() {
	*arr = (*arr)[0:0]
}

func (arr *Array[T]) BrokenEmpty() {
	*arr = make(Array[T], 0)
}

//冒泡排序
func (arr *Array[T]) BubbleSort(callback func(a T, b T) int) {
	var len = len(*arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if callback((*arr)[j], (*arr)[j+1]) > 0 {
				temp := (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

//选择排序
func (arr *Array[T]) SelectSort(callback func(a T, b T) int) {
	var len = len(*arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		tempMinIndex := i
		for j := i + 1; j < len; j++ {
			if callback((*arr)[j], (*arr)[tempMinIndex]) < 0 {
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

//快速排序
/* func (arr *Array[T]) QuickSort(callback func(a T, b T) int) {
	var len = len(*arr)

} */

//插入排序
func (arr *Array[T]) InsertSort(callback func(a T, b T) int) {
	var len = len(*arr)
	for i := 1; i < len; i++ {
		prevIndex := i - 1
		current := (*arr)[i]
		for prevIndex >= 0 && callback((*arr)[prevIndex], current) > 0 {
			(*arr)[prevIndex+1] = (*arr)[prevIndex]
			prevIndex--
		}
		(*arr)[prevIndex+1] = current
	}
}

//希尔排序
func (arr *Array[T]) ShellSort(callback func(a T, b T) int) {
	var len = len(*arr)
	for gap := math.Floor(float64(len / 2)); gap > 0; gap = math.Floor(gap / 2) {
		for i := gap; int(i) < len; i++ {
			j := i
			current := (*arr)[int(i)]
			for j-gap >= 0 && callback(current, (*arr)[int(j-gap)]) < 0 {
				(*arr)[int(j)] = (*arr)[int(j-gap)]
				j = j - gap
			}
			(*arr)[int(j)] = current
		}
	}
}

//归并排序
func (arr *Array[T]) MergeSort(callback func(a T, b T) int) {
	mergeSortSelf[T](arr, 0, len(*arr)-1, callback)
}
func mergeSortSelf[T any](arr *Array[T], start, end int, callback func(a T, b T) int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSortSelf[T](arr, start, mid, callback)
	mergeSortSelf[T](arr, mid+1, end, callback)
	merge[T](arr, start, mid, end, callback)
}
func merge[T any](arr *Array[T], start, mid, end int, callback func(a T, b T) int) {
	rightIndex := start
	leftIndex := mid + 1
	tmpIndex := 0
	tmp := make([]T, 1+end-start)
	for rightIndex <= mid && leftIndex <= end {
		if callback((*arr)[rightIndex], (*arr)[leftIndex]) <= 0 {
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
