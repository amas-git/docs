package fp

import "testing"

func TestFib(t *testing.T) {
	var tcase = []struct {
		n int
		e int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}

	for _, tc := range tcase {
		if v := Fib(tc.n); v != tc.e {
			t.Errorf("%v", tc.n)
		}
	}

	for _, tc := range tcase {
		if v := FibIter(tc.n); v != tc.e {
			t.Logf("(%v):%v  ", tc.n, v)
		}
	}
}

func benchmarkFib(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(1, b)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(10, b)
}

func BenchmarkFib100(b *testing.B) {
	benchmarkFib(100, b)
}
func _BenchmarkFibIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIter(i)
	}
}
