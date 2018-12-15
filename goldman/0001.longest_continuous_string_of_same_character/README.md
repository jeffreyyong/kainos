## Title

Given two strings 'X' and 'Y', print the length of the longest common substring. If two or more substrings have the same value for longest common substring, then print any of them. 

Input :  X = "GeeksforGeeks", 
         Y = "GeeksQuiz"
Output : Geeks

Input : X = "zxabcdezy", 
        Y = "yzabcdezx"
Output : abcdez


## Thought Process

##### Naive Approach
Let strings X and Y be the lengths m and n respectively. Generate all possible substrings of X which requires a time complexity of O(m^2) and search each substring in the string Y which can be achieved in O(n) time complexity. Overall time complexity will be O(n * m^2)


##### Efficient Approach
The longest suffix matrix LCSuff[][] is built up and the index of the cell having the maximum value is tracked. Let that index be represented by (row, col) pair. Now the final longest common substring is built with the help of that index by diagonally traversing up the LCSuff[][] matrix until LCSuff[row][col] != 0 and during the iteration obtaining the characters either from X[row-1] or Y[col-1] and adding them from right to left in the resultant common string.

