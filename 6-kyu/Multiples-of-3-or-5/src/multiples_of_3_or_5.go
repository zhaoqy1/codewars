package src

func Multiple3And5(number int) int {
	num := 0
	for i := 0; i < number; i++ {
		if i < 0 {
			num = num
		}
		if i%3 == 0 || i%5 == 0 {
			num += i
		}
	}
	return num
}
