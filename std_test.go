package stdpdqsort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/zhangyunhao116/fastrand"
)

func TestAll(t *testing.T) {
	fuzzTestSort(t, func(data []int) {
		insertionSort(IntSlice(data), 0, len(data))
	})
	fuzzTestSort(t, func(data []int) {
		heapSort(IntSlice(data), 0, len(data))
	})
	fuzzTestSort(t, func(data []int) {
		Ints(data)
	})
	fuzzTestPartition(t, func(data []int, pivotidx int) int {
		idx, _ := partition(IntSlice(data), 0, len(data), pivotidx)
		return idx
	})
}

func TestPartialInsertionSort(t *testing.T) {
	randomTestTimes := fastrand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := fastrand.Intn(100)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := fastrand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
		}
		sort.Ints(v1)
		if partialInsertionSort(IntSlice(v2), 0, len(v2)) {
			for idx := range v1 {
				if v1[idx] != v2[idx] {
					t.Fatal("invalid sort:", idx, v1[idx], v2[idx])
				}
			}
		}
	}
}

func fuzzTestSort(t *testing.T, f func(data []int)) {
	const times = 2048
	randomTestTimes := fastrand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := fastrand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		v3 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := fastrand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
			v3[j] = randomValue
		}
		sort.Ints(v1)
		f(v2)
		for idx := range v1 {
			if v1[idx] != v2[idx] {
				t.Fatal("invalid sort:", idx, v1[idx], v2[idx], fmt.Sprintf("\n%#v\n", v3))
			}
		}
	}
}

func fuzzTestPartition(t *testing.T, f func(data []int, pivotidx int) int) {
	const times = 2048
	randomTestTimes := fastrand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := fastrand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := fastrand.Intn(randomLenth)
			v1[j] = randomValue
		}
		pivotidx := fastrand.Intn(len(v1))
		newpivotidx := f(v1, pivotidx)
		pivot := v1[newpivotidx]
		for i, v := range v1 {
			if i < newpivotidx && v > pivot {
				t.Fatal(i, v, pivotidx, pivot, fmt.Sprintf("\n%v\n", v1))
			}
			if i > newpivotidx && v < pivot {
				t.Fatal(i, v, pivotidx, pivot, fmt.Sprintf("\n%v\n", v1))
			}
		}
	}
}
