package array

type Array[T any] []T

func New[T any]() Array[T] {
	arr := make([]T, 0)
	return arr
}
func PNew[T any]() *Array[T] {
	// arr := new(Array[T])
	arr := make(Array[T], 0)
	return &arr
}
func (arr *Array[T]) Push(content T) {
	*arr = append(*arr, content)
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
