package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

/*
headコマンドに似たコマンドのgo実装
仕様は以下
- 引数で渡された1つもしくは複数のファイルの先頭の最大N行をそのまま出力
- Nはデフォルトでは10
- オプション-nでNを指定できる
 */
func main()  {
	var n = flag.Int("n", 10, "head [-n lines] [file ...]")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Input at least one file. Usage: head [-n lines] [file ...]")
	}

	for i, f := range flag.Args() {
		// 複数ファイルの場合、改行で出力を区切る
		if i > 0 {
			fmt.Printf("\n")
		}

		// ファイルオープン
		var file, err = os.Open(f)
		if err != nil {
			fmt.Println("Error occured: ", err.Error())
			os.Exit(1)
		}

		// 1ファイルだけの場合、ファイル名は出力しない
		if len(flag.Args()) != 1 {
			fmt.Printf("==>  %s  <==\n", f)
		}

		headCommand(*n, file)
	}
}

func headCommand(lines int, fp *os.File) {
	var scanner = bufio.NewScanner(fp)
	for i := 0; i < lines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	fp.Close()
}
