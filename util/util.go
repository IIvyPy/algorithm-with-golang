package util

import "fmt"

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
		fmt.Println("iji", i,j, pvalue)
		if list[j] < pvalue{
			fmt.Println("iiiii", i,j, pvalue)
			if list[i] > pvalue{
				fmt.Println("iiiiji", i,j,pvalue)
				list[i], list[j] = list[j], list[i]
			}
			i++
		}
		j--
	}
	list[i], list[begin] = list[begin], list[i]
	return i
}

func QuickSort(testList []int, begin int, end int){
	for begin < end{
		mid := partision(testList, begin, end)
		QuickSort(testList, begin, mid)
		QuickSort(testList, mid+1, end)
	}
}

