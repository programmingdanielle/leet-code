package easy

import "fmt"

func MergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	ans := make([]byte, 0, m+n)
	fmt.Println("before", ans)
	for i := 0; i < m || i < n; i++ {
		if i < m {
			ans = append(ans, word1[i])
		}
		if i < n {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

// // TAKEAWAY/NOTEWORTHY MENTIONS:
// * Go does not have a data type for characters.
// * Instead of characters, it uses bytes & runes to represent a character value.
// * String is a read-only slice of bytes
// * Indexing a string yields its bytes, not characters
// initializing a slice -- make([]dataType, length, *optional* capacity
