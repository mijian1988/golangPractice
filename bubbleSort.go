package main

import (
	"fmt"
)

/*
 *  冒泡排序
 */
func main() {
	arr := [...]int{21,32,12,33,34,34,87,24}
	var n = len(arr)
	fmt.Println("--------排序前--------\n",arr)

	for i := 0; i <= n-1; i++ {
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}

	fmt.Println("--------排序后--------\n",arr)
}
