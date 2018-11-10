# [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## Title

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4


Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1


## Thought process

First assume that old = [0, 1, 2, 4, 5, 6, 7], using the binary search method, it is easy to use the index number of `5`, when old transforms into new = [4, 5, 6, 7, 0, 1, 2], you can also use the binary search method because the elements in old and new have a clear correspondence.

Old[i] == new[j] as long as i and j satisfy the relationship

```go
j=i+4
If j > len(old) {
    j -= len(old)
}
```

Where the maximum value of 4 = old is index number in new + 1

Therefore, if we only have new in our hands, we can pretend that we are still using the binary search method for old. When we need to obtain the value of old[i] for comparison and judgment, we can use the value of new[j] instead.

## to sum up

This question is an upgraded version of the binary search method