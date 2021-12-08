# Goのif
## ifの書き方
入力された文字列がgoだった場合Welcomと出力する
```
if name == "go" {
  fmt.Println("Welcome")
}
```
比較する対象を()で閉じない<br>
name `==` "go"<br>
==は`条件演算子`という

## 条件演算子の種類
| 演算時 | 例 | 読み方 | 役割 |
|---|---|---|---|
| == | a == b| イコールイコール | 等しい。代入区別する |
| != | a != b | エクスクラメーションイコール | 等しくない |
| > | a > b | 大なり | aはbより大きい。aがbに等しい場合は成立しない |
| < | a < b | 小なり | aはbより小さい。aがbに等しい場合は成立しない |
| >= | a >= b | 以上 | aはb以上。aがbに等しい場合は成立する |
| <= | a <= b | 以下 | aはb以下。aがbに等しい場合は成立する |

この辺はJSでも使うのと一緒ですが、型をつける文===とかはないっぽい。<br>

## else if
結論JSとかと一緒！
```
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner((os.Stdin))

	sc.Scan()
	number, _ := strconv.Atoi(sc.Text())
	fmt.Println((number))

	if number == 10 {
		fmt.Println(strconv.Itoa(number) + "は10に等しい")
	} else if number > 10 {
		fmt.Println(strconv.Itoa(number) + "は10より大きい")
	} else {
		fmt.Println((strconv.Itoa(number) + "は10未満"))
	}
}
```
`strconv.Itoa`は整数を文字列に変換
