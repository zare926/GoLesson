package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	name := sc.Text()
	fmt.Println("Hello " + name)

	sc.Scan()
	number, _ := strconv.Atoi(sc.Text())
	fmt.Println(number)
}
