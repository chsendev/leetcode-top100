package leetcode_top100

import (
	"fmt"
	"testing"
)

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[i] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func sort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		if num[i] == 0 {
			for j := 0; j < len(num)-i-1; j++ {
				if num[j] < num[j+1] {
					num[j], num[j+1] = num[j+1], num[j]
				}
			}
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	//moveZeroes(nums)
	sort(nums)
	fmt.Println(nums)
}
