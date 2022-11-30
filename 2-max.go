package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var max1 = 1
	var max2 int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	resInt := strings.Split(line, " ")

	for i := 0; i < n; i++ {
		resVal, _ := strconv.Atoi(resInt[i]) //  max1=5 max2=1
		if resVal > max2 && resVal < max1 {
			max2 = resVal
		} else if resVal >= max1 {
			max2 = max1
			max1 = resVal
		}

	}
	fmt.Println(max2)
}
