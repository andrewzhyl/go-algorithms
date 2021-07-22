package searches

// 二分查找
func BinarySearch(array []int, key int) int {
	low := 0
	high := len(array) -1
	for(low <= high){
		mid := (low+high) / 2
		if key == array[mid]{
			return mid
		}
		if key < array[mid] {
			high = mid -1
		} else {
			low = mid + 1
		} 
	}
	return -1
}

// 递归二分查找
func IterBinarySearch(array []int, key int, low int, high int) int {
	if high < low || len(array) == 0 {
		return -1
	}
	mid := (low + high) / 2
	if key == array[mid]{
		return mid
	}
	if key < array[mid] {
		return IterBinarySearch(array, key, low, mid - 1)
	} else {
		return IterBinarySearch(array, key, mid + 1, high)
	} 
}

// 插值查找
func InterpolationSearch(array []int, key int) int {
	low := 0
	high := len(array) -1
	for(low <= high){
		mid := low + (high-low) * (key-array[low]) / (array[high] - array[low])
		if key == array[mid]{
			return mid
		}
		if key < array[mid] {
			high = mid -1
		} else {
			low = mid + 1
		} 
	}
	return -1
}