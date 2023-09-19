/*
Copyright © 2023 github.com/alpkeskin
*/
package haveibeenpwned

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/alpkeskin/mosint/v3/internal/config"
	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type HaveIBeenPwned struct {
	Response HaveIBeenPwnedResponse
}

type HaveIBeenPwnedResponse []struct {
	Name               string    `json:"Name"`
	Title              string    `json:"Title"`
	Domain             string    `json:"Domain"`
	BreachDate         string    `json:"BreachDate"`
	AddedDate          time.Time `json:"AddedDate"`
	ModifiedDate       time.Time `json:"ModifiedDate"`
	PwnCount           int       `json:"PwnCount"`
	Description        string    `json:"Description"`
	LogoPath           string    `json:"LogoPath"`
	DataClasses        []string  `json:"DataClasses"`
	IsVerified         bool      `json:"IsVerified"`
	IsFabricated       bool      `json:"IsFabricated"`
	IsSensitive        bool      `json:"IsSensitive"`
	IsRetired          bool      `json:"IsRetired"`
	IsSpamList         bool      `json:"IsSpamList"`
	IsMalware          bool      `json:"IsMalware"`
	IsSubscriptionFree bool      `json:"IsSubscriptionFree"`
}

func New() *HaveIBeenPwned {
	return &HaveIBeenPwned{}
}

func (h *HaveIBeenPwned) Lookup(email string) {
	spinner := spinner.New("HaveIBeenPwned Lookup")
	spinner.Start()

	key := config.Cfg.Services.HaveIBeenPwnedApiKey
	if strings.EqualFold(key, "") {
		spinner.StopFail()
		spinner.SetMessage("HaveIBeenPwned Api Key is empty")
		return
	}

	url := fmt.Sprintf("https://haveibeenpwned.com/api/v3/breachedaccount/%s?truncateResponse=false", email)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	req.Header.Set("hibp-api-key", key)
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

	defer resp.Body.Close()
	var response HaveIBeenPwnedResponse
	json.Unmarshal(body, &response)
	h.Response = response
	spinner.Stop()
}

func (h *HaveIBeenPwned) Print() {
	key := config.Cfg.Services.HaveIBeenPwnedApiKey
	if strings.EqualFold(key, "") {
		return
	}

	if len(h.Response) == 0 {
		return
	}

	fmt.Println("[*] HaveIBeenPwned Breaches")
	for _, v := range h.Response {
		fmt.Println(color.GreenString("[+]"), "Name:", v.Name)
		fmt.Println(color.GreenString("[+]"), "Domain:", v.Domain)
		fmt.Println(color.GreenString("[+]"), "Breach Date:", v.BreachDate)
		fmt.Println(color.GreenString("[+]"), "DataClasses:")
		for _, d := range v.DataClasses {
			fmt.Println("    -", d)
		}
		fmt.Println()
	}
}
