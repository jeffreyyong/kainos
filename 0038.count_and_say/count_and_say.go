package problem0038

import (
	"fmt"
)

func countAndSay(n int) string {
	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)
}

func say(buf []byte) []byte {
	// res length will not exceed twice the buf, so you can specify the capacity in advance to speed up append
	res := make([]byte, 0, len(buf)*2)

	i, j := 0, 1
	for i < len(buf) {
		// Use j to find the next different element
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}

		// fmt.Println(fmt.Sprintf("%v", byte(j-i+'0')))
		fmt.Println(string(byte(2 + '0')))
		// res in res[i] indicates the number of res[i+1], i is 0,2,4,6,...
		res = append(res, byte(j-i+'0'), buf[i])

		// Move i to j
		i = j
	}

	return res
}
