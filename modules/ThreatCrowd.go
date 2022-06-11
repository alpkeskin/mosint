// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

func doRequest(url string, kind string) []string {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	fasthttp.Do(req, resp)
	var data_array []string
	if resp.StatusCode() != 200 {
		return data_array
	}
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
					data_array = append(data_array, splt[i][1:])
				} else if i == (length - 1) {
					data_array = append(data_array, splt[i][:len(splt[i])-1])
				} else {
					data_array = append(data_array, splt[i])
				}
			}
		}
	}
	return data_array
}

func RelatedDomains(email string) []string {
	return doRequest("https://www.threatcrowd.org/searchApi/v2/email/report/?email="+email, "domains")
}

func RelatedEmails(email string) []string {
	return doRequest("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain="+strings.Split(email, "@")[1], "emails")
}
