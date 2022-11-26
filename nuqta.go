package main

import "fmt"

func main() {
	var n, a, b, c, d int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b, &c, &d)
		x := 2*c - a
		y := 2*d - b
		fmt.Println(x, y)
	}

}
