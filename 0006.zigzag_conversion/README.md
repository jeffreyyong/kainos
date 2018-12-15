# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## Title

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```text
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

```go
func convert(text string, nRows int) string
```

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

## Thought process

After entering "ABCDEFGHIJKLMNOPQRSTUVWXYZ" and parameter 5, the answer "AGMSYBFHLNRTXZCEIKOQUWDJPV" is obtained.

```text
A   I   Q   Y
B  HJ  PR  XZ
C G K O S W
DF  LN  TV
E   M   U
```

It can be seen that the index number of each line character in the original string is 

1. 0行，0,    8,         16,       24
1. 1行，1,  7,9,      15,17,    23,25
1. 2行，2, 6, 10,  14,   18,  22
1. 3行，3,5,  11,13,     19,21
1. 4行，4,    12,        20

let p=numRows×2-2，can summarise the following rules

1. 0行， 0×p，1×p，...
1. r行， r，1×p-r，1×p+r，2×p-r，2×p+r，...
1. 最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p，...
