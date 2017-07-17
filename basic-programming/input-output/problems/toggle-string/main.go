// You have been given a String S consisting of uppercase and lowercase English alphabets.
// You need to change the case of each alphabet in this String.
// That is, all the uppercase letters should be converted to lowercase
// and all the lowercase letters should be converted to uppercase.
// You need to then print the resultant String to output.

// Input Format
// The first and only line of input contains the String S

// Output Format
// Print the resultant String on a single line.

// Constraints
//  1 ≤ |S| ≤ 100

//  where |S| denotes the length of string S.
//
//  SAMPLE INPUT
//  abcdE
//
//  SAMPLE OUTPUT
//  ABCDe

package main

import "fmt"

func main() {

	s := ""
	fmt.Scan(&s)
	S := ""

	// for _, r := range s {

	// sR := string(r)
	// c := strings.ToUpper(sR)
	// if c == string(r) {
	// 	c = strings.ToLower(sR)
	// }
	// S += c

	// }

	for _, r := range s {
		S += string(r ^ 32)
	}
	fmt.Println(S)
}
