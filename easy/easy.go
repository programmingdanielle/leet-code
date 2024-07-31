package easy

func MergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	ans := make([]byte, 0, m+n)
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// gcd (greatest common divisor) -- can use the Euclidean algorithim to retrieve the gcd
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
func GcdOfStrings(str1 string, str2 string) string {
	// grab both lengths of the strings
	lengthOne, lengthTwo := len(str1), len(str2)

	// pass in lengths of strings to gcd helper to find greatest common divisor
	divisorLen := gcd(lengthOne, lengthTwo)

	// use the gcd as a means to identify the largest string that divides both str1 & str2
	divisor := str1[:divisorLen]

	// ensure that the divisor is correct by comparing the strings reversed; they should be the same
	if str1+str2 == str2+str1 {
		return divisor
	}

	// if the above statement isn't true, there is NOT a larger string that divides both str1 & str2
	return ""
}

// Euclidean algorithim
// an algorithim that replaces the larger of the two numbers by its remainder when divided by the smaller of the two
// algorithim will stop when it reaches a zero remainder
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
