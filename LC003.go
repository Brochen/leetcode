package main

import "fmt"

/**
3. Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	max := 1

	ms := 0
	me := ms + 1

	for ms < len(s) && me < len(s) {
		if len(s) - ms <= max {
			break
		}
		f := true
		for i := ms; i < me; i++ {
			if s[me] == s[i] {
				ms = i + 1
				me = me + 1
				f = false
				break
			}
		}
		if f {
			if me - ms + 1 > max {
				max = me - ms + 1
			}
			me += 1
		}
	}

	return max
}

func main()  {

	s := ""

	fmt.Println(lengthOfLongestSubstring(s))
}