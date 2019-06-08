package util

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
