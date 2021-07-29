package sorts

import(
	// "fmt"
)

// 堆排序
func HeapSort(array []int) {
	length := len(array)

	// 循环建初始堆
	// 把无序数组构建成最大堆, 这里-2,是因为从索引0开始、另外就是叶子节点【最后一层是不需要堆化的】
	for i := length / 2 - 1; i >= 0; i--{ // 把 L 中的 r 构建成一个大顶堆
		HeapAdjust(array, i, length)
	}

	// 循环删除堆顶元素，并且移到集合尾部，调整堆产生新的堆顶
	for i := length - 1; i > 0; i-- {
		swap(array, 0, i) // 将堆顶记录和当前未经排序子序列的最后一个记录交换
		HeapAdjust(array, 0, i) // 重新调整为大顶堆
	}
}

// 创建大根堆
func HeapAdjust(array []int, root int, length int){
	var temp, child int
	temp = array[root]    // temp 保存父节点的值，用于最后赋值

	//沿关键字较大的孩子结点向下筛选
	for  {  
		child = 2 * root + 1
		if child >= length {
			break 
		}

		// 如果有右孩子，且右孩子大于左孩子的值，则定位到右孩子
        // 这里其实是比较左、右子树的大小，选择更大的
		if child + 1 < length && array[child] < array[child+1]{       
			child++
		}

		// 如果父节点大于任何一个孩子的值，则直接跳出
		if temp >= array[child] { 
			break 
		}  

		 // 当左、右子树比父节点更大，进行交换
		 array[root] = array[child]
		 root = child
	}
	array[root] = temp
}

// 交换
func swap(array []int, i int, j int){
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}