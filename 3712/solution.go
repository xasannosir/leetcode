package solution

func sumDivisibleByK(nums []int, k int) int {
    total := 0
    freq := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }

    for num, occur := range freq {
        if occur % k == 0 {
            total += num * occur
        }
    }

    return total
}
