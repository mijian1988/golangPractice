package main

import (
	"fmt"
	"strconv"
)

func convertToBin(num int) string{
	ret := ""
	for ;num > 0;num /=2{
		lsb := num % 2
		ret = strconv.Itoa(lsb)+ret
	}
	return ret
}

func main() {
	fmt.Println(convertToBin(5))
}
