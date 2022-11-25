package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	res := math.Mod(a, b)
	if res != 0 {
		res = math.Floor(a/b) + 1
		fmt.Println(math.Floor(res))
		return
	}
	fmt.Println(math.Floor(a / b))
}
