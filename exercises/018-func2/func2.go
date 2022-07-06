//
// Problem:
// Functions may have multiple arguments and multiple return values.
// To create a function with multiple return values surround the
// return types in parentheses.
//
//  func fruits() (string, string) {
//  	return "banana", "tangerine"
//  }
//
// If two or more consecutive arguments have the same type one can omit
// the type name until the last argument.
//
//  func divmod(a, b int) (div, mod int) {
//  	div = a/b
//  	mod = a%b
//  	return div, mod
//  }
//

package main

import "fmt"

func main() {
	fmt.Println(return1And2(1, 2))
}

func return1And2(a, b int) (int, int) {
	return a, b
}
