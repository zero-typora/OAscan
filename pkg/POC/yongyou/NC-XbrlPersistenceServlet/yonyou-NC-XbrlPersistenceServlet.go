package ERP_NC_MLBL

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest()
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/service/~xbrl/XbrlPersistenceServlet"
	resp, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 用友NC XbrlPersistenceServlet反序列化漏洞不存在（该漏洞验证不准确建议使用dnslog直接打https://mp.weixin.qq.com/s/y1EMp6Q5W5fKJPAI9FiA0w）")
		return
	}
	if resp.Status == "200 OK" {
		color.Green.Println("[+] 用友NC XbrlPersistenceServlet反序列化漏洞存在该漏洞验证不准确建议使用dnslog直接打https://mp.weixin.qq.com/s/y1EMp6Q5W5fKJPAI9FiA0w -> " + url)
		return
	}
	color.Red.Println("[-] 用友NC XbrlPersistenceServlet反序列化漏洞不存在（该漏洞验证不准确建议使用dnslog直接打）exp（https://mp.weixin.qq.com/s/y1EMp6Q5W5fKJPAI9FiA0w）")
}
