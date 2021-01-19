package main

import "fmt"

func singleNumbers(nums []int) []int {
	eor := 0
	for _, v := range nums {
		eor ^= v
	}
	rightOne := eor & (^eor + 1)
	onlyOne := 0
	for _, v := range nums {
		if v&rightOne == 0 {
			onlyOne ^= v
		}
	}
	return []int{onlyOne, onlyOne ^ eor}
}

func main() {
	arr := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(arr))
}
