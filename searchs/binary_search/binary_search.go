package searches

// "fmt"

// 二分查找
func BinarySearch(array []int, key int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		if key < array[mid] {
			right = mid - 1
		} else if key > array[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 递归二分查找
func IterBinarySearch(array []int, key int, low int, high int) int {
	if low > high || len(array) == 0 {
		return -1
	}
	mid := (low + high) / 2
	if key < array[mid] {
		return IterBinarySearch(array, key, low, mid-1)
	} else if key > array[mid] {
		return IterBinarySearch(array, key, mid+1, high)
	} else {
		return mid
	}
}

// 插值查找
func InterpolationSearch(array []int, key int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)*(key-array[low])/(array[high]-array[low])
		if key == array[mid] {
			return mid
		}
		if key < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
