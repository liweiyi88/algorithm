package romantointeger

func romanToInt(s string) int {

	st := make(map[uint8]int, 0)

	st['I'] = 1
	st['V'] = 5
	st['X'] = 10
	st['L'] = 50
	st['C'] = 100
	st['D'] = 500
	st['M'] = 1000

	var t int
	var lastRoman uint8

	for i := len(s) - 1; i >= 0; i-- {
		roman := s[i]
		v, ok := st[roman]

		if ok {
			if roman == 'I' && (lastRoman == 'V' || lastRoman == 'X') {
				t -= v
			} else if roman == 'X' && (lastRoman == 'L' || lastRoman == 'C') {
				t -= v
			} else if roman == 'C' && (lastRoman == 'D' || lastRoman == 'M') {
				t -= v
			} else {
				t += v
			}
		}

		lastRoman = roman
	}

	return t
}
