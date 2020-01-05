package password

// ValidElfPassword returns true if the password meets the following criterea:
// - It is a six-digit number.
// - The value is within the range given in your puzzle input.
// - Two adjacent digits are the same (like 22 in 122345).
// - Going from left to right, the digits never decrease; they only ever
// 	 increase or stay the same (like 111123 or 135679).
func ValidElfPassword(password string) bool {
	if len(password) != 6 {
		return false
	}

	ch := rune(password[0])
	hasDouble := false
	for _, next := range password[1:] {
		if next < ch {
			return false
		}

		if next == ch {
			hasDouble = true
		}

		ch = next
	}

	return hasDouble
}

// ValidElfPassword2 returns true if the password meets the following
// criteria:
// - It is a six-digit number.
// - The value is within the range given in your puzzle input.
// - Two adjacent digits are the same (like 22 in 122345).
// - Going from left to right, the digits never decrease; they only ever
// 	 increase or stay the same (like 111123 or 135679).
// - two adjacent matching digits may not be part of a larger group of
//   matching digits (meaning, there must be at least one EXACT pair)
func ValidElfPassword2(password string) bool {
	if len(password) != 6 {
		return false
	}

	ch := rune(password[0])
	for _, next := range password[1:] {
		if ch > next {
			return false
		}

		ch = next
	}

	counts := make(map[rune]int)
	for _, ch := range password {
		if _, ok := counts[ch]; !ok {
			counts[ch] = 1
			continue
		}

		counts[ch]++
	}

	for _, v := range counts {
		if v == 2 {
			return true
		}
	}

	return false
}
