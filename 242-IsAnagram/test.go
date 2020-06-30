package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
		hashMap[t[i]]--
	}

	for _, count := range hashMap {
		if count != 0 {
			return false
		}
	}

	return true
}
