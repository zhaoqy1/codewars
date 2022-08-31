package src

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var str string
	for _, v := range numbers {
		str = str + strconv.Itoa(int(v))
	}
	s := fmt.Sprintf("(%v) %v-%v", str[0:3], str[3:6], str[6:10])
	return s
}
