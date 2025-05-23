//Solution passes all tests with 0ms runtime (meets or beats 100% of competitors) and 6.5MB memory
//(meets or beats 98.63% of competitors)
package p4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_len :=len(nums1) + len(nums2);
	bridge_mids := total_len % 2 == 0;
	next_num := 0;

	nextNum := func() int {
		if len(nums1) > 0 {
			if len(nums2) > 0 {
				if nums1[0] < nums2[0] {
					next_num, nums1 = nums1[0], nums1[1:];
				} else {
					next_num, nums2 = nums2[0], nums2[1:];
				}
			} else {
					next_num, nums1 = nums1[0], nums1[1:];
			}
		} else {
					next_num, nums2 = nums2[0], nums2[1:];
		}
		return next_num
	};

	i := 0;
	if bridge_mids {
		i++;
	}
	for i < total_len/2 {
		nextNum();
		i++;
	}

	if bridge_mids {
		a := float64(nextNum());
		b := float64(nextNum());
		return (a+b)/2;

	} else {
			return float64(nextNum());
	}
}
