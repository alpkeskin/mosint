/*
Copyright © 2023 github.com/alpkeskin
*/
package intelx

import (
	"context"
	"fmt"
	"strings"

	"github.com/IntelligenceX/SDK/Go/ixapi"
	"github.com/alpkeskin/mosint/v3/internal/config"
	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Intelx struct {
	Response []string
}

func New() *Intelx {
	return &Intelx{}
}

func (i *Intelx) Search(email string) {
	spinner := spinner.New("Breached Data Searching")
	spinner.Start()

	key := config.Cfg.Services.IntelXApiKey
	if strings.EqualFold(key, "") {
		spinner.StopFail()
		spinner.SetMessage("IntelX Api Key is empty")
		return
	}

	api := ixapi.IntelligenceXAPI{}
	ctx := context.Background()
	api.Init("", key)
	results, selectorInvalid, err := api.Search(ctx, email, 2, config.Cfg.Settings.IntelXMaxResults, ixapi.DefaultWaitSortTime, ixapi.DefaultTimeoutGetResults)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}
	if len(results) == 0 && selectorInvalid {
		spinner.Stop()
		return
	}

	response := generateResultText(ctx, &api, results)

	i.Response = append(i.Response, response...)
	spinner.Stop()
}

func (i *Intelx) Print() {
	key := config.Cfg.Services.IntelXApiKey
	if strings.EqualFold(key, "") {
		return
	}

	fmt.Println("[*] IntelX Search Results")
	if len(i.Response) == 0 {
		fmt.Println(color.RedString("[!]"), "No results found")
		return
	}

	for _, v := range i.Response {
		fmt.Println(color.GreenString("[+]"), v)
	}
}

func generateResultText(ctx context.Context, api *ixapi.IntelligenceXAPI, Records []ixapi.SearchResult) (response []string) {
	for n, record := range Records {
		url := fmt.Sprintf("https://intelx.io/?did=%s", record.SystemID.String())

		response = append(response, url)

		if n >= config.Cfg.Settings.IntelXMaxResults {
			break
		}
	}

	return response
}
