package A8_htmlofficeservlet_uploadfile

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
	url = url + "/seeyon/htmlofficeservlet"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(url)
	if err != nil {
		color.Red.Println("[-]	致远A8 htmlofficeservlet 任意文件上传漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(resp.String(), "DBSTEP") {
		color.Green.Println("[+]	致远A8_htmlofficeservlet 任意文件上传漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]	致远A8_htmlofficeservlet 任意文件上传漏洞不存在")
	}
}
