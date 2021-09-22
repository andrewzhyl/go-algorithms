package sorts

// 希尔排序
func ShellSort(array []int) {
	var l = len(array)
	// 增量每次都是2
	for step := l/2; step > 0; step /= 2 {
		//从增量那组开始进行插入排序，直至完毕
		for i := step; i < l; i++ { 
			Insert(array, step, i)
		}
	}
}

func Insert(arr []int, step int, i int){
	var j int
	// 获取未排序区间中的第一个元素 
	inserted := arr[i]
	// 开始找插入位置，遍历已排序区间元素,从尾部开始  
	// 拿未排序区间中的第一个元素和已排序区间中的元素依次比较，直到找到适合插入的位置
	for j = i; j - step >= 0 &&  arr[j - step] > inserted; j -= step{
		arr[j] = arr[j - step] // 往后移动元素，留出插入位置
	}
	// 将未排序区间拿出来的元素插入到合适的位置
	arr[j] = inserted
}