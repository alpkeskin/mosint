/*
Copyright © 2023 github.com/alpkeskin
*/
package breachdirectory

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

type BreachDirectory struct {
	Response BreachDirectoryResponse
}

type BreachDirectoryResponse struct {
	Success bool `json:"success"`
	Found   int  `json:"found"`
	Result  []struct {
		HasPassword bool     `json:"has_password"`
		Sources     []string `json:"sources"`
		Password    string   `json:"password,omitempty"`
		Sha1        string   `json:"sha1,omitempty"`
		Hash        string   `json:"hash,omitempty"`
	} `json:"result"`
}

func New() *BreachDirectory {
	return &BreachDirectory{}
}

func (b *BreachDirectory) Lookup(email string) {
	spinner := spinner.New("Breached Password Searching")
	spinner.Start()

	key := config.Cfg.Services.BreachDirectoryApiKey
	if strings.EqualFold(key, "") {
		spinner.StopFail()
		spinner.SetMessage("BreachDirectory Api Key is empty")
		return
	}

	url := fmt.Sprintf("https://breachdirectory.p.rapidapi.com/?func=auto&term=%s", email)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

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

	var response BreachDirectoryResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	b.Response = response
	spinner.Stop()
}

func (b *BreachDirectory) Print() {
	key := config.Cfg.Services.BreachDirectoryApiKey
	if strings.EqualFold(key, "") {
		return
	}

	if !b.Response.Success {
		fmt.Println(color.RedString("[!]"), "BreachDirectory Internal Error")
		return
	}

	fmt.Println("[*] Breached Sources and Passwords")
	for _, v := range b.Response.Result {
		for _, v2 := range v.Sources {
			fmt.Println(color.GreenString("[+]"), v2)
		}
		if v.HasPassword {
			fmt.Println("[*] Password")
			fmt.Println(color.GreenString("[+]"), v.Password)
			fmt.Println(color.GreenString("[+]"), v.Sha1)
			fmt.Println(color.GreenString("[+]"), v.Hash)
		}
		if !v.HasPassword {
			fmt.Println(color.RedString("[!]"), "Password Not Found")
		}
	}
}
