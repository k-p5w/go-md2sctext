package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

// ImportFile is インポートするためのファイル
type ImportFile struct {
	pages string
	data  PageItem
}

// PageItem is ページデータ
type PageItem struct {
	title string
	lins  []string
}

// ReadDir is フォルダ内探索
func ReadDir(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// ファイルが配置されているフォルダを探す
	for _, file := range files {
		lines := make([]string, 0, 10)
		filename := file.Name()
		ext := filepath.Ext(filename)

		// mdファイルの場合
		if ext == ".md" {
			// filelen := len(filename) - len(ext)
			// pagename := filename[0:filelen]

			fullpath := filepath.Join(dir, filename)
			f, err := os.Open(fullpath)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			defer f.Close()
			// 関数return時に閉じる
			s := bufio.NewScanner(f)
			for s.Scan() {
				// テキストファイルの中身を出力する
				// fmt.Print(strconv.Quote(s.Text()))
				lines = append(lines, s.Text())
				// sliceに放り込む
			}
			if s.Err() != nil {
				// non-EOF error.
				log.Fatal(s.Err())
			}

			Edittext(lines)

		}
	}

}

// Edittext is 本文加工用
func Edittext(txt []string) string {
	//  ここに来るまでには普通のテキストなのでそれをURLエンコードとやらをする。
	jointxt := ""

	v := url.Values{}

	// for ii := 0; ii < len(txt); ii++ {
	for idx, val := range txt {

		fmt.Printf("行:%v 内容:%v\n", idx, val)
		jointxt += val + "\n"

	}

	v.Set("body", jointxt)
	ret := v.Encode()
	return ret
}
