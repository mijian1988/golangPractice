package main

import (
	"fmt"
)

/* 
 *二分查找函数,假设有序数组的顺序是从小到大（很关键，决定左右方向）
 */
func binarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int)  {
	//判断 leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}

func main() {
	arr := []int{21,32,12,33,34,34,87,24}
	var n = len(arr)

  // 冒泡排序
	for i := 0; i <= n-1; i++ {
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}

	binarySearch(&arr, 0, len(arr) - 1, 33)
}
