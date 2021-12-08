# Goでデータを受け取る
## パッケージをimport
データの読み込みに必要なパッケージをimport
```
import (
	"bufio"
	"fmt"
	"os"
)
```
main関数
```
func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	name := sc.Text()
	fmt.Println("Hello " + name)
}
```
`bufio.NewScanner(os.Stdin)`これが入力タブになる<br>
  `Scan()`で入力タブからデータを受け取って、`Text()`が受け取ったテキストデータ<br>
  最終的に変数nameに代入していて、Hello 入力した名前となるからくり<br>
  Javaでもこの流れやるよね。
  ```
  package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	name := sc.Text()
	fmt.Println("Hello " + name)
}
  ```
  ターミナルへ入力
  ```
  zare
  ```
  結果
  ```
  Hello zare
  ```

## 整数を受け取る
この流れだと整数はstringへ変換が必要なようです。<br>
strconvを追加でimport<br>
Atoi関数で読み取った文字列を数値へ変換<br>
しれっと登場してますが、2つの変数に代入しています。<br>
2つ目の整数には、処理が正常に終わったか判定値が入るようですが、今回は使わないので_としてます。<br>
`ブランク識別子`というようです。
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
	name := sc.Text()
	fmt.Println("Hello " + name)

	sc.Scan()
	number, _ := strconv.Atoi(sc.Text())
	fmt.Println(number)
}
``` 