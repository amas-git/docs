package fp

func Fib(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func FibIter(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		n1 := 1
		n2 := 2
		for i := 2; i < n; i++ {
			n1, n2 = n2, n2+n1
		}
		return n2
	}
}
