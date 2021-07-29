package sorts

import(
	// "fmt"
)

// 归并排序
func MergingSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergingSort(items[:len(items)/2])
	second := MergingSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}