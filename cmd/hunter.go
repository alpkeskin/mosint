// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type HunterStruct struct {
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

func Hunter(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	key := GetAPIKey("Hunter")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.hunter.io/v2/domain-search?domain="+strings.Split(email, "@")[1]+"&api_key="+key, nil)
	req.Header.Set("User-Agent", "mosint")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &hunter_result)
}
