package trie

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrie(t *testing.T) {

	Convey("#Add & #Contains", t, func() {
		trie := NewTrie()
		words := []string{"sam", "john", "tim", "jose", "rose",
			"cat", "dog", "dogg", "roses"}
		for i := 0; i < len(words); i++ {
			trie.Add(words[i])
		}

		wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
			"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
		result := []bool{
			true, true, true, true, true, true, true, true, true, false, false, false,
		}

		for i := 0; i < len(wordsToFind); i++ {
			found := trie.Contains(wordsToFind[i])
			So(trie.Contains(wordsToFind[i]), ShouldEqual, result[i])
			if found {
				fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
			} else {
				fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
			}
		}
	})
}
