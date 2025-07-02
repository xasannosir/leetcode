package solution

func divideArray(nums []int) bool {
    counter := make(map[int]int)

    for _, num := range nums {
        counter[num]++
    }

    for _, v := range counter {
        if v % 2 != 0 {
            return false
        }
    }

    return true
}