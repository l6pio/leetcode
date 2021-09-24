package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	k1 := x
	k2 := 0
	for x > 0 {
		r := x % 10
		x /= 10
		k2 = k2*10 + r
	}
	return k1 == k2
}

func main() {
	ret := isPalindrome(12)
	println(ret)
}
