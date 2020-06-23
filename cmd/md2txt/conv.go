package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

// ImportFile is インポートするためのファイル
type ImportFile struct {
	Items []PageItem `json:"pages"`
}

// PageItem is ページデータ
type PageItem struct {
	Title string   `json:"title"`
	Lins  []string `json:"lines"`
}

// ReadDir is フォルダ内探索
func ReadDir(dir string, tag string) ImportFile {
	var pg ImportFile
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

			pageinfo := Edittext(lines, tag)
			pg.Items = append(pg.Items, pageinfo)
		}
	}

	return pg
}

// Edittext is 本文加工用
func Edittext(txt []string, tags string) PageItem {
	//  ここに来るまでには普通のテキストなのでそれをURLエンコードとやらをする。
	jointxt := ""
	var pi PageItem
	v := url.Values{}
	//検索条件
	reg := `^# (.*)`
	r := regexp.MustCompile(reg)
	// for ii := 0; ii < len(txt); ii++ {
	for _, val := range txt {
		ret := r.FindStringSubmatch(val)
		// 0行目で先頭が#だけであればタイトルにする
		if len(pi.Title) == 0 {
			pi.Title = ret[1]
			pi.Lins = append(pi.Lins, ret[1])
		} else {

			pi.Lins = append(pi.Lins, val)
		}
		// fmt.Printf("行:%v 内容:%v\n", idx, val)
		jointxt += val + "\n"

	}

	pi.Lins = append(pi.Lins, tags)

	v.Set("body", jointxt)
	ret := v.Encode()
	fmt.Println(ret)
	return pi
}

// Writedata is ファイルへの書き出し
func Writedata(item ImportFile) {

	// jsonエンコード
	outputJSON, err := json.Marshal(&item)
	if err != nil {
		panic(err)
	}

	//書き込みファイル作成
	file, err := os.OpenFile("./input.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	fmt.Fprintln(file, string(outputJSON))

}
