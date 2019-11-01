/*
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.
Example 1:
Input: a = "11", b = "1"
Output: "100"
Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
*/
package add_binary

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var result = ""
	carry := uint(0)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		aBit := uint(0)
		bBit := uint(0)
		if i >= 0 {
			aBit = uint(a[i] - '0')
		}
		if j >= 0 {
			bBit = uint(b[j] - '0')
		}
		sum := aBit + bBit + carry
		if sum >= 2 {
			carry = 1
		} else {
			carry = 0
		}
		sum = sum % 2
		result = strconv.Itoa(int(sum)) + result
	}
	if carry > 0 {
		return "1" + result
	} else {
		return result
	}
}
