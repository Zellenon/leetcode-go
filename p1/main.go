//Solution passes all tests with 0ms runtime (meets or beats 100% of competitors) and 5.95MB memory
//(meets or beats 40.05% of competitors)

package p1

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int);
	for i, x := range nums {
		var remainder = target - x;
		if j, ok := hmap[remainder]; ok {
			return []int{i,j};
		} else {
			hmap[x] = i;
		}
	}
	return []int{-1,-1};
}
