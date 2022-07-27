package eq

func EqualSliceInt(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, val := range nums1 {
		if nums2[i] != val {
			return false
		}
	}
	return true
}

func EqualSliceIntElement(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	m := make(map[int]struct{}, len(nums1))
	for _, val := range nums1 {
		m[val] = struct{}{}
	}

	for _, val := range nums2 {
		if _, ok := m[val]; !ok {
			return false
		}
	}
	return true
}

func EqualSliceString(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	nums1Map := make(map[string]struct{}, len(s1))
	for _, v := range s1 {
		nums1Map[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := nums1Map[v]; !ok {
			return false
		}
	}
	return true
}
