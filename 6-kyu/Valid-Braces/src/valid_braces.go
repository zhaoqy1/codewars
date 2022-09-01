package src

func ValidBraces(str string) bool {
	var stack []rune
	for _, r := range str {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			//转换成rune类型时，{是123.}125，[是91，]是93，(是40，)是41，使用-1或-2就可以确定匹配到对应的闭合括号
			if r-1 == top || r-2 == top {
				//如果可以闭合，则删掉已闭合括号
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, r)
	}
	//如果全部都为闭合括号，len(stack)为0
	return len(stack) == 0
}
