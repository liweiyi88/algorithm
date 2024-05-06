package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	charMap := make(map[int]map[rune]int, 0)

	for _, str := range strs {
		for i, char := range str {

			if charMap[i] == nil {
				charMap[i] = make(map[rune]int, 0)
			}

			charMap[i][char]++
		}
	}

	prefix := make([]rune, 0)
	firstStr := strs[0]

	for i, char := range firstStr {
		if charMap[i][char] == len(strs) {
			prefix = append(prefix, char)
		} else {
			break
		}
	}

	return string(prefix)
}

// a better solution
// func longestCommonPrefix(strs []string) string {
//     p := strs[0]
//     for _, s := range strs {
//         i := 0
//         for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {}
//         p = p[:i]
//     }
//     return p
// }
