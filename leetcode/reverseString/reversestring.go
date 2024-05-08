package reversestring

func reverseString(s []byte) {
	for i, j := len(s)-1, 0; j < i; i-- {
		s[i], s[j] = s[j], s[i]
		j++
	}
}
