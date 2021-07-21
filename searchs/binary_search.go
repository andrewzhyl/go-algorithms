package searches

func BinarySearch(array []int, key int) int {
	low := 0
	high := len(array) -1
	for(low <= high){
		mid := (low+high) / 2
		if key < array[mid] {
			high = mid -1
		} else if(key > array[mid]){
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}