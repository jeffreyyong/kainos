# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## Topic
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## Thought Process for Solution
The text[left:i+1] is used to represent the longest substring containing s[i] in s[:i+1]. Location[s[i]] is the serial number of the second occurence of the character s[i] in s[:i+1]. 
When left < location[s[i]], the character s[i] appears twice. need to set left = location[s[i]] + 1. The guaranteed character s[i] appears only once.


## Conclusion
Using location to save the last occurrence of the serial number of the character, avoiding the query work. Location is the same as m in [Two Sum](./algorithms/0001.two-sum)
```go
// m is responsible for saving the serial number of the map[int]int
	m := make(map[int]int, len(nums))
```
