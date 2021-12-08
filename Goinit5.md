# Go繰り返し処理
## for
```
for i := 0; i < 3; i++ {
  fmt.Println("Hello world")
}
```
変数を一つ用意して繰り返し使う。<br>
↑でいう`i`は
`カウンター変数`という<br>
慣習としてiやをを使う<br>

## 例
```
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
```
最初に整数(count)を入れて、その後に文字列(name)を入れる。