package main

var vowels = [10]string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

func reverseVowels(s string) string {
	lastVowel := len(s)
	for i := 0; i < lastVowel; i++ {
		iVowel := getVowel(s[i : i+1])
		if iVowel != "" {
			vowelFound := false
			for j := lastVowel - 1; j > i && !vowelFound; j-- {
				jVowel := getVowel(s[j : j+1])
				if jVowel != "" {
					if jVowel != iVowel {
						newS := s[:i] + jVowel + s[i+1:j] + iVowel
						if j != len(s)-1 {
							newS += s[j+1:]
						}
						s = newS
					}
					vowelFound = true
					lastVowel = j
				}
			}
		}
	}
	return s
}

func getVowel(s string) string {
	for _, v := range vowels {
		if v == s {
			return v
		}
	}
	return ""
}
