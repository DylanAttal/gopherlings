//
// Problem:
// Dividing an integer by zero is an undefined operation. To deal with
// this Go will `panic` to avoid undefined behaviour. Panics result in
// crashes if not caught by a recover(), more on that later.
//
// Furthermore, in Go variables that are declared without
// initialization are "zero-valued" For example, integers
// will be equal to zero when declared without assigning to them.

package main

import "fmt"

func main() {
	var z int

	if z == 0 {
		z = 7
	}

	fmt.Println(49/z + 7)
}
