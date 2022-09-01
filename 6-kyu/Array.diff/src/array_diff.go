package src

func ArrayDiff(a, b []int) []int {
	// your code here
	c := make(map[int]bool)
	for _, v := range b {
		c[v] = true
	}
	result := make([]int, 0)
	for _, v := range a {
		if !c[v] {
			result = append(result, v)
		}
	}
	return result
}
