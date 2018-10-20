# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Title
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example:
```
Input: "cbbd"
Output: "bb"
```
## Thought process
Can use the following program to determine whether the string s[i:j+1] is a palindrome.
```go
func isPalindromic(s *string, i, j int ) bool {
    for  i < j {
        if (*s)[i] != (*s)[j] {
            return false
        } 
        i++
        j--
    }
    return true
}
```
However, this does not take advantage of the characteristics of the palindrome
1. When l is an odd number the 'median segment' of the palindrome will only be "x" or "xxx" or "xxxxx"
1. When l is an even number the 'median segment' of the palindrome will only be "xx" or "xxxx" or "xxxxxx"
1. The positive middle segment of any two palidnromes substring of the same string does not overlap