package mymath

import (
	"strconv"
)

const version = "1.0.0.0"

func ToFloat32(in int) float32 {
	return float32(in)
}

func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic("Can't convert string to number!")
		//return -9999999
	} else {
		return num
	}
}

func ToBool(str string) bool {
	bool1, err := strconv.ParseBool(str)
	if err != nil {
		//panic("Can't convert string to boolean!")
		return false
	} else {
		return bool1
	}
}
