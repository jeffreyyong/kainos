# [31. Next Permutation](https://leetcode.com/problems/next-permutation/)

## Topic

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

```text
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

## 解题思路

1　　2　　7　　4　　3　　1

Next permutation:

1　　3　　1　　2　　4　　7

Steps are as follows:

From the back to the front, find the longest descending order

1　　2　　`7　　4　　3　　1`

Arrange this in descending order and convert it into ascending order

1　　2　　`1　　3　　4　　7`

The element before the sequence is interchanged with the first element
in the sequence that is larger than it
1　　`3`　　1　　`2`　　4　　7
