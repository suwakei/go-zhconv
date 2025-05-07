package zhconv



// indexRune finds the index of a rune within a string.
// Returns -1 if the rune is not found.
// Consider moving to a shared utility package if used in multiple places.
func indexRune(s []rune, r rune) int {
	// Ranging over a string iterates through runes directly.
	for i, char := range s {
		if char == r {
			return i
		}
	}
	return -1
}
