package pro717

func isOneBitCharacter(bits []int) bool {
	last := -1
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
		} else {
			last = i
		}
	}

	return last == len(bits)-1
}
