package sorts

import(
	"fmt"
)

// 冒泡排序
func BubbleSort(array []int) {
	l := len(array)
	count := 0
	for i := 0; i < l-1; i++ {
		flag := false
		for j := l-1; j > i; j--{
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				flag = true
				count++
			}
		}
		if !flag{
			break
		}
	}
	fmt.Println("次数：", count)
}