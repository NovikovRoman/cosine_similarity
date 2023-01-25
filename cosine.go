package cosine_similarity

import (
	"errors"
	"math"
)

func Cosine(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}

	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += b[k] * b[k]
			continue
		}

		if k >= length_b {
			s1 += a[k] * a[k]
			continue
		}

		sumA += a[k] * b[k]
		s1 += a[k] * a[k]
		s2 += b[k] * b[k]
	}

	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("Vectors should not be null (all zeros)")
	}

	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}
