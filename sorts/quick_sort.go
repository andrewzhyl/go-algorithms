package sorts

// "fmt"

// 快速排序
func QuickSort(items []int, low int, high int) {
	var pivot int
	for low < high {
		pivot = Partition(items, low, high) // 算出枢轴值 pivot
		QuickSort(items, low, pivot-1)      // 对低子表递归排序
		// QuickSort(items, pivot + 1, high)		 // 对高子表递归排序
		low = pivot + 1
	}
}

// 交换顺序表中的记录，使枢轴记录到位，并返回其所在位置
func Partition(items []int, low int, high int) int {
	var pivotkey int
	pivotkey = items[low] // 用子表的第一个记录作为枢轴值
	temp := pivotkey
	for low < high {
		for low < high && items[high] >= pivotkey {
			high--
		}
		// swap(items, low, high)  // 将比枢轴记录小的记录交换到低端
		items[low] = items[high]
		for low < high && items[low] <= pivotkey {
			low++
		}
		// swap(items, low, high)  // 将比枢轴记录大的记录交换到高端
		items[high] = items[low]
	}
	items[low] = temp
	return low // 返回枢轴位置
}
