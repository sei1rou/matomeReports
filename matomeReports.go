package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type course struct {
	code string
	name string
	count int
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Erro:", err)
	}
}

func main() {
	flag.Parse()

	//ログファイル準備
	logfile, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	failOnError(err)
	defer logfile.Close()

	log.SetOutput(logfile)

	log.Print("Start\r\n")

	// ファイルを読み込んで二次元配列に入れる
	records := readFile(flag.Arg(0))

	// コース別人数集計
	courseReports := courseCounter(records)
	// 項目別人数集計

	//ファイルへ書き出す
	saveFile(flag.Arg(0), records)

	log.Print("Finesh !\r\n")

}

func readFile(fileName string) [][]string {
	//入力ファイル準備
	infile, err := os.Open(fileName)
	failOnError(err)
	defer infile.Close()

	reader := csv.NewReader(transform.NewReader(infile, japanese.ShiftJIS.NewDecoder()))
	reader.Comma = '\t'

	//CSVファイルを２次元配列に展開
	readRecords := make([][]string, 0)
	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		readRecords = append(readRecords, record)
	}

	return readRecords
}

func saveFile(fileName string, saveRecords [][]string) {
	//出力ファイル準備
	outDir, outfileName := filepath.Split(fileName)
	pos := strings.LastIndex(outfileName, ".")
	outfile, err := os.Create(outDir + outfileName[:pos] + "d.txt")
	failOnError(err)
	defer outfile.Close()

	writer := csv.NewWriter(transform.NewWriter(outfile, japanese.ShiftJIS.NewEncoder()))
	writer.Comma = '\t'
	writer.UseCRLF = true

	for _, out_record := range saveRecords {
		writer.Write(out_record)
	}

	writer.Flush()

}

func courseCounter(counterRecords [][]string) string {
	courses := make([]course, 0)
	reports := ""

	codePos := 0
	namePos := 0
	recTitle := couterRecords[0]
	//コースコードの位置を確認
	for i,item := range recTitle {
		if item == "健診ｺｰｽcd" {
			codePos = i
		} else if item == "健診ｺｰｽ" {
			namePos = i
		}
	}

	if codePos == 0 {
		reports := "「健診ｺｰｽcd」がファイルにありません\r\n"
	}

	if namePos == 0 {
		reports := "「健診ｺｰｽ」がファイルにありません\r\n"
	}

	//コース
	if reports == "" {
		recLen := len(counterRecords)
		courseNum := 0
		courses
		for j==1, j < recLen; j++ {
			if courseNum - len(courses) = 0
		}
	}
}
