# [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)

## Topic

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

```text
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

## Thought process
In the current situation, you have several options. Then try every choice. If you have found that a certain choice is definitely not possible (because some of the rules are violated), then return; if a selection tries to find the correct solution, add it to the solution set.

There are two aspects to consider: 1. choices and limitations; 2. end conditions

For this question, you have two choices at any time:

1. Add the left parenthesis.
2. Add the right parenthesis.

There are also following restrictions:

1. If the left parenthesis has been used up, you can't add the left parenthesis.
2. If there are as many right and left parentheses as you have, you can't add right parentheses. Because the newly added right parenthesis must not match.



The end condition is:

1. The left and right parentheses have been used up.

Correctness after the end:
After the left and right parentheses are used up, they must be correct solutions. Because 1. There are as many left and right parentheses, 2. Each right parenthesis must have a left parenthesis paired with it. So once it is over, you can join the solution set.

Recursive function passed parameters:
There are "run out" and "as many" words in the limit and end conditions, so you need to know the number of left and right brackets.

Of course you still need to know the current situation substr and the solution set res.

Therefore, put the above ideas together is the code:

```text

If (left and right parentheses are exhausted) {
  Join the solution set and return
}
//Otherwise start experimenting with various options
If (left parenthesis left to be used) {
  Add a left parenthesis and continue recursion
}
If (the right parenthesis can be used, and the right parenthesis is smaller than the left parenthesis) {
  Add a closing parenthesis and continue recursion
}
```