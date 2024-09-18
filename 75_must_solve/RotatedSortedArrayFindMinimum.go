package main

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 153. Find Minimum in Rotated Sorted Array
// o(logn)
func findMin3(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]
	for l <= r {
		if nums[l] <= nums[r] {
			res = Min(res, nums[l])
			break
		}
		m := (l + r) / 2
		res = Min(res, nums[m])
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			// Minimum is in the right part
			left = mid + 1
		} else {
			// Minimum is in the left part or at mid
			right = mid
		}
	}
	return nums[left]
}

// O(n)
func findMin1(nums []int) int {
	n := len(nums)
	res := nums[0]
	i := 1
	for i < n {
		if nums[i-1] < nums[i] {
			i++
		} else {
			res = Min(res, nums[i])
			break
		}
	}
	return res
}

//154. Find Minimum in Rotated Sorted Array II

func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	if r < 0 {
		return 0
	}
	res := nums[0]
	for l <= r {
		for ; l < r && nums[l] == nums[r]; r-- {
		}
		for ; l < r && nums[l] == nums[l+1]; l++ {
		}
		for ; l < r && nums[r-1] == nums[r]; r-- {
		}
		if nums[l] <= nums[r] {
			res = Min(res, nums[l])
			break
		}
		m := (l + r) / 2
		res = Min(res, nums[m])
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}

// 154. Find Minimum in Rotated Sorted Array II
/*func main() {
	nums := []int{6, 7, 1, 2, 3, 4, 5}
	fmt.Println(findMin(nums))
	fmt.Println(findMin2([]int{}))
}*/
