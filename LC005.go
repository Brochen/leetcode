package main

import (
	"fmt"
)

/**
5. Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
 */


/**
马拉车算法
http://www.cnblogs.com/grandyang/p/4475985.html
 */
func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}
	ns := make([]byte, len(s) * 2 + 3, len(s) * 2 + 3)
	ns[0] = '$'
	for i := 1; i <= len(s); i++ {
		ns [i * 2 - 1] = '#'

		ns[i * 2] = s[i - 1]
	}
	ns[2 *len(s) + 1] = '#'
	ns[2 *len(s) + 2] = 0

	id, mx := 0, 0

	maxLength := 0
	retCenter := 0

	pi := make([]int, len(ns), len(ns))
	for i := 1; i < len(ns) - 1; i++ {
		if mx > i {
			pi[i] = mx - i;
			if pi[i] > pi[2 * id - i] {
				pi[i] = pi[2 * id - i]
			}
		} else {
			pi[i] = 1
		}

		for ns[i + pi[i]] == ns[i - pi[i]] {
			pi[i]++
		}
		if mx < pi[i] + i {
			mx = pi[i] + i
			id = i
		}

		if maxLength < pi[i] {
			maxLength = pi[i]
			retCenter = i
		}
	}

	index := (retCenter - maxLength) / 2;
	return string(s[index : index + maxLength - 1])
}
func longestPalindrome_2(s string) string {

	var n = len(s)
	if n <= 0 {
		return ""
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if isPalindrome(s, j, n - i + j - 1) {
				return string(s[j : n - i + j])
			}
		}
	}

	return ""
}

func isPalindrome(s string, start, end int) bool  {

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func main()  {

	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ab"))
	fmt.Println(longestPalindrome("aba"))
	fmt.Println(longestPalindrome("ababa"))
	fmt.Println(longestPalindrome("abcba"))
	fmt.Println(longestPalindrome("abba"))
	fmt.Println(longestPalindrome("cabcdcba"))
	fmt.Println(longestPalindrome("babad"))
}
