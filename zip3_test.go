package slicezip

import (
	"strconv"
	"testing"
)

func TestZip3(t *testing.T) {
	names := []string{
		"Leia Organa",
		"Luke Skywalker",
		"Sheev Palpatine",
		"Han Solo",
	}

	numbers := []int{1, 2, 3, 4}

	chars := []byte{'L', 'L', 'S', 'H'}

	// Test each element
	for i, z := range Zip3(names, numbers, chars) {
		t.Run("Row "+strconv.Itoa(i), func(t *testing.T) {
			// Test the first element
			if z.First != names[i] {
				t.Errorf("Expected %s, got %s", names[i], z.First)
			}

			// Test the second element
			if z.Second != numbers[i] {
				t.Errorf("Expected %d, got %d", numbers[i], z.Second)
			}

			// Test the third element
			if z.Third != chars[i] {
				t.Errorf("Expected %c, got %c", chars[i], z.Third)
			}
		})
	}
}
