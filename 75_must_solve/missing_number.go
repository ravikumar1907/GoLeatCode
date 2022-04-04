package main

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

/*func main() {
	a := []int{0, 1}
	fmt.Println(missingNumber(a))
	a = []int{2, 2, 2, 2, 2}
	fmt.Println(findDuplicate(a))
}*/
