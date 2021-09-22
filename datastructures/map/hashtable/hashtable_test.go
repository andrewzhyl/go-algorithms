package hashtable

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHashTable(t *testing.T) {

	Convey("#TestHashTable", t, func() {
		a := NewHashTable()
		a.Add("name", "andrew")
		a.Add("gender", "male")
		fmt.Println(a)
		a.Remove("gender")
		a.Set("name", "andrewzhang")
		fmt.Println(a)
	})
}
