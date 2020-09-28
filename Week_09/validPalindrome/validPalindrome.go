package validPalindrome
func validPalindrome(s string) bool {
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{
			return isPalindrome(s[left:right])||isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string)bool{
	left :=0
	right :=len(s)-1
	for left<right{
		if s[left]!=s[right]{
			return false
		}
		left++
		right--
	}
	return true
}
