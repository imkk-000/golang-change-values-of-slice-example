// Ref: https://stackoverflow.com/questions/49428716/pass-slice-as-function-argument-and-modify-the-original-slice/49429099#49429099
package main

import (
	"fmt"
)

func noReturn(a []int, data ...int) {
	a = append(a, data...)
}

func returnS(a []int, data ...int) []int {
	return append(a, data...)
}

func main() {
	a := make([]int, 1)
	noReturn(a, 1, 2, 3)
	fmt.Println(a, len(a), cap(a)) // append changes will not visible since slice size grew on demand changing underlying array

	a = returnS(a, 1, 2, 3)
	fmt.Println(a, len(a), cap(a)) // append changes will be visible here since your are returning the new updated slice
}
