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

	// Challenge 2
	fmt.Println("Challenge-2")
	fmt.Println(spe.FindParityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(spe.FindParityOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(spe.FindParityOutlier([]int{11, 13, 15, 19, 9, 13, -21}))

	// Challenge 3
	fmt.Println("Challenge-3")
	fmt.Println(spe.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	// Challenge 4
	fmt.Println("Challenge-4")
	fmt.Println(spe.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	fmt.Println(spe.BlueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))
}
