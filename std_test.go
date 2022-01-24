// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdpdqsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
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
	randomTestTimes := rand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(100)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
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
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		v3 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
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
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
		}
		pivotidx := rand.Intn(len(v1))
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
