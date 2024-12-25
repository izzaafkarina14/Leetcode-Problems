package easy

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	input := x
	reversedigit := 0

	for x > 0 {
		lastdigit := x % 10
		reversedigit = reversedigit*10 + lastdigit

		x = x / 10
	}
	return input == reversedigit
}