# markdownファイルをscrapboxに取り込むためのツール / README

markdownファイルをscrapboxに取り込むために作った便利ツールです。

## Features / どんなことができる

### このツールで行っていること
指定フォルダのmdファイルをinput.jsonに出力する

### 使い方

- 以下のように実行するとフォルダにあるmd拡張子のファイルを読み込んで、ツールと同じフォルダにinput.jsonというファイルを作ります
> .\md2txt.exe -d C:\home\scrapbox\input -t 映画

#### 仕様というか備忘録

動作時に同名ファイルがあったら事前にリネームしています

input.jsonのレイアウトとかはこちらのページを参考に作ってます
- https://scrapbox.io/help-jp/ページをインポート・エクスポートする

## Requirements /　使うために必要なもの

なし

## Release Notes

2020/06/28:じぶんのscrapboxに取り込むファイルを作ったのでこれで終わりとする。

### 0.0.1

初回リリース。
