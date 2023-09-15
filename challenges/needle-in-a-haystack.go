package challenges

// Challenge 3: Needle in a Haystack
func (s SpeSkillTest) FindNeedle(haystack []string, needle string) int {
	for i, str := range haystack {
		if str == needle {
			return i
		}
	}
	return -1 // Return -1 if the needle is not found in the haystack
}
