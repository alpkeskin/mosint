// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package controllers

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/alpkeskin/mosint/cmd/utils"
)

func IPAPI(email string) {
	defer utils.ProgressBar.Add(10)
	splt := strings.Split(email, "@")
	ips, _ := net.LookupIP(splt[1])
	ip4api := ""
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ip4api = ipv4.String()
		}
	}
	if ip4api != "" {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", utils.IPAPIURL+ip4api+"/json/", nil)
		req.Header.Set("User-Agent", "mosint")
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &utils.Ipapi_result)
	}
}
