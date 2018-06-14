package main

import "fmt"

/**
	寻找最长不含有重复字符的字串
*/
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefghjk"))
	fmt.Println(lengthOfNonRepeatingSubStr("测试中文"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	// stores last occurred pos +1
	// 0 means not seen
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}

	return maxLength
}
