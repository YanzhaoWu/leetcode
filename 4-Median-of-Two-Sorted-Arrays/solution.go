package main

import "fmt"

func getMedian(nums []int) float64 {
	var length int = len(nums)
	if length == 0 {
		return 0
	} else {
		return (float64(nums[(length - 1) / 2]) + float64(nums[length / 2] - nums[(length - 1)/  2]) / 2.0) // Avoid overflow
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMedianWithOneElement(num int, nums []int) float64 {
	var length int = len(nums)
	
	if length % 2 == 0 {
		if (num <= nums[length / 2] && num >= nums[length / 2 - 1]) {
			return float64(num)
		} else {
			if (num < nums[length / 2 - 1]) {
				return float64(nums[length / 2 - 1])
			} else {
				return float64(nums[length / 2])
			}
		}
	} else {
		if (num <= nums[length / 2 + 1] && num >= nums[length / 2 - 1]) {
			return float64(num + nums[length / 2]) / 2.0
		} else {
			if (num < nums[length / 2 - 1]) {
				return float64(nums[length / 2 - 1] + nums[length / 2]) / 2.0
			} else {
				return float64(nums[length / 2] + nums[length / 2 + 1]) / 2.0
			}
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var ph1, ph2 int = 0, 0
	var pe1, pe2 int = len(nums1), len(nums2)
	var offset int
	
	for ; ph1 + 1 < pe1 && ph2 + 1 < pe2; {
		offset = min((pe1 - ph1) / 2, (pe2 - ph2) / 2)
		if nums1[pe1 - offset] > nums2[pe2 - offset] {
			pe1 -= offset
		} else {
			pe2 -= offset
		}
		if nums1[ph1 + offset - 1] < nums2[ph2 + offset - 1] {
			ph1 += offset
		} else {
			ph2 += offset
		}
		//fmt.Println(ph1, pe1, ph2, pe2)
	}
	
	if pe1 == ph1 && pe2 == ph2 {
		return 0
	} else if pe1 == ph1 {
		return getMedian(nums2[ph2:pe2])
	} else if pe2 == ph2 {
		return getMedian(nums1[ph1:pe1])
	}
	
	if (pe1 == ph1 + 1 && pe2 == ph2 + 1) {
		return float64(nums1[ph1] + nums2[ph2]) / 2.0
	} else {
		if (pe1 == ph1 + 1) {
			return getMedianWithOneElement(nums1[ph1], nums2[ph2:pe2])
		} else if (pe2 == ph2 + 1) {
			return getMedianWithOneElement(nums2[ph2], nums1[ph1:pe1])
		} else {
			return 0
		}
	}
	return 0
}


func main() {
	a1 := []int {1, 3, 4, 5}
	a2 := []int {2, 6, 7, 8}
	fmt.Println(findMedianSortedArrays(a1, a2))
}