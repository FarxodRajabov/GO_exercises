package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a int
	var maxNum int
	numbers := make(map[string]int)
	fmt.Scan(&a)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	res := strings.Split(line, " ")
	for i := 0; i < a; i++ {
		numbers[res[i]] += 1
		if numbers[res[i]] > maxNum {
			maxNum = numbers[res[i]]
		}
	}
	fmt.Println(a - maxNum)
}
