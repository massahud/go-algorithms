package sortmultiple

import (
	"fmt"
	"testing"
)

const (
	sortStart = 10
	sortEnd   = 10_000_000
	sortMult  = 10
)

func BenchmarkSortArrayBigSort(b *testing.B) {
	for sz := sortStart; sz <= sortEnd; sz *= sortMult {
		b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
			b.StopTimer()
			lists := make([][]int, 10)
			for i := 0; i < len(lists); i++ {
				lists[i] = make([]int, sz)
			}
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				randomLists(lists)
				b.StartTimer()
				x := SortArrayBigSort(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkSortArrayMap(b *testing.B) {
	for sz := sortStart; sz <= sortEnd; sz *= sortMult {
		b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
			b.StopTimer()
			lists := make([][]int, 10)
			for i := 0; i < len(lists); i++ {
				lists[i] = make([]int, sz)
			}
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				randomLists(lists)
				b.StartTimer()
				x := SortArrayMap(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkSortArrayMapPreallocated(b *testing.B) {
	for sz := sortStart; sz <= sortEnd; sz *= sortMult {
		b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
			b.StopTimer()
			lists := make([][]int, 10)
			for i := 0; i < len(lists); i++ {
				lists[i] = make([]int, sz)
			}
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				randomLists(lists)
				b.StartTimer()
				x := SortArrayMapPreAllocated(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkSortMapNoBigArray(b *testing.B) {
	for sz := sortStart; sz <= sortEnd; sz *= sortMult {
		b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
			b.StopTimer()
			lists := make([][]int, 10)
			for i := 0; i < len(lists); i++ {
				lists[i] = make([]int, sz)
			}
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				randomLists(lists)
				b.StartTimer()
				x := SortMapNoBigArray(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkSortMapPreAllocatedNoBigArray(b *testing.B) {
	for sz := sortStart; sz <= sortEnd; sz *= sortMult {
		b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
			b.StopTimer()
			lists := make([][]int, 10)
			for i := 0; i < len(lists); i++ {
				lists[i] = make([]int, sz)
			}
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				randomLists(lists)
				b.StartTimer()
				x := SortMapPreAllocatedNoBigArray(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}
