package pro537

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	r1, imag1 := parseComplexNumber(num1)
	r2, imag2 := parseComplexNumber(num2)

	return fmt.Sprintf("%d+%di", r1*r2-imag1*imag2, r1*imag2+r2*imag1)
}

func parseComplexNumber(num1 string) (real, imaginary int) {
	idx := strings.Index(num1, "+")

	real, _ = strconv.Atoi(num1[:idx])
	imaginary, _ = strconv.Atoi(num1[idx+1 : len(num1)-1])

	return
}
