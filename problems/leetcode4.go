package problems

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k1 := (m + n + 1) / 2
	k2 := (m + n + 2) / 2
	return (findKthSmallest(nums1, nums2, 0, 0, m-1, n-1, k1) + findKthSmallest(nums1, nums2, 0, 0, m-1, n-1, k2)) * 0.5
}

func findKthSmallest(nums1, nums2 []int, s1, s2, e1, e2, k int) float64 {
	l1 := e1 - s1 + 1
	l2 := e2 - s2 + 1
	if l1 > l2 {
		return findKthSmallest(nums2, nums1, s2, s1, e2, e1, k)
	}
	if l1 == 0 {
		return float64(nums2[s2+k/2-1])
	}
	if k == 1 {
		return float64(minInt(nums1[s1], nums2[s2]))
	}

	p1 := s1 + minInt(k/2, l1) - 1
	p2 := s2 + minInt(k/2, l2) - 1
	if nums1[p1] < nums2[p2] {
		return findKthSmallest(nums1, nums2, p1+1, p2, e1, e2, k-(p1-s1+1))
	} else {
		return findKthSmallest(nums1, nums2, p1, p2+1, e1, e2, k-(p2-s2+1))

	}

}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
