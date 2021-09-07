package triemap

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrieMap(t *testing.T) {
	words := []string{"apple", "apple", "apple", "orange", "orange", "banana", "strawberry"}
	Convey("#TrieMap", t, func() {
		m := NewTrieMap()
		for _, w := range words {
			node := m.GetNode(w)
			if node != nil {
				m.Set(w, int(node.Value)+1)
			} else {
				m.Add(w, 1)
			}
		}
		fmt.Println("total different words:", m.GetSize())
		fmt.Println("Frequency of apple:", m.Get("apple"))
		fmt.Println("Frequency of orange:", m.Get("orange"))
		fmt.Println("Frequency of banana:", m.Get("banana"))
		fmt.Println("Frequency of strawberry:", m.Get("strawberry"))
	})
}
