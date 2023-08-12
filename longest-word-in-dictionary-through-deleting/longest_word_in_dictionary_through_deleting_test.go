package longest_word_in_dictionary_through_deleting

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortedString(t *testing.T) {
	s := []string{"ab", "ac", "ad", "aac", "aa"}
	sort.Strings(s)
	fmt.Println(s)
}
