/*
Copyright © 2023 github.com/alpkeskin
*/
package psbdmp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Psbdmp struct {
	Urls []string
}

func New() *Psbdmp {
	return &Psbdmp{}
}

func (p *Psbdmp) Search(email string) {
	spinner := spinner.New("Pastebin Dumps Searching")
	spinner.Start()

	type emailData struct {
		ID     string `json:"id"`
		Tags   string `json:"tags"`
		Length int    `json:"length"`
		Time   string `json:"time"`
		Text   string `json:"text"`
	}

	url := fmt.Sprintf("https://psbdmp.ws/api/v3/search/%s", email)
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

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	var data []emailData
	err = json.Unmarshal([]byte(body), &data)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	for _, v := range data {
		binUrl := fmt.Sprintf("https://psbdmp.ws/dump/%s", v.ID)
		p.Urls = append(p.Urls, binUrl)
	}

	spinner.Stop()
}

func (p *Psbdmp) Print() {
	if len(p.Urls) == 0 {
		return
	}

	fmt.Println("[*] Pastebin Dumps")
	for _, v := range p.Urls {
		fmt.Println(color.GreenString("[+]"), v)
	}
}
