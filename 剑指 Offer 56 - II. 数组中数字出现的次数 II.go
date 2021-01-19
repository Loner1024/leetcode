package main

import "fmt"

func singleNumber(nums []int) int {
	singleMap := make(map[int]int)
	for _, v := range nums {
		singleMap[v] += 1
	}
	for k, v := range singleMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	arr := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber(arr))
}
