package searches

import(
	// "fmt"
)

func FibonacciSearch(array []int, key int) int {
	low := 0
	n := len(array)
	high := n - 1
	k := 0

	// 1. 构建斐波那契数列
	fa := Fibonacci(n)

	// 2. 计算数组长度 n 位于斐波那契数列的位置 k
	for(n > fa[k]){
		k++
	}

	// 3. 根据 fa[k] 的值作为长度，对数组进行填充
	for j := n; j < fa[k]; j++ {
		array = append(array, array[n - 1])
	}

	// 4：对区间不断分割
	for(low <= high){

		// 计算分割位置，左区间长度为 F(n-2)，右区间长度为 F(n-1)
        // -1 是因为下标从 0 开始，即 D[0] 到 D[20] 对应 21 个元素
		mid := low + fa[k - 1] -1 // 计算当前分隔的下标
		if key == array[mid]{

			// 此时搜索到的目标值是填充的值
            if mid < n {
            	return mid
            } else{
                return n - 1
            }
		}
		if key < array[mid] {
			// 若 key 比当前元素小, 则 key 值应该在 low 至 mid - 1 之间，剩下的范围个数为 fa(k-1) - 1
			high = mid -1
			k--
		} else {
			// 若 key 比这个元素大, 则 key 至应该在 mid + 1 至 high 之间，剩下的元素个数为 fa(k) - fa(k-1) - 1 = fa(k-2) - 1
			low = mid + 1
			k -=  2
		} 
	}

	return -1
}

func Fibonacci(n int) []int {
	var items []int
	if n < 1 {
		return items
	}
	items = append(items, 0)
	items = append(items, 1)

	for i := 2; i < n; i++{
		items = append(items, items[i - 1] + items[i - 2])
	}
	return items
}