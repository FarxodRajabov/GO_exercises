package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	res := strings.Split(line, " ")

	for i := 0; i < a; i++ {
		resIn, _ := strconv.Atoi(res[i])
		if resIn != i+1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
