package E_125_Valid_Palindrome

func validChar(c byte) (bool, byte) {
	if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		if c >= 65 && c <= 90 {
			return true, c + 32
		}
		return true, c
	}
	return false, 0
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		ok, leftChar := validChar(s[left])
		if !ok {
			left += 1
			continue
		}

		ok, rightChar := validChar(s[right])
		if !ok {
			right -= 1
			continue
		}

		if leftChar != rightChar {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
