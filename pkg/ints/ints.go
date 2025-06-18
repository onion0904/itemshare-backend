package ints

import "strconv"

func Digit(num int32) int32 {
	str := strconv.Itoa(int(num))
	return int32(len(str))
}
