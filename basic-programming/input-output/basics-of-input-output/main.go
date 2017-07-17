// Read different types of data from standard input, process them as shown in output format and print the answer to standard output.

// Input format:
// First line contains integer N.
// Second line contains string S.

// Output format:
// First line should contain N x 2.
// Second line should contain the same string S.

// Constraints:
//  0 ≤ N ≤ 10
//  0 ≤ N ≤ 10
//  where |S| = length of string

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	if n <= 10 && n >= 0 {
		fmt.Println(n)
	}

	if len(s) <= 15 && len(s) >= 0 {
		fmt.Println(s)
	}

}
