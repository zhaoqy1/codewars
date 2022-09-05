package src

func Zeros(n int) int {
	result := 0
	//查看阶乘表格总结出规律n每增加5，尾数0的个数增加1，每增加25再额外加1
	for i := 5; i <= n; i *= 5 {
		for j := i; j <= n; j += i {
			result++
		}

	}
	return result
}
