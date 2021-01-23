package main

import "fmt"

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	help := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			cnt += j - (mid + 1)
			help = append(help, nums[i])
			i++
		} else {
			help = append(help, nums[j])
			j++
		}
	}
	for i <= mid {
		help = append(help, nums[i])
		cnt += end - (mid + 1) + 1
		i++
	}
	for j <= end {
		help = append(help, nums[j])
		j++
	}
	for i := start; i <= end; i++ {
		nums[i] = help[i-start]
	}
	return cnt
}

func main() {
	// arr := []int{4, 5, 6, 7}
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
	// fmt.Println(reversePairs(arr))
}
