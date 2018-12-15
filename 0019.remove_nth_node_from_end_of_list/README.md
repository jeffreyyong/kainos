# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 题目
Given a linked list, remove the nth node from the end of list and return its head.

For example,
```
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.
```
**Note**:
1. Given n will always be valid. There is no case where the length of the chain is < n
2. Try to do this in one pass.

## Thought Process
1. Get the parent node d of the nth node from the bottom
2. Then according to whether the parent node d is a head node, select different deletion methods.
