/*
Copyright © 2023 github.com/alpkeskin
*/
package hunter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/config"
	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Hunter struct {
	Response HunterResponse
}

type HunterResponse struct {
	Data struct {
		Domain       string      `json:"domain"`
		Disposable   bool        `json:"disposable"`
		Webmail      bool        `json:"webmail"`
		AcceptAll    bool        `json:"accept_all"`
		Pattern      string      `json:"pattern"`
		Organization string      `json:"organization"`
		Country      string      `json:"country"`
		State        interface{} `json:"state"`
		Emails       []struct {
			Value      string `json:"value"`
			Type       string `json:"type"`
			Confidence int    `json:"confidence"`
			Sources    []struct {
				Domain      string `json:"domain"`
				URI         string `json:"uri"`
				ExtractedOn string `json:"extracted_on"`
				LastSeenOn  string `json:"last_seen_on"`
				StillOnPage bool   `json:"still_on_page"`
			} `json:"sources"`
			FirstName    string      `json:"first_name"`
			LastName     string      `json:"last_name"`
			Position     string      `json:"position"`
			Seniority    string      `json:"seniority"`
			Department   string      `json:"department"`
			Linkedin     interface{} `json:"linkedin"`
			Twitter      interface{} `json:"twitter"`
			PhoneNumber  interface{} `json:"phone_number"`
			Verification struct {
				Date   string `json:"date"`
				Status string `json:"status"`
			} `json:"verification"`
		} `json:"emails"`
		LinkedDomains []interface{} `json:"linked_domains"`
	} `json:"data"`
	Meta struct {
		Results int `json:"results"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Params  struct {
			Domain     string      `json:"domain"`
			Company    interface{} `json:"company"`
			Type       interface{} `json:"type"`
			Seniority  interface{} `json:"seniority"`
			Department interface{} `json:"department"`
		} `json:"params"`
	} `json:"meta"`
}

func New() *Hunter {
	return &Hunter{}
}

func (h *Hunter) Lookup(email string) {
	spinner := spinner.New("Related Emails Searching")
	spinner.Start()

	key := config.Cfg.Services.HunterApiKey
	if strings.EqualFold(key, "") {
		spinner.StopFail()
		spinner.SetMessage("Hunter Api Key is empty")
		return
	}

	domain := email[strings.IndexByte(email, '@')+1:]
	url := fmt.Sprintf("https://api.hunter.io/v2/domain-search?domain=%s&api_key=%s", domain, key)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	req.Header.Set("User-Agent", "mosint")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	defer resp.Body.Close()
	var response HunterResponse
	json.Unmarshal(body, &response)
	h.Response = response
	spinner.Stop()
}

func (h *Hunter) Print() {
	key := config.Cfg.Services.HunterApiKey
	if strings.EqualFold(key, "") {
		return
	}

	fmt.Println("[*] Hunter.io Search Results")

	fmt.Println(color.GreenString("[+]"), "Disposable:", h.Response.Data.Disposable)
	fmt.Println(color.GreenString("[+]"), "Webmail:", h.Response.Data.Webmail)
	fmt.Println(color.GreenString("[+]"), "AcceptAll:", h.Response.Data.AcceptAll)
	fmt.Println(color.GreenString("[+]"), "Pattern:", h.Response.Data.Pattern)
	fmt.Println(color.GreenString("[+]"), "Organization:", h.Response.Data.Organization)
	fmt.Println(color.GreenString("[+]"), "Country:", h.Response.Data.Country)

	fmt.Println("[*] Related Emails:")
	if len(h.Response.Data.Emails) == 0 {
		fmt.Println(color.RedString("[!]"), "No related emails found")
		return
	}
	for _, v := range h.Response.Data.Emails {
		fmt.Println(color.GreenString("[+]"), v.Value)
	}
}
