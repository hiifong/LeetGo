package sort_characters_by_frequency

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	s := "aassssffssa"
	fmt.Println(frequencySort(s))
}
