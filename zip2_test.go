package slicezip

import (
	"strconv"
	"testing"
)

func TestZip2(t *testing.T) {
	names := []string{
		"Leia Organa",
		"Luke Skywalker",
		"Sheev Palpatine",
		"Han Solo",
	}

	numbers := []int{1, 2, 3, 4}

	// Test each element
	for i, z := range Zip2(names, numbers) {
		t.Run("Row "+strconv.Itoa(i), func(t *testing.T) {
			// Test the first element
			if z.First != names[i] {
				t.Errorf("Expected %s, got %s", names[i], z.First)
			}

			// Test the second element
			if z.Second != numbers[i] {
				t.Errorf("Expected %d, got %d", numbers[i], z.Second)
			}
		})
	}
}
