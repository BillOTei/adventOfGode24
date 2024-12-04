package utils

import (
	"fmt"
	"strconv"
)

func ParseUint64(s string) uint64 {
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	return u64
}

func RemoveIndex(s []uint64, index int) []uint64 {
	ret := make([]uint64, 0)
	ret = append(ret, s[:index]...)

	return append(ret, s[index+1:]...)
}
