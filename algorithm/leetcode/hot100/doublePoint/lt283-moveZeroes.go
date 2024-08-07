// [L283-简单] 移动零
package main

import "fmt"

// 双指针
func moveZeros(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	fmt.Println(nums)
}
