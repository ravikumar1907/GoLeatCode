package main

import (
	"sort"
)

// 1. Two Sum
func twoSum(nums []int, target int) []int {
	preMap := make(map[int]int)
	result := []int{-1, -1}

	for i, v := range nums {
		preMap[v] = i
	}
	for i, v := range nums {
		tmp := target - v
		if j, ok := preMap[tmp]; ok && i != j {
			result[0] = i
			result[1] = j
			break
		}
	}
	return result
}

// two sum given sorted array
// 167. Two Sum II - Input Array Is Sorted
func twosumSorted(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

// 15. 3Sum
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

// 18. 4Sum
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)
	for h := 0; h < n-3; h++ {

		if h > 0 && nums[h] == nums[h-1] {
			continue
		}

		for i := h + 1; i < n-2; i++ {
			if i > h+1 && nums[i] == nums[i-1] {
				continue
			}

			j := i + 1
			k := n - 1
			for j < k {
				sum := nums[h] + nums[i] + nums[j] + nums[k]
				if sum == target {
					result = append(result, []int{nums[h], nums[i], nums[j], nums[k]})
					k--
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if sum > 0 {
					k--
				} else {
					j++
				}
			}
		}
	}
	return result
}

func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a[i], a[j] = a[j], a[i]
		}
	}
}

// 560. Subarray Sum Equals K
func subarraySum(nums []int, k int) int {
	n := len(nums)
	hMap := make(map[int]int)
	hMap[0] = 1
	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if hMap[sum-k] != 0 {
			// if v, ok := hMap[sum-k]; ok {
			ans += hMap[sum-k]
		}
		hMap[sum]++
	}
	return ans
}

// 1679. Max Number of K-Sum Pairs
func maxOperations(nums []int, k int) int {
	hMap := make(map[int]int)
	cnt := 0
	for _, v := range nums {
		if hMap[k-v] > 0 {
			cnt++
			hMap[k-v]--
		} else {
			hMap[v]++
		}
	}
	return cnt
}

/*
func main() {
	nums := []int{11, 7, 2, 15}
	fmt.Printf("%v\n", twoSum(nums, 9))
	nums = []int{3, 2, 4}
	fmt.Printf("%v\n", twoSum(nums, 6))
	nums = []int{3, 3}
	fmt.Printf("%v\n", twoSum(nums, 6))
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v\n", threeSum(nums))
	nums = []int{2, 2, 2, 2, 2}
	fmt.Printf("%v\n", fourSum(nums, 8))
	//	nums = []int{3, 1, 3, 4, 3}
	//	fmt.Println(maxOperations(nums, 6))
}
*/
