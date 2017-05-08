package main

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	str := ""
	for _, v := range s {
		if (v >= 'a' && v <= 'z') ||
			(v >= '0' && v <= '9') {
			str = str + string(v)
		}
		if v >= 'A' && v <= 'Z' {
			str = str + string(v-'A'+'a')
		}
	}
	if str == "" {
		return true
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
