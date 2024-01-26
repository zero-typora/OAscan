package main

import (
	"OAScan/pkg/config" // 确保导入路径正确
	"OAScan/pkg/qi"
	"flag"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
)

var (
	Url    string
	file   string
	oaType int
)

func usage() {
	myFigure := figure.NewFigure("OA Scan", "", true)
	fmt.Println("\033[32m" + myFigure.String() + "\033[0m")
	fmt.Println("帮助:")
	fmt.Println("  -u url")
	fmt.Println("      指定目标 URL")
	fmt.Println("  -f targets.txt")
	fmt.Println("      指定包含多个目标的文件")
	fmt.Println("  -o 1")
	fmt.Println("      指定 OA 系统类型")
	fmt.Printf("        %-20s %s\n", "1 用友OA", "8 蓝凌")
	fmt.Printf("        %-20s %s\n", "2 致远OA", "9 金河")
	fmt.Printf("        %-20s %s\n", "3 帆软", "10 新点")
	fmt.Printf("        %-20s %s\n", "4 泛微", "11 智明")
	fmt.Printf("        %-20s %s\n", "5 万户", "12 华天")
	fmt.Printf("        %-20s %s\n", "6 金蝶OA", "13 万户")
	fmt.Printf("        %-20s %s\n", "7 通达", "14 信呼")
}

func main() {
	flag.StringVar(&Url, "u", "", "指定目标 URL")
	flag.StringVar(&file, "f", "", "指定包含多个目标的文件")
	flag.IntVar(&oaType, "o", 1, "指定 OA 系统类型")

	flag.Usage = usage
	flag.Parse()

	if Url == "" && file == "" {
		usage()
		os.Exit(0)
	}

	qi.Logo(Url)

	Sc := config.WorkExp{
		Url:    Url,
		OAType: oaType,
	}

	if file != "" {
		lines, err := qi.ReadLinesFromFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, line := range lines {
			Sc.Url = line
			qi.Scan(Sc.Url)
			Sc.ScanRun()
		}
	} else {
		Sc.ScanRun()
	}
}
