package pro796

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(s+s, goal)
}
