package sorts

import (
	// "fmt"
)

// 直接插入排序
func InsertionSort(array []int) {
	l := len(array)
	var j int
	// 先遍历未排序区间元素，从第二个元素开始
	for i := 1; i < l; i++ { 
		// 获取未排序区间中的第一个元素 
		temporary := array[i]
		// 开始找插入位置，遍历已排序区间元素,从尾部开始  
		// 拿未排序区间中的第一个元素和已排序区间中的元素依次比较，直到找到适合插入的位置
		for j = i; j > 0 && array[j-1] > temporary; j-- {
			array[j] = array[j-1] // 往后移动元素，留出插入位置
		}
		// 将未排序区间拿出来的元素插入到合适的位置
		array[j] = temporary
	}
}