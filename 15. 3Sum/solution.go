func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)

    for i := range nums {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        j := i + 1
        k := len(nums) - 1

        for j < k {
            sum := nums[i] + nums[j] + nums[k]

            switch {
                case sum < 0: j += 1
                case sum > 0: k -= 1
                case sum == 0: 
                    res = append(res, []int{nums[i], nums[j], nums[k]})
                    j += 1
                    for j < k && nums[j] == nums[j -1] {
                        j += 1
                    }
            }

        }
    }
    return res
}
