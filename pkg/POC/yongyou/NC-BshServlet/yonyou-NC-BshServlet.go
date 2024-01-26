package NC_BshServlet

import (
	"bufio"
	_ "embed"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

//go:embed payload.txt
var Payload string

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(1 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(baseURL string) {
	scanner := bufio.NewScanner(strings.NewReader(Payload))
	for scanner.Scan() {
		line := scanner.Text()
		testUrl := baseURL + line                             // 创建一个新的 URL 用于测试
		resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
			"User-Agent": UA,
		}).Get(testUrl)
		if err != nil {
			continue // 如果出现错误，继续尝试下一行
		} else {
			if resp.Status == "200 OK" && strings.Contains(strings.ToLower(resp.String()), "beanshell") {
				color.Green.Println("[+] 用友 NC bsh.servlet.BshServlet 远程命令执行漏洞存在 -> " + testUrl)
				return
			}
		}
	}
	color.Red.Println("[-] 用友 NC bsh.servlet.BshServlet 远程命令执行漏洞不存在")
}
