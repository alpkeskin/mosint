// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/valyala/fasthttp"
)

type IPAPIStruct struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

func DNS_lookup(email string) *tablewriter.Table {
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
	//table.Render()
	return table
}

func IPAPI(email string) IPAPIStruct {
	data := IPAPIStruct{}
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

		json.Unmarshal(bodyBytes, &data)
	}
	return data
}
