package reloaded

func Apostrophe(s []string) []string {
	apostropheCount := 0
	for i, current := range s {
		// Check if the current element is an apostrophe and if the count is 0
		if current == "'" && apostropheCount == 0 {
			// If so, concatenate the apostrophe with the next element in the slice
			s[i+1] = current + s[i+1]
			// Remove the apostrophe element from the slice
			s = append(s[:i], s[i+1:]...)
			// Increment the count to indicate that an apostrophe has been handled
			apostropheCount++
		}
	}

	for i, current := range s {
		// Check if the current element is an apostrophe
		if current == "'" {
			// If so, concatenate the apostrophe with the previous element in the slice
			s[i-1] += current
			// Reset the count to 0 for the next iteration
			apostropheCount = 0
			// Remove the apostrophe element from the slice
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
