package modules

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/valyala/fasthttp"
)

func DNS_lookup(email string) {
	splt := strings.Split(email, "@")

	data := [][]string{}
	iprecords, _ := net.LookupIP(splt[1])
	for _, ip := range iprecords {
		row := []string{"IP", ip.String()}
		data = append(data, row)
	}
	nameserver, _ := net.LookupNS(splt[1])
	for _, ns := range nameserver {
		row := []string{"NS", ns.Host}
		data = append(data, row)
	}
	mxrecords, _ := net.LookupMX(splt[1])
	for _, mx := range mxrecords {
		row := []string{"MX", mx.Host}
		data = append(data, row)
	}
	txtrecords, _ := net.LookupTXT(splt[1])
	for _, txt := range txtrecords {
		row := []string{"TXT", txt}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Record"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func IPapi(email string) {
	splt := strings.Split(email, "@")
	ips, _ := net.LookupIP(splt[1])
	ip4api := ""
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			ip4api = ipv4.String()
		}
	}
	if ip4api != "" {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)

		req.SetRequestURI("https://ipapi.co/" + ip4api + "/json/")

		fasthttp.Do(req, resp)

		bodyBytes := resp.Body()
		var dat map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}
		color.Magenta("IPapi.co data:")
		println(fmt.Sprintf("IP: %v", dat["ip"]))
		println(fmt.Sprintf("|-- City: %v", dat["city"]))
		println(fmt.Sprintf("|-- Region: %v", dat["region"]))
		println(fmt.Sprintf("|-- Region: %v", dat["region"]))
		println(fmt.Sprintf("|-- Country Name: %v", dat["country_name"]))
		println(fmt.Sprintf("|-- Country calling code: %v", dat["country_calling_code"]))
		println(fmt.Sprintf("|-- Timezone: %v", dat["timezone"]))
		println(fmt.Sprintf("|-- asn: %v", dat["asn"]))
		println(fmt.Sprintf("|-- org: %v", dat["org"]))
	}
}
