package solutions

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]uint16
	for i := 0; i < len(s); i++ {
		count[s[i]-97]++
		count[t[i]-97]--
	}
	for i := 0; i < len(count); i++ {
		if count[i] > 0 {
			return false
		}
	}
	return true
}
