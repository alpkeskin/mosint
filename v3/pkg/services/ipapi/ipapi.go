/*
Copyright © 2023 github.com/alpkeskin
*/
package ipapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Ipapi struct {
	Response IpApiResponse
}

type IpApiResponse struct {
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

func New() *Ipapi {
	return &Ipapi{}
}

func (i *Ipapi) GetInfo(email string) {
	spinner := spinner.New("IPApi Lookup")
	spinner.Start()

	domain := email[strings.IndexByte(email, '@')+1:]
	ip := domainToIP(domain)
	if strings.EqualFold(ip, "") {
		spinner.StopFail()
		spinner.SetMessage("IP not found")
		return
	}
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	req.Header.Add("User-Agent", "mosint")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	var r IpApiResponse
	err = json.Unmarshal(body, &r)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	i.Response = r
	spinner.Stop()
}

func (i *Ipapi) Print() {
	fmt.Println("[*] IP Lookup")
	fmt.Println(color.GreenString("[+]"), "IP:", i.Response.IP)
	fmt.Println(color.GreenString("[+]"), "City:", i.Response.City)
	fmt.Println(color.GreenString("[+]"), "Region:", i.Response.Region)
	fmt.Println(color.GreenString("[+]"), "Country:", i.Response.CountryName)
	fmt.Println(color.GreenString("[+]"), "Country Code:", i.Response.CountryCode)
	fmt.Println(color.GreenString("[+]"), "Timezone:", i.Response.Timezone)
	fmt.Println(color.GreenString("[+]"), "Organization:", i.Response.Org)
	fmt.Println(color.GreenString("[+]"), "ASN:", i.Response.Asn)
}

func domainToIP(domain string) string {
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String()
		}
	}
	return ""
}
