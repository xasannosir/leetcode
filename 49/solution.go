package solution

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	groups := make(map[string][]int)

	for idx, str := range strs {
		items := strings.Split(str, "")
		sort.Strings(items)

		item := strings.Join(items, "")

		groups[item] = append(groups[item], idx)
	}

	for _, idxs := range groups {
		item := make([]string, 0)
		for _, idx := range idxs {
			item = append(item, strs[idx])
		}
		res = append(res, item)
	}

	return res
}
