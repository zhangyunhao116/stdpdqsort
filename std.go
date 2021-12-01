package stdpdqsort

import (
	"math/bits"
	"strconv"
)

// Sort sorts data.
// The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	limit := strconv.IntSize - bits.LeadingZeros(uint(n))
	recurse(data, 0, n, 0, false, limit)
}

// recurse sorts `v` recursively.
// The algorithm base on pattern-defeating quicksort(pdqsort), but without the optimizations from block quciksort.
// pdqsort paper: https://arxiv.org/pdf/2106.05123.pdf
func recurse(data Interface, a, b, pred int, predExist bool, limit int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true // whether the last partitioning was reasonably balanced.
		wasPartitioned = true // whether the slice looks like already partitioned.
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSort(data, a, b)
			return
		}

		// Fall back to heapsort if too many bad choices were made.
		if limit == 0 {
			heapSort(data, a, b)
			return
		}

		// If the last partitioning was imbalanced, try breaking patterns.
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
		if b-a-mid < mid {
			wasBalanced = (b - a - mid) >= b-a/8
		} else {
			wasBalanced = mid >= b-a/8
		}
		wasPartitioned = wasP

		if mid-a > b-mid {
			recurse(data, a, mid, pred, predExist, limit)
			a = mid + 1
			pred = mid
			predExist = true
		} else {
			recurse(data, mid+1, b, mid, true, limit)
			b = mid
		}
	}
}

func partition(v Interface, a, b, pivotidx int) (newpivotidx int, wasPartitioned bool) {
	v.Swap(a, pivotidx)
	i, j := a+1, b-1

	for {
		for i <= j && v.Less(i, a) {
			i++
		}
		for i <= j && !v.Less(j, a) {
			j--
		}
		if i > j {
			break
		}
		v.Swap(i, j)
		i++
		j--
	}
	v.Swap(j, a)
	return j, j == pivotidx
}

// partitionEqual partitions `v` into elements equal to `v[pivotidx]` followed by elements greater than `v[pivotidx]`.
// It assumed that `v` does not contain elements smaller than the `v[pivotidx]`.
func partitionEqual(v Interface, a, b, pivotidx int) int {
	v.Swap(a, pivotidx)

	L := a + 1
	R := b
	for {
		for L < R && !v.Less(a, L) {
			L++
		}
		for L < R && v.Less(a, R-1) {
			R--
		}
		if L >= R {
			break
		}
		R--
		v.Swap(L, R)
		L++
	}
	return L
}

// partialInsertionSort partially sorts a slice, returns `true` if the slice is sorted at the end.
func partialInsertionSort(v Interface, a, b int) bool {
	const (
		maxSteps         = 5  // maximum number of adjacent out-of-order pairs that will get shifted
		shortestShifting = 50 // don't shift any elements on short arrays, that has a performance cost.
	)
	i := a + 1
	for j := 0; j < maxSteps; j++ {
		for i < b && !v.Less(i, i-1) {
			i++
		}

		if i == b {
			return true
		}

		if b-a < shortestShifting {
			return false
		}

		v.Swap(i-1, i)

		// Shift the smaller one to the left.
		shiftTail(v, a, i)
		// Shift the greater one to the right.
		shiftHead(v, i, b)
	}

	return false
}

func shiftTail(v Interface, a, b int) {
	l := b - a
	if l >= 2 {
		for i := l - 1; i >= 1; i-- {
			if !v.Less(a+i, a+i-1) {
				break
			}
			v.Swap(a+i, a+i-1)
		}
	}
}

func shiftHead(v Interface, a, b int) {
	l := b - a
	if l >= 2 {
		for i := 1; i < l; i++ {
			if !v.Less(a+i, a+i-1) {
				break
			}
			v.Swap(a+i, a+i-1)
		}
	}
}

func nextPowerOfTwo(length int) uint {
	shift := uint(strconv.IntSize - bits.LeadingZeros(uint(length)))
	return uint(1 << shift)
}

// breakPatterns scatters some elements around in an attempt to break some patterns
// that might cause imbalanced partitions in quicksort.
func breakPatterns(v Interface, a, b int) {
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
			v.Swap(pos-1+i, a+other)
		}
	}
}

// choosePivot chooses a pivot in `v`.
//
// `v` might be reordered in this function.
//
// [0,8): choose a static pivot.
// [8,ShortestNinther): use the simple median-of-three method.
// [ShortestNinther,∞): use the Tukey’s ninther method.
func choosePivot(v Interface, x, y int) (pivotidx int, likelySorted bool) {
	const (
		shortestNinther = 50
		maxSwaps        = 4 * 3
	)

	l := y - x

	var (
		swaps int
		a     = x + l/4*1
		b     = x + l/4*2
		c     = x + l/4*3
	)

	if l >= 8 {
		if l >= shortestNinther {
			// Tukey’s ninther method.
			// Find medians in the neighborhoods of `a`, `b`, `c`.
			sortAdjacent(v, &a, &swaps)
			sortAdjacent(v, &b, &swaps)
			sortAdjacent(v, &c, &swaps)
		}
		// Find the median among `a`, `b`, `c`.
		sort3(v, &a, &b, &c, &swaps)
	}

	if swaps < maxSwaps {
		return b, (swaps == 0)
	} else {
		// The maximum number of swaps was performed.
		// Reversing will probably help.
		reverseRange(v, x, y)
		return 2*x + (l - 1 - b), true
	}
}

// sort2 swaps `a, b` so that `v[a] <= v[b]`.
func sort2(v Interface, a, b, swaps *int) {
	if v.Less(*b, *a) {
		*a, *b = *b, *a
		*swaps++
	}
}

// sort3 swaps `a, b, c` so that `v[a] <= v[b] <= v[c]`.
func sort3(v Interface, a, b, c, swaps *int) {
	sort2(v, a, b, swaps)
	sort2(v, b, c, swaps)
	sort2(v, a, b, swaps)
}

// sortAdjacent finds the median of `v[a - 1], v[a], v[a + 1]` and stores the index into `a`.
func sortAdjacent(v Interface, a, swaps *int) {
	t1 := *a - 1
	t2 := *a + 1
	sort3(v, &t1, a, &t2, swaps)
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
