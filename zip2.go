package slicezip

type Tuple2[T1, T2 any] struct {
	First  T1
	Second T2
}

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
