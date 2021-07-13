package sortmultiple

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	bigStart = 10_000
	bigEnd   = 10_000_000
	bigMult  = 10
)

func BenchmarkCreateBigListWithAppend(b *testing.B) {
	for sz := bigStart; sz <= bigEnd; sz *= bigMult {
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
				x := CreateBigListWithAppend(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkCreateBigListWithAppendPreAllocated(b *testing.B) {
	for sz := bigStart; sz <= bigEnd; sz *= bigMult {
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
				x := CreateBigListWithAppendPreAllocated(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func BenchmarkCreateBigListWithCopy(b *testing.B) {
	for sz := bigStart; sz <= bigEnd; sz *= bigMult {
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
				x := CreateBigListWithCopy(lists...)
				if len(x) <= 0 {
					b.Error("never")
				}
			}
		})
	}
}

func randomLists(lists [][]int) {
	for i := 0; i < len(lists); i++ {
		for j := 0; j < len(lists[i]); j++ {
			lists[i][j] = rand.Int()
		}
	}
}
