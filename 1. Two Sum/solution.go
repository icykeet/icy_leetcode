func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i := range nums {
        complement := target - nums[i]
        if j, ok := m[complement]; ok {
            return []int{j,i}
        }
        m[nums[i]] = i
    }

    return []int{}
}
