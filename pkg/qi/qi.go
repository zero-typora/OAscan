package qi

import (
	"bufio"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"os"
	"strings"
)

var (
	green = []*color.Style256{color.S256(46), color.S256(47), color.S256(48), color.S256(49), color.S256(50), color.S256(51)}
)

// ReadLinesFromFile 从指定的文件中逐行读取内容
func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件错误: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("文件读取错误: %v", err)
	}

	return lines, nil
}

// gradient 创建一个彩色渐变的文本字符串
func gradient(text string, colorArr []*color.Style256) string {
	lines := strings.Split(text, "\n")

	var output string
	t := len(text) / 6
	i, j := 0, 0

	for l := 0; l < len(lines); l++ {
		str := strings.Split(lines[l], "")
		for _, x := range str {
			j++
			output += colorArr[i].Sprint(x)
			if j > t {
				i++
				j = 0
			}
		}
		if len(lines) != 0 {
			output += "\n"
		}
	}

	return strings.TrimRight(output, "\n")
}

// randomPattern 生成图案
func randomPattern() string {
	myFigure := figure.NewFigure("OA Scan", "", true)
	return myFigure.String()
}

// Logo 打印带有彩色渐变的随机图案，并显示OA漏洞综合检查工具信息
// Logo 打印带有彩色渐变的随机图案，并显示OA漏洞综合检查工具信息
func Logo(url string) {

	fmt.Println(gradient(randomPattern(), green))
	fmt.Println("OA漏洞综合检查工具v1.0")
	fmt.Println("data:2024-01-16")
}
func Scan(url string) {
	fmt.Printf("Scan: %s\n", url)
}
