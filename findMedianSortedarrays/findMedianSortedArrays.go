package main

import "math"

func main() {
	a := []int{1, 3}
	b := []int{2}
	println(a, b, findMedianSortedArrays(a, b))
	a = []int{2}
	b = []int{1, 3, 4, 5}
	println(a, b, findMedianSortedArrays(a, b))
}

type SortedArrayIter struct {
	Arr      []int
	iter_idx int
}

func SortedArrayIterIsDone(nums *SortedArrayIter) bool {
	return nums.iter_idx >= len(nums.Arr)
}

func getNextValue(numsA *SortedArrayIter, numsB *SortedArrayIter) int {
	var val int
	if SortedArrayIterIsDone(numsA) {
		val = numsB.Arr[numsB.iter_idx]
		numsB.iter_idx += 1
	} else if SortedArrayIterIsDone(numsB) {
		val = numsA.Arr[numsA.iter_idx]
		numsA.iter_idx += 1
	} else {
		if numsA.Arr[numsA.iter_idx] < numsB.Arr[numsB.iter_idx] {
			val = numsA.Arr[numsA.iter_idx]
			numsA.iter_idx += 1
		} else {
			val = numsB.Arr[numsB.iter_idx]
			numsB.iter_idx += 1
		}
	}
	return val
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1_iter := SortedArrayIter{
		Arr:      nums1,
		iter_idx: 0,
	}

	nums2_iter := SortedArrayIter{
		Arr:      nums2,
		iter_idx: 0,
	}

	idx_low := int(math.Floor(float64(len(nums1)+len(nums2)-1) / 2.0))
	idx_high := int(math.Ceil(float64(len(nums1)+len(nums2)-1) / 2.0))
	println("idx low = ", idx_low)
	println("idx high = ", idx_high)
	var value_idx_low int
	var value_idx_high int
	var val int

	for idx_low >= 0 {
		// if one array is finished, lookup both values in the other array
		if SortedArrayIterIsDone(&nums1_iter) {
			value_idx_low = nums2[nums2_iter.iter_idx+idx_low]
			value_idx_high = nums2[nums2_iter.iter_idx+idx_high]
			return float64(value_idx_high+value_idx_low) / 2.0
		} else if SortedArrayIterIsDone(&nums2_iter) {
			value_idx_low = nums1[nums1_iter.iter_idx+idx_low]
			value_idx_high = nums1[nums1_iter.iter_idx+idx_high]
			return float64(value_idx_high+value_idx_low) / 2.0
		}
		val = getNextValue(&nums1_iter, &nums2_iter)

		idx_low -= 1
		idx_high -= 1

		println("idx_low=", idx_low, ", idx_high=", idx_high, ", num1_idx=", nums1_iter.iter_idx, ", nums2_idx=", nums2_iter.iter_idx)
	}
	value_idx_low = val

	if idx_high != idx_low {
		if SortedArrayIterIsDone(&nums1_iter) {
			value_idx_high = nums2[nums2_iter.iter_idx]
		} else if SortedArrayIterIsDone(&nums2_iter) {
			value_idx_high = nums1[nums1_iter.iter_idx]
		} else {
			value_idx_high = getNextValue(&nums1_iter, &nums2_iter)
		}
		return float64(value_idx_high+value_idx_low) / 2.0
	} else {
		return float64(value_idx_low)
	}
}
