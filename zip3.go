package slicezip

type Tuple3[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

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
