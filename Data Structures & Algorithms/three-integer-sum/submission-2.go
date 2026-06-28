import (
	"slices"
)

func threeSum(nums []int) [][]int {
	i, k := 0, len(nums)-1
	res := make([][]int, 0)

	slices.Sort(nums)

	for i < k {
		if i > 0 && nums[i] == nums[i-1]{
			i++
			continue
		}

		j := i+1
		for j < k {
			sum := nums[i]+nums[j]+nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
		i++
		k = len(nums)-1
	}

	return res
}
