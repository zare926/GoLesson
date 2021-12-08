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
	count, _ := strconv.Atoi(sc.Text())
	fmt.Println(count)

	for i := 0; i < count; i++ {
		sc.Scan()
		name := sc.Text()
		fmt.Println("Hello " + name)
	}
}
