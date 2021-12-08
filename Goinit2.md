# Goの変数
## 変数に型をつける
ひとまずコード
```
package main

import "fmt"

func main() {
	var greeting string
	greeting = "Hello World"
	fmt.Println(greeting)
}
```
varをつけて変数名を書く、その後ろに型をつける。<br>
今回であれば変数名に文字を代入する。<br>
ちなみに代入する値を途中で変えることもできる。
```
package main

import "fmt"

func main() {
	var greeting string
	greeting = "①Hello World"
	fmt.Println(greeting)
	greeting = "②Hello go"
	fmt.Println(greeting)
}
```
実行結果
```
①Hello World
②Hello go
```
上から読まれてますね。


