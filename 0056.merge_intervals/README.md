# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## Topic

Given a collection of intervals, merge all overlapping intervals.

```text
For example,
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

## 解题思路


1. Sort the intervals first, follow the Start increment
2. Process the overlapping cases in turn.