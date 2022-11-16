package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[1:]
	printSlice(s)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	printSlice(s)

	kb := [2]string{"Penn", "Teller"}
	ab := &kb
	ac := kb
	ac[0] = "dhchdsvchds"
	fmt.Println(*ab)
	fmt.Println(ac)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
