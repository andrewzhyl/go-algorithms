package kmp

// from: https://github.com/TheAlgorithms/Go/blob/master/strings/single-string-matching/kmp/kmp.go

import(
	"fmt"
)

const notFoundPosition int = -1

type Result struct {
	resultPosition     int
	numberOfComparison int
}

func Search(text string, word string) Result {
	m, i, c := 0, 0, 0
	t := kmpTable(word)
	for m+i < len(text) {
		fmt.Printf("\n   comparing characters %c %c at positions %d %d", text[m+i], word[i], m+i, i)
		c++
		if word[i] == text[m+i] {
			fmt.Printf(" - match")
			if i == len(word)-1 {
				return Result{
					m, c,
				}
			}
			i++
		} else {
			m = m + i - t[i]
			if t[i] > -1 {
				i = t[i]
			} else {
				i = 0
			}
		}
	}
	return Result{notFoundPosition,
		c,
	}
}

// 计算返回 word 的 next 数组
func GetNext(word string) (next []int){
	next = make([]int, len(word))
	next[0] = 0                  	// next[0] 必然是 0
	x, now := 1, 0             		// 因此从 next[1] 开始求
	for x < len(word){
		if word[x] == word[now] {   // 如果 word[x] == word[now]，则可以向右扩展一位
			now++
			next[x] = now
			x++
		} else if now > 0{
			now = next[now-1] 		// 缩小 now 
		} else {
			next[x] = 0          	// now 已经是 0，无法再缩
			x++
		}
	}
	return next
}