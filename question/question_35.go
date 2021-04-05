package question

func searchInsert(nums []int, target int) int {
	if len(nums) == 0{
		return 0
	}

	left := 0
	right := len(nums) - 1
	var middle int
	for left <= right{
		middle = (left + right)/2
		if nums[middle] > target{
			right = middle - 1
		}else if nums[middle] < target{
			left = middle + 1
		}else{
			return middle
		}
	}
	return left
}
