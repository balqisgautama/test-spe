package challenges

// Challenge 4: The Blue Ocean Reverse
func (s SpeSkillTest) BlueOcean(firstList, secondList []int) []int {
	// Create a map to store the values from the second list
	secondListValues := make(map[int]bool)
	for _, val := range secondList {
		secondListValues[val] = true
	}

	// Initialize a result slice
	result := []int{}

	// Iterate through the first list and add values not present in the second list to the result
	for _, val := range firstList {
		if !secondListValues[val] {
			result = append(result, val)
		}
	}

	return result
}
