package main

func longestCommonPrefix(strArr []string) string {
	return longestCommonPrefixHelper(strArr, 0, len(strArr)-1)
}

func longestCommonPrefixHelper(strArr []string, start int, end int) string {
	if start == end {
		return strArr[start]
	}

	mid := (start + end) / 2
	left := longestCommonPrefixHelper(strArr, start, mid)
	right := longestCommonPrefixHelper(strArr, mid+1, end)

	i := 0
	str := ""
	for i < len(left) && i < len(right) {
		leftRune := []rune(left)
		rightRune := []rune(right)
		if leftRune[i] == rightRune[i] {
			str = str + string(leftRune[i])
		} else {
			break
		}
		i++
	}
	return str
}
