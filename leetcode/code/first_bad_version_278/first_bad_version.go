package first_bad_version_278

func firstBadVersion(version int) int {
	low := 1
	height := version

	for low < height {
		mid := ((height - low) >> 1) + low
		if isBadVersion(mid) {
			height = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func isBadVersion(version int) bool {
	return false
}
