# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## Topic
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

That is to say, there are many vertical line segments at the 1, 2, ..., n points on the x-axis, and the lengths are a1, a2, ..., an in order. Find the two segments so that they and x draw the largest area. The area formula is Min(ai, aj) X |j - i|


## Thought process 
The exhaustive method is the complexity of O(n^2), which triggers the time limit of leetcode.

 The solution to the complexity of O(n) is to hold two pointers i, j; respectively pointing to the beginning and end of the length array. If ai is less than aj, move i backwards (i++). Conversely, move j forward (j--). If the current area is larger than the recorded area, replace it. The basis of this idea is that if the length of i is less than j, no matter how j is moved, the short board is at i, it is impossible to find a larger value than the currently recorded area, and only by moving i to find a new possible larger area.