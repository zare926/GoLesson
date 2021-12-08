# Goを勉強していく
## Goとは
- シンプルな言語仕様、オブジェクト指向の継承を持っていない
- ネットワーク処理やマルチスレッド、並行処理に強い
- 標準ライブラリや開発用コマンドが充実
- Googleが中心となって開発
- システムツールの開発に利用されており、DockerなどもGoを利用している

<br>
<br>

## Goの導入
ターミナルからbrewを使ってインストールする。
```
brew install go
```
インストール後、以下コマンドで【コマンドが見つかりません】とならなければOK
```
go
```
あとは単純にversion確認でもOK
```
go version
```

## ビルドして実行
とりあえず代表的なHello Worldでやってみます。
```
//ファイル名 helloworld.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
このファイルをビルドしないと実行できません。<br>
ファイルがあるディレクトリまで行き、以下コマンドでビルドする。
```
go build helloworld.go
```
ビルド完了したら以下コマンドで実行
```
./helloworld.go
```

ここまではJavaと一緒
