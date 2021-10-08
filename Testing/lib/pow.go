package lib

func Pow(x int, y int) (v int) {

	v = 1

	for i := y; i > 0; i-- {
		v = v * x
	}

	return
}
