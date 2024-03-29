func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
    number := 0
    start := x
    for x > 0 {
        number += (x % 10)
        x /= 10
        number *= 10
    }

    return start == number / 10
}