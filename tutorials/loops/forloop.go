package loops

// Insert all the elements up to the limit in a list.
func PrintCounter(limit int) []int {
	var x []int

	for i := 0; i < limit; i++ {
		x = append(x, i)
	}
	return x
}
