package modules

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

func doRequest(url string, kind string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	var dat map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &dat); err != nil {
		panic(err)
	}
	if dat["response_code"] == "1" {
		q := dat[kind]
		str := fmt.Sprintf("%v", q)
		splt := strings.Split(str, " ")
		length := len(splt)
		if length > 0 {
			for i := 0; i < length; i++ {
				if i == 0 {
					color.Green("[+] " + splt[i][1:])
				} else if i == (length - 1) {
					color.Green("[+] " + splt[i][:len(splt[i])-1])
				} else {
					color.Green("[+] " + splt[i])
				}
			}
		} else {
			color.Red("[-] Not found!")
		}
	} else {
		color.Red("[-] Not found!")
	}
}

func Related_domains(email string) {
	doRequest("https://www.threatcrowd.org/searchApi/v2/email/report/?email="+email, "domains")
}

func Subdomains(email string) {
	splt := strings.Split(email, "@")
	color.Magenta("Subdomains:")
	doRequest("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain="+splt[1], "subdomains")
}

func Related_emails(email string) {
	splt := strings.Split(email, "@")
	doRequest("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain="+splt[1], "emails")
}
