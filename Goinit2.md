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
	greeting = "Hello World"
	fmt.Println(greeting)
	greeting = "Hello go"
	fmt.Println(greeting)
}
```
実行結果
```
Hello World
Hello go
```
上から読まれてますね。<br>
代入する値が整数ならintの型を付ける。
```
var number int
number = 100
```

## 型推論
var 変数名 型を書かずとも **:**  で型を推測してくれる
```
package main

import "fmt"

func main() {
	greeting := "Hello World"
	fmt.Println(greeting)
	greeting = "Hello go"
	fmt.Println(greeting)
  number := 100
	fmt.Println(number)
}
```
ちなみに同じ変数に別の値を再代入する時は:不要のようです。

## 変数名の付け方
- 1文字目:  アルファベット、アンダーバー
- 2文字目以降: アルファベット、アンダーバー、`数字`
- 日本語は使わない！
- ハイフンで繋がない
- 1文字目が`小文字`: `そのパッケージだけで`使える変数
- 1文字目が`大文字`: `他のパッケージからも`みえる変数

## 実例
- greeting
- number1
- playerName(キャメルケース)