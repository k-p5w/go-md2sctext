package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	targetdir := flag.String("d", "dir", "対象フォルダを指定")
	tag := flag.String("tag", "tag", "一律に設定したい文字列を記載")

	flag.Parse()
	log.Printf("args= %q %q\n", *targetdir, *tag)
	// フォルダから.mdを読み込む
	targetdir2 := "C:/home/scrapbox/input"
	tags := "#映画 "
	jsondata := ReadDir(targetdir2, tags)
	fmt.Println(jsondata)

	Writedata(jsondata)
}
