package sortmultiple

import (
	"fmt"
	"testing"
)

const (
	sortStart = 10
	sortEnd   = 1_000_000
	sortMult  = 10
)

func BenchmarkSort(b *testing.B) {
	for fName, fn := range functions {

		b.Run(fName, func(b *testing.B) {
			for sz := sortStart; sz <= sortEnd; sz *= sortMult {
				b.StopTimer()
				lists := make([][]int, 10)
				for i := 0; i < len(lists); i++ {
					lists[i] = make([]int, sz)
				}
				b.Run(fmt.Sprintf("sz=%d", sz), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						copy := copyLists(lists)
						randomLists(copy)
						b.StartTimer()
						x := fn(copy...)
						if len(x) <= 0 {
							b.Error("error")
						}
					}
				})
			}
		})
	}
}
