// You have been given an array A of size N consisting of positive integers.
// You need to find and print the product of all the number in this array Modulo pow(10, 9)+7.

// Input Format:
// The first line contains a single integer N denoting the size of the array.
// The next line contains N space separated integers denoting the elements of the array

// Output Format:
// Print a single integer denoting the product of all the elements of the array Modulo pow(10, 9)+7.

// Constraints:
//  1 ≤ N ≤ pow(10, 3)
//  1 ≤ A[i] ≤ pow(10, 3)

//  SAMPLE INPUT
//  5
//  1 2 3 4 5

//  SAMPLE OUTPUT
//  120

package main

import "fmt"

func main() {

	l := 0
	fmt.Scan(&l)

	m := 1
	a := 0

	for i := 0; i < l; i++ {
		fmt.Scan(&a)
		m = (m * a) % 1000000007
	}
	fmt.Println(m)
}
