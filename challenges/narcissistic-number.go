package challenges

import "fmt"

// Challenge 1: Narcissistic Number
func (s SpeSkillTest) IsNarcissistic(num int) bool {
	// Convert the number to a string to count its digits
	numStr := fmt.Sprintf("%d", num)
	numDigits := len(numStr)

	// Calculate the sum of each digit raised to the power of the number of digits
	sum := 0
	for _, digit := range numStr {
		digitValue := int(digit - '0')
		sum += intPow(digitValue, numDigits)
	}

	// Check if the sum equals the original number
	return sum == num
}

// Helper function to calculate the power of an integer
func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
