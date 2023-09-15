package main

import (
	"fmt"
	"github.com/balqisgautama/test-spe/challenges"
)

func main() {
	// Example usages
	spe := challenges.SpeSkillTest{}

	// Challenge 1
	fmt.Println("Challenge-1")
	fmt.Println(spe.IsNarcissistic(153))
	fmt.Println(spe.IsNarcissistic(111))
}
