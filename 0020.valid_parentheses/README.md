# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## Topic
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

## Thought Process
Stack is a structure of last in, first out (LIFO), which can be used here to avoid complicated judgment structures. However, the standard library of the Go language does not have a stack structure, so I implemented one manually.