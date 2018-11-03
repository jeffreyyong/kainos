# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## Topic
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

## Thought process
Combine the two sorted lists, and still require sorting after the merge. The steps to complete the question are as follows:
1. First deal with the case where one of the lists is nil and return directly to the other chain, which simplifies the subsequent judgment conditions.
2. Set the link header head and the node pointer for the mobile node
3. Use the for loop to compare repeatedly, each time pick a smaller node and put it on node.Next
4. Process the remaining nodes in l1 or l2