// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdpdqsort

// Sort sorts data.
// The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	limit := usize - bitsLeadingZeros(uint(n))
	recurse(data, 0, n, 0, false, limit)
}

// recurse sorts `data` recursively.
// The algorithm based on pattern-defeating quicksort(pdqsort), but without the optimizations from BlockQuicksort.
// pdqsort paper: https://arxiv.org/pdf/2106.05123.pdf
func recurse(data Interface, a, b, pred int, predExist bool, limit int) {
	const MaxInsertion = 12

	var (
		wasBalanced    = true // whether the last partitioning was reasonably balanced
		wasPartitioned = true // whether the slice was already partitioned
	)

	for {
		length := b - a

		if length <= MaxInsertion {
			insertionSort(data, a, b)
			return
		}

		// Fall back to heapsort if too many bad choices were made.
		if limit == 0 {
			heapSort(data, a, b)
			return
		}

		// If the last partitioning was imbalanced, we need to breaking patterns.
		if !wasBalanced {
			breakPatterns(data, a, b)
			limit--
		}

		pivotidx, likelySorted := choosePivot(data, a, b)

		// The slice is likely already sorted.
		if wasBalanced && wasPartitioned && likelySorted {
			if partialInsertionSort(data, a, b) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if predExist && !data.Less(pred, pivotidx) {
			mid := partitionEqual(data, a, b, pivotidx)
			a = mid
			continue
		}

		mid, wasP := partition(data, a, b, pivotidx)
		wasPartitioned = wasP

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8
		if leftLen > rightLen {
			wasBalanced = rightLen >= balanceThreshold
			recurse(data, a, mid, pred, predExist, limit)
			a = mid + 1
			pred = mid
			predExist = true
		} else {
			wasBalanced = leftLen >= balanceThreshold
			recurse(data, mid+1, b, mid, true, limit)
			b = mid
		}
	}
}

func partition(data Interface, a, b, pivotidx int) (newpivotidx int, wasPartitioned bool) {
	data.Swap(a, pivotidx)
	i, j := a+1, b-1

	for i <= j && data.Less(i, a) {
		i++
	}
	for i <= j && !data.Less(j, a) {
		j--
	}
	if i > j {
		data.Swap(j, a)
		return j, true
	}
	data.Swap(i, j)
	i++
	j--

	for {
		for i <= j && data.Less(i, a) {
			i++
		}
		for i <= j && !data.Less(j, a) {
			j--
		}
		if i > j {
			break
		}
		data.Swap(i, j)
		i++
		j--
	}
	data.Swap(j, a)
	return j, false
}

// partitionEqual partitions `data` into elements equal to `data[pivotidx]` followed by elements greater than `data[pivotidx]`.
// It assumed that `data` does not contain elements smaller than the `data[pivotidx]`.
func partitionEqual(data Interface, a, b, pivotidx int) int {
	data.Swap(a, pivotidx)

	L := a + 1
	R := b
	for {
		for L < R && !data.Less(a, L) {
			L++
		}
		for L < R && data.Less(a, R-1) {
			R--
		}
		if L >= R {
			break
		}
		R--
		data.Swap(L, R)
		L++
	}
	return L
}

// partialInsertionSort partially sorts a slice, returns `true` if the slice is sorted at the end.
func partialInsertionSort(data Interface, a, b int) bool {
	const (
		MaxSteps         = 5  // maximum number of adjacent out-of-order pairs that will get shifted
		ShortestShifting = 50 // don't shift any elements on short arrays
	)
	i := a + 1
	for j := 0; j < MaxSteps; j++ {
		for i < b && !data.Less(i, i-1) {
			i++
		}

		if i == b {
			return true
		}

		if b-a < ShortestShifting {
			return false
		}

		data.Swap(i-1, i)

		// Shift the smaller one to the left.
		if i-a >= 2 {
			for j := i - 1; j >= 1; j-- {
				if !data.Less(j, j-1) {
					break
				}
				data.Swap(j, j-1)
			}
		}
		// Shift the greater one to the right.
		if b-i >= 2 {
			for j := 1; j < b; j++ {
				if !data.Less(j, j-1) {
					break
				}
				data.Swap(j, j-1)
			}
		}
	}
	return false
}

// breakPatterns scatters some elements around in an attempt to break some patterns
// that might cause imbalanced partitions in quicksort.
func breakPatterns(data Interface, a, b int) {
	length := b - a
	if length >= 8 {
		// Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
		random := uint(length)
		random ^= random << 13
		random ^= random >> 17
		random ^= random << 5

		modulus := nextPowerOfTwo(length)
		pos := a + length/8

		for i := 0; i < 3; i++ {
			other := int(random & (modulus - 1))
			if other >= length {
				other -= length
			}
			data.Swap(pos-1+i, a+other)
		}
	}
}

// choosePivot chooses a pivot in `data`.
//
// `data` might be reordered in this function.
//
// [0,8): choose a static pivot.
// [8,ShortestNinther): use the simple median-of-three method.
// [ShortestNinther,∞): use the Tukey’s ninther method.
func choosePivot(data Interface, lo, hi int) (pivotidx int, likelySorted bool) {
	const (
		ShortestNinther = 50
		MaxSwaps        = 4 * 3
	)

	l := hi - lo

	var (
		swaps int
		a     = lo + l/4*1
		b     = lo + l/4*2
		c     = lo + l/4*3
	)

	if l >= 8 {
		if l >= ShortestNinther {
			// Tukey’s ninther method.
			sortAdjacent(data, &a, &swaps)
			sortAdjacent(data, &b, &swaps)
			sortAdjacent(data, &c, &swaps)
		}
		// Find the median among `a`, `b`, `c` and stores it into `b`.
		sort3(data, &a, &b, &c, &swaps)
	}

	if swaps < MaxSwaps {
		return b, (swaps == 0)
	} else {
		// The maximum number of swaps was performed.
		// Reversing will probably help.
		reverseRange(data, lo, hi)
		return 2*lo + (l - 1 - b), true
	}
}

// sort2 swaps `a, b` so that `data[a] <= data[b]`.
func sort2(data Interface, a, b, swaps *int) {
	if data.Less(*b, *a) {
		*a, *b = *b, *a
		*swaps++
	}
}

// sort3 swaps `a, b, c` so that `data[a] <= data[b] <= data[c]`.
func sort3(data Interface, a, b, c, swaps *int) {
	sort2(data, a, b, swaps)
	sort2(data, b, c, swaps)
	sort2(data, a, b, swaps)
}

// sortAdjacent finds the median of `data[a - 1], data[a], data[a + 1]` and stores the index into `a`.
func sortAdjacent(data Interface, a, swaps *int) {
	t1 := *a - 1
	t2 := *a + 1
	sort3(data, &t1, a, &t2, swaps)
}

func reverseRange(data Interface, a, b int) {
	i := a
	j := b - a - 1
	for i < j {
		data.Swap(i, j)
		i++
		j--
	}
}

func nextPowerOfTwo(length int) uint {
	shift := uint(usize - bitsLeadingZeros(uint(length)))
	return uint(1 << shift)
}
