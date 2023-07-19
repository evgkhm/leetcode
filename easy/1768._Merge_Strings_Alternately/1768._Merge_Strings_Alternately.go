package _768__Merge_Strings_Alternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	idx1, idx2 := 0, 0
	for idx1+idx2 < len(word1)+len(word2) {
		if idx1 < len(word1) {
			res.WriteByte(word1[idx1])
			idx1++
		}

		if idx2 < len(word2) {
			res.WriteByte(word2[idx2])
			idx2++
		}
	}
	return res.String()
}
