package main

import (
	"flag"
	"log"
)

func main() {
	targetdir := flag.String("d", "dir", "対象フォルダを指定")

	flag.Parse()
	log.Printf("listening on %q\n", *targetdir)
	// フォルダから.mdを読み込む
	targetdir2 := "C:/home/scrapbox/input"
	ReadDir(targetdir2)
}
