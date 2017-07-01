package fibonacci_closure

type generator func() int

func fibonacci() generator {
	a, b := 1, 0
	return func() int {
		a, b = b, a + b
		return a
	}
}

func GetFibonacciList(length int) (list []int) {
	f := fibonacci()
	for i := 0; i < length; i++ {
		list = append(list, f())
	}
	return list
}