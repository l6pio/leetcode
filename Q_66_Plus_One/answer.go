package Q_66_Plus_One

// https://leetcode-cn.com/problems/plus-one/

func reverse(digits []int, digitsLen int) {
	left, right := 0, digitsLen-1
	for left < right {
		tmp := digits[left]
		digits[left] = digits[right]
		digits[right] = tmp
		left++
		right--
	}
}

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	remain := 1

	reverse(digits, digitsLen)

	for i := 0; i < digitsLen; i++ {
		tmp := digits[i] + remain
		digits[i] = tmp % 10
		remain = tmp / 10
	}

	if remain > 0 {
		digits = append(digits, 1)
		digitsLen++
	}

	reverse(digits, digitsLen)
	return digits
}
