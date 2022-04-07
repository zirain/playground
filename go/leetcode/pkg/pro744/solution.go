package pro744

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}

	for _, b := range letters {
		if b > target {
			return b
		}
	}

	return letters[0]
}
