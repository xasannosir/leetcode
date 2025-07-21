package solution

func findWinners(matches [][]int) [][]int {
    winner := make(map[int]int)
    loser := make(map[int]int)
    notLost := make([]int, 0)
    oneLost := make([]int, 0)

    merge := func(first, second []int) []int {
        var sorted []int
        i, j := 0, 0

        for i < len(first) && j < len(second) {
            if first[i] < second[j] {
                sorted = append(sorted, first[i])
                i++
            } else {
                sorted = append(sorted, second[j])
                j++
            }
        }

        sorted = append(sorted, first[i:]...)
        sorted = append(sorted, second[j:]...)

        return sorted
    }

    var mergeSort func(nums []int) []int
    mergeSort = func(nums []int) []int {
        if len(nums) < 2 {
            return nums
        }

        mid := len(nums) / 2
        first := mergeSort(nums[:mid])
        second := mergeSort(nums[mid:])

        return merge(first, second)
    }

    for _, match := range matches {
        if _, ok := loser[match[0]]; !ok {
            winner[match[0]]++
        }
        delete(winner, match[1])
        loser[match[1]]++
    }

    for k := range winner {
        notLost = append(notLost, k)
    }
    
    for k, v := range loser {
        if v == 1 {
            oneLost = append(oneLost, k)
        }
    }

    return [][]int{mergeSort(notLost), mergeSort(oneLost)}
}