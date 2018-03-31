package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// 1行ずつ処理を行う場合
func readLine1() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		fmt.Println(buf.Text())
	}
}

func readCsv() {
	f, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	reader := csv.NewReader(f)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(line)

		for _, col := range line {
			fmt.Println(col)
		}
	}

}

func main() {
	readLine1()
	readCsv()
}
