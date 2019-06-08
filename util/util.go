package util

import "fmt"

/*
经典的冒泡算法
input: intList []int
output:
 */
func bubble(intList []int){
	length := len(intList)

	for i := 0; i < length; i++{
		for j := i; j < length; j++{
			if intList[i] < intList[j]{
				intList[i],intList[j] = intList[j],intList[i]
			}
		}
	}
}

/*
若一个数组中有一个数字出现的次数超过半数，找之
input: intList []int
output:
    - 找到：
		超过半数的数字， true
    - 未找到：
        0, false
 */
func findMostNum(intList []int)(int, bool){
	//count := 0
	length := len(intList)
	comap := map[int]int{}
	for i := 0; i < len(intList); i++{
		_, isExact := comap[intList[i]]
		if isExact{
			//有这个数
			comap[intList[i]]++
			if comap[intList[i]] >= length/2{
				//break
				return intList[i], true
			}
		}else{
			//没有这个数
			comap[intList[i]] = 1
		}
	}
	fmt.Println("ddd", comap)
	return 0, false
}

/*
给一个数组和目标值，找到两个数之和等于该目标值
input: intList []int， target int
output:
    - 找到：
		两数的索引
    - 未找到：
        -1, -1
 */
func FindTwoSum(input []int, target int)(int, int){
	leftNum := 0
	rightNum := len(input) - 1

	if target <= input[leftNum]{
		return -1, -1
	}

	for leftNum != rightNum{
		if input[rightNum] > target - input[leftNum]{
			rightNum--
		} else if input[rightNum] < target - input[leftNum]{
			leftNum++
		}else{
			return leftNum, rightNum
		}
	}

	return -1, -1
}

func partision(list []int, begin int, end int) int{
	pvalue := list[begin]
	i := begin
	j := end

	for i <= j{
		fmt.Println("init", i,j, pvalue)
		if list[j] < pvalue{
			fmt.Println("find right value", i,j, pvalue)
			if list[i] > pvalue{
				fmt.Println("find left value", i,j,pvalue)
				list[i], list[j] = list[j], list[i]
			}
			i++
		}
		j--
	}
	list[i], list[begin] = list[begin], list[i]
	return i
}

/*
快排
input: intList []int， begin int, end int
output:
 */
func QuickSort(testList []int, begin int, end int){
	for begin < end{
		mid := partision(testList, begin, end)
		QuickSort(testList, begin, mid)
		QuickSort(testList, mid+1, end)
	}
}

/*
寻找重复的数字, 若找到则返回true，反之为false
input: arr []int
output: bool
 */
func findRepeatNum(arr []int)bool{
	if len(arr) == 0{
		return false
	}

	for i := 0; i < len(arr); i++{
		for arr[i] != i{
			if arr[i] == arr[arr[i]]{
				return true
			}
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
			i = arr[i]
		}
	}

	return false
}

/*
二分查找，有序数组，用的是非递归
input: arr []int, keyNum int
output: int, 找到则对应的索引，反之为-1
 */
func BinarySearch(arr []int, keyNum int) int{
	length := len(arr)
	if length == 0{
		return -1
	}
	start := 0
	end := length - 1
	middle := (start + end)/2

	for start <= end{
		if keyNum > arr[middle]{
			start = middle + 1
			middle = (start + end)/2
		}else if keyNum < arr[middle]{
			end = middle - 1
			middle = (start + end)/2
		}else{
			return middle
		}
	}
	return -1
}

/*
选择排序
input: arr []int
output:
 */
func SelectionSort(arr []int){
	outer := len(arr) - 1
	inner := len(arr)

	for i := 0; i < outer; i++{
		k := i
		//找出后面的最小的那个值得索引
		for j := i; j < inner; j++{
			if arr[k] > arr[j]{
				k = j
			}
		}
		//如果没找到的话则继续循环，如果找到了的话则换就成
		if k != i{
			arr[k], arr[i] = arr[i], arr[k]
		}
	}
}