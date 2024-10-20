func sortedSquares(nums []int) []int {
    size := len(nums)
    res := make([]int, size)
    left := 0
    right := size - 1

    for i := (size - 1); i >= 0; i-- {
        if checkNum(nums[left]) > checkNum(nums[right]) {
            res[i] = getSquare(nums[left])
            left++
        } else if (checkNum(nums[right]) > checkNum(nums[left])) {
            res[i] = getSquare(nums[right])
            right--
        } else {
            if i == 0 {
                res[i] = getSquare(nums[right])
                break
            }

            res[i], res[i -1] = getSquare(nums[right]), getSquare(nums[left])
            i--
            left++
            right--
        }
    }

    return res
}

func checkNum(number int) float64 {
    return math.Abs(float64(number))
}

func getSquare(number int) int {
    return number * number
}
