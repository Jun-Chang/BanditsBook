package algorithms


func maxInt64(is []int64) (int, int64) {
	var i int
	var m int64
	for j, v := range is {
		if v > m {
			i = j
			m = v
		}
	}
	return i, m
}
func maxFloat64(fs []float64) (int, float64) {
	var i int
	var m float64
	for j, v := range fs {
		if v > m {
			i = j
			m = v
		}
	}
	return i, m
}

func sumInt64(is []int64) int64 {
	var s int64
	for _, i := range is {
		s += i
	}
	return s
}

func sumFloat64(fs []float64) float64 {
	var s float64
	for _, f := range fs {
		s += f
	}
	return s
}
