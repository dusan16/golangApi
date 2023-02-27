package lib

/*
* Deduplicate Operation
* in this fuction we achive the in-place space complexity as it can be seen that complexity is O(n),
* we have to arrays of size equeal to the input which we assume is some big number n,
* which is in total size 2*n plus some constant value c which when we take the highest degree polynomial is O(n)
 */
func Deduplicate(array []int64) (tmpArray []int64) {
	var foundMatch bool = false
	tmpArray = append(tmpArray, array[0])
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(tmpArray); j++ {
			if array[i] == tmpArray[j] {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			tmpArray = append(tmpArray, array[i])
		}
		foundMatch = false
	}
	return
}

/*
* In this function we look for multiple occurences of the numbers if number has more than one occurence we summ all those occurences
* in the map where we store key value as the number and its nuber of occurences in value part
 */
func GetPairs(array []int64) (resultMap map[int64]int64, arrayOfPairs []int64) {
	var tmpMap = make(map[int64]int64)
	resultMap = make(map[int64]int64)

	for i := 0; i < len(array); i++ {
		if _, ok := tmpMap[array[i]]; ok && tmpMap[array[i]] < 1 {
			tmpMap[array[i]] = 1
		} else {
			tmpMap[array[i]]++
		}
	}

	for i := 0; i < len(array); i++ {
		if tmpMap[array[i]] > 1 {
			resultMap[array[i]] = tmpMap[array[i]]
		}
	}

	for key, _ := range resultMap {
		arrayOfPairs = append(arrayOfPairs, key)
	}

	return
}
