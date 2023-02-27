package lib

// deduplicate operation
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

// get pairs operation
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
