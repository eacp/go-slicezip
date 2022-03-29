package slicezip

// Tuple3 is a container for 3 items. It is used for the
// resulting slice of Zip3.
type Tuple3[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

// Zip3 zips two slices of length n into an slice of n Tuple3 elements
// where slice[i].First = s1[i], slice[i].Second = s2[i]
// and slice[i].Third = s3[i]
//
// If the slices have different lengths, the shortest one will be used.
func Zip3[T1, T2, T3 any](s1 []T1, s2 []T2, s3 []T3) []Tuple3[T1, T2, T3] {
	// Use the smallest one
	size := len(s1)
	if len(s2) < size {
		size = len(s2)
	}
	if len(s3) < size {
		size = len(s3)
	}

	zipped := make([]Tuple3[T1, T2, T3], size)

	for i := 0; i < size; i++ {
		zipped[i].First = s1[i]
		zipped[i].Second = s2[i]
		zipped[i].Third = s3[i]
	}

	return zipped
}
