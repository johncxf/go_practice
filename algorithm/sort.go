// 排序算法
package main

import "fmt"

// 冒泡排序
func bubbleSort(arr []int) []int {
	count := len(arr)
	if 1 >= count {
		return arr
	}
	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			// 从小到大
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// 从大到小
			// if arr[j] < arr[j+1] {
			// 	arr[j], arr[j+1] = arr[j+1], arr[j]
			// }
		}
	}
	return arr
}

// 快速排序
func quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	temp := arr[left]
	i := left
	j := right
	for i != j {
		// 哨兵 j 向左移动，查找小于基准数 temp 时停下
		for arr[j] >= temp && i < j {
			j--
		}
		// 哨兵 i 向右移动，查找大于基准数 temp 时停下
		for arr[i] <= temp && i < j {
			i++
		}
		// 交换 i、j 位置
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 将基准数归位
	arr[left], arr[i] = arr[i], arr[left]
	quickSort(arr, left, i-1)
	quickSort(arr, j+1, right)
}

// 选择排序
func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 递归合并数组（从小到大）
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// 归并排序
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	// 递归的进行数组拆分
	// 左区间：[left, mid]，右区间：[mid, right]
	middle := length / 2
	left := mergeSort(arr[0:middle])
	right := mergeSort(arr[middle:])

	// 递归进行数组合并
	return merge(left, right)
}

func main() {
	arr := []int{4, 5, 6, 7, 8, 3, 2, 1}
	arr = mergeSort(arr)
	// arr := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	// quickSort(arr, 0, len(arr)-1)
	// arr = selectSort(arr)
	fmt.Println(arr)
}
