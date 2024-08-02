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

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// 0 is empty
	// 1 is not empty
	// number of new flowers to be planted

	count := 0

	for i, flower := range flowerbed {
		if flower == 1 {
			continue
		}

		// confused why test cases on this fail when they are false instead of true
		leftIsZero := true
		rightIsZero := true
		if i > 0 {
			if flowerbed[i-1] == 1 {
				leftIsZero = false
			}
		}
		if i < len(flowerbed)-1 {
			if flowerbed[i+1] == 1 {
				rightIsZero = false
			}
		}
		if leftIsZero && rightIsZero {
			count++
			flowerbed[i] = 1
		}
	}
	return n <= count
}

// func KidsWithCandies(candies []int, extraCandies int) []bool {
// 	// create array called addedCandies of all candies + extraCandies
// 	// compare addedCandies against original candies ....
// 	// if the iteration of addedCandies > compared to all candies,
// 	// append a true bool to a bool array

// 	addedCandies := []int{}

// 	for i := range candies {
// 		addedCandies = append(addedCandies, candies[i]+extraCandies)
// 	}

// 	boolAnswer := []bool{}
// 	// boolHolder := true

// 	type Test struct {
// 		int
// 		bool
// 	}

// 	// mapForStuff := make(map[int]Test)

// 	testBool := true

// 	for d, i := range addedCandies {
// 		for _, b := range candies {
// 			// if addedCandies[i] IS NOT greater compared to candies[b]....
// 			if addedCandies[i] > candies[b] {
// 				// mapForStuff[addedCandies[d]] = Test{addedCandies[i], false}

// 		}
// 	}

// 	fmt.Println(mapForStuff)
// 	fmt.Println(addedCandies)
// 	fmt.Println("expectation: [true,true,true,false,true]")

// 	return boolAnswer
// }
