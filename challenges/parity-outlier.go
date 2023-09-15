package challenges

// Challenge 2: Parity Outlier
func (s SpeSkillTest) FindParityOutlier(numbers []int) int {
	evenCount, oddCount := 0, 0
	var even, odd int

	// Count the number of even and odd integers in the array
	for _, num := range numbers {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}
	}

	// Determine and return the outlier based on parity
	if evenCount == 1 {
		return even
	} else if oddCount == 1 {
		return odd
	} else {
		return 0 // Return 0 if no outlier is found (shouldn't happen based on the challenge)
	}
}
