package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// パラメータの分解
	targetdir := flag.String("d", "", "読み込みフォルダを指定。このフォルダ内のmdファイルを読み込みます")
	tag := flag.String("t", "", "フォルダ内のファイルに一律で設定したいキーワードを入力。e.g.映画")

	flag.Parse()
	log.Printf("args= %q %q\n", *targetdir, *tag)
	// フォルダから.mdを読み込む
	jsondata := ReadDir(*targetdir, *tag)
	fmt.Println(jsondata)

	// ファイルの書き出し
	Writedata(jsondata)
}
