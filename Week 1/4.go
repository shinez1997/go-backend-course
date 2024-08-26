//4. Viết chương trình nhập giải bài toán twosum : https://leetcode.com/problems/two-sum/

func twoSum(a []int, target int) []int {
	idx := make(map[int]int)
	for i, x := range a {
		if id, ok := idx[target-x]; ok == true {
			return []int{id, i}
		}
		idx[x] = i
	}
	return nil
}