package main

import "fmt"

/**
6. ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
 */
func convert(s string, numRows int) string {
	if len(s) <= 0 || numRows <= 1 {
		return  s
	}

	cn := 2 * numRows - 2

	rs := make([]byte, 0, len(s))

	for i := 0; i < numRows; i++ {

		j := i
		for j < len(s) {
			rs = append(rs, s[j])
			if i == 0 || i == numRows - 1 {
				//第一行或最后一行
				//nothing todo
			} else {
				t := j + (numRows - i - 1) * 2
				if t < len(s) {
					rs = append(rs, s[t])
				}
			}
			j += cn
		}
	}

	return string(rs)
}
func convert2(s string, numRows int) string {
	if len(s) <= 0 || numRows <= 1 {
		return  s
	}
	//初始化一个二维数组

	sn := 2 * numRows - 2;

	columns := (len(s) / sn) * (numRows - 1)
	if len(s) % sn != 0 {
		columns += sn
	}


	ns := make([][]byte, numRows, numRows)
	for i := 0; i < numRows; i++ {
		ns[i] = make([]byte, columns, columns)
		for j := 0; j< columns; j++ {
			ns[i][j] = 0
		}
	}

	c, r := 1, 1
	direction := true //down
	for i := 0; i < len(s); i++ {
		fmt.Println(r, c)
		ns[r - 1][c - 1] = s[i]
		if direction {
			r++
		} else {
			r--
			c++
		}
		if r == numRows {
			direction = false
		} else if r == 1 {
			direction = true
		}
	}

	fmt.Println(ns)
	rets := make([]byte, len(s), len(s))
	index := 0
	for i:= 0; i < numRows; i++ {
		for j := 0; j < columns; j++ {
			if ns[i][j] > 0 {
				rets[index] = ns[i][j]
				index++
			}
		}
	}

	return string(rets)
}

func main()  {
	//fmt.Println(convert("A", 1))
	fmt.Println(convert("PAYPALISHIRING", 3))
}