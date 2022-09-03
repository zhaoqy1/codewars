package src

func Cakes(recipe, available map[string]int) int {
	num := 0
	for k, v := range recipe {
		availableNum, ok := available[k]
		if !ok {
			return 0
		}
		if num == 0 {
			num = availableNum / v
		}
		if num > availableNum/v {
			num = availableNum / v
		}
	}
	return num
}
