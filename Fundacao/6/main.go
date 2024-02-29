package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Gabriel"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
