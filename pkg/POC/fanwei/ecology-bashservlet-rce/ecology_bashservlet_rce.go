package ecology_bashservlet_rce

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
	vulnUrl := url + "/weaver/bsh.servlet.BshServlet"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-02] 泛微OA E-Cology BshServlet 远程代码执行漏洞 不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-02] 泛微OA E-Cology BshServlet 远程代码执行漏洞 不存在")
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && strings.Contains(body, "BeanShell") {
		cheurl := url + "/weaver/bsh.servlet.BshServlet"
		color.Green.Println("[+02] 泛微OA E-Cology BshServlet 远程代码执行漏洞 存在 -> ", cheurl)
	} else {
		color.Red.Println("[-02] 泛微OA E-Cology BshServlet 远程代码执行漏洞 不存在")
	}
}
