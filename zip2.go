package slicezip

// Tuple2 is a container for 2 items. It is used for the
// resulting slice of Zip2.
type Tuple2[T1, T2 any] struct {
	First  T1
	Second T2
}

// Zip2 zips two slices of length n into an slice of n Tuple2 elements
// where slice[i].First = s1[i] and slice[i].Second = s2[i]
//
// If the slices have different lengths, the shortest one will be used.
func Zip2[T1, T2 any](s1 []T1, s2 []T2) []Tuple2[T1, T2] {
	// Use the smallest one
	size := len(s1)
	if len(s2) < size {
		size = len(s2)
	}

	zipped := make([]Tuple2[T1, T2], size)

	for i := 0; i < size; i++ {
		zipped[i].First = s1[i]
		zipped[i].Second = s2[i]
	}

	return zipped
}
