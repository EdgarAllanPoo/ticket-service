package utility

import (
	"fmt"
	"strconv"
)

func StrToUint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 0)

	if err != nil {
		fmt.Println("Error :", err)
		return 0
	}

	uintValue := uint(num)

	return uintValue
}
