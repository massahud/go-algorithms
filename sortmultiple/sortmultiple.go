package sortmultiple

import (
	"math"
	"sort"
)

func SortArrayBigSort(lists ...[]int) []int {
	bigList := CreateBigListWithCopy(lists...)
	sort.Ints(bigList)
	ret := make([]int, 0, len(bigList))
	for _, v := range bigList {
		if v > 0 && (len(ret) == 0 || ret[len(ret)-1] < v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func SortArrayMap(lists ...[]int) []int {

	bigList := CreateBigListWithCopy(lists...)
	m := make(map[int]struct{})
	for _, v := range bigList {
		if v > 0 {
			m[v] = struct{}{}
		}
	}
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func SortArrayMapPreAllocated(lists ...[]int) []int {

	bigList := CreateBigListWithCopy(lists...)
	m := make(map[int]struct{}, len(bigList)*16)
	for _, v := range bigList {
		if v > 0 {
			m[v] = struct{}{}
		}
	}
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func SortMapNoBigArray(lists ...[]int) []int {

	m := make(map[int]struct{})
	for i := range lists {
		for _, v := range lists[i] {
			if v > 0 {
				m[v] = struct{}{}
			}
		}
	}
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func SortMapPreAllocatedNoBigArray(lists ...[]int) []int {

	var sz int
	for i := range lists {
		sz += len(lists[i])
	}
	m := make(map[int]struct{}, sz*16)
	for i := range lists {
		for _, v := range lists[i] {
			if v > 0 {
				m[v] = struct{}{}
			}
		}
	}
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func SortArrayMerge(lists ...[]int) []int {

	var sz int
	for i := range lists {
		sort.Ints(lists[i])
		sz += len(lists[i])
	}
	if sz == 0 {
		return []int{}
	}
	ret := make([]int, 0, sz)
	for {
		min := int(math.MaxInt64)
		listMin := -1
		exit := true
		for i := range lists {
			if len(lists[i]) > 0 {
				exit = false
				if lists[i][0] > 0 {
					if lists[i][0] <= min {
						min = lists[i][0]
						listMin = i
					}
				} else {
					lists[i] = lists[i][1:]
				}
			}
		}
		if exit {
			break
		}
		if listMin < 0 {
			continue
		}
		lists[listMin] = lists[listMin][1:]
		if len(ret) == 0 || ret[len(ret)-1] < min {
			ret = append(ret, min)
		}
	}
	return ret
}

func CreateBigListWithCopy(lists ...[]int) []int {
	var sz int
	for i := range lists {
		sz += len(lists[i])
	}
	bigList := make([]int, sz)
	var p int
	for i := range lists {
		w := len(lists[i])
		copy(bigList[p:p+w], lists[i])
		p += w
	}
	return bigList
}

func CreateBigListWithAppend(lists ...[]int) []int {
	bigList := make([]int, 0)
	for i := range lists {
		bigList = append(bigList, lists[i]...)
	}
	return bigList
}

func CreateBigListWithAppendPreAllocated(lists ...[]int) []int {
	var sz int
	for i := range lists {
		sz += len(lists[i])
	}
	bigList := make([]int, 0, sz)
	for i := range lists {
		bigList = append(bigList, lists[i]...)
	}
	return bigList
}
