package sorts

import(
	"fmt"
)

// 简单选择排序
func SelectionSort(array []int) {
	l := len(array)
	count := 0
	var min int
	for i := 0; i < l-1; i++ {
		min = i
		for j := i+1; j < l; j++{
			if array[min] > array[j] {
				min = j
				count++
			}
		}
		if(i != min){
			array[i], array[min] = array[min], array[i]
		}
	}
	fmt.Println("次数：", count)
}