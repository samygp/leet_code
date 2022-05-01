package arrays

import "fmt"

func IsPalindrome() {
	str1, str2 := "tacocat", "tacoca"
	fmt.Println(fmt.Sprintf("IsPalindrome: %s: %v", str1, isPalindrome(str1)))
	fmt.Println(fmt.Sprintf("IsPalindrome: %s: %v", str2, isPalindrome(str2)))
}

func isPalindrome(daString string) bool {
	var bitList int32
	for _, char := range daString {
		shiftBits(charNumber(char), &bitList)
	}
	return bitList == 0 || ((bitList-1)&bitList) == 0
}

func shiftBits(numVal int32, bitList *int32) {
	var mask int32 = 1 << numVal
	if (*bitList & mask) == 0 {
		*bitList |= mask
	} else {
		*bitList = *bitList &^ mask
	}
}

func charNumber(c rune) int32 {
	if c < 91 {
		return c - 65
	}
	return c - 97
}
