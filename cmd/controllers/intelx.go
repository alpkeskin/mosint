// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package controllers

import (
	"context"
	"fmt"

	"github.com/IntelligenceX/SDK/Go/ixapi"
	"github.com/alpkeskin/mosint/cmd/utils"
)

func Intelx(email string) {
	defer utils.ProgressBar.Add(10)
	key := utils.GetAPIKey("Intelx")
	if key == "" {
		return
	}
	search(context.Background(), key, email, 2)
}

func search(ctx context.Context, Key, Selector string, Sort int) {

	search := ixapi.IntelligenceXAPI{}
	search.Init("", Key)
	results, selectorInvalid, err := search.Search(ctx, Selector, Sort, utils.IntelxDefaultMaxResults, ixapi.DefaultWaitSortTime, ixapi.DefaultTimeoutGetResults)

	if err != nil {
		fmt.Printf("Error querying results: %s\n", err)
		return
	} else if len(results) == 0 && selectorInvalid {
		return
	}

	text := generateResultText(ctx, &search, results)
	fmt.Println(text)
}

func generateResultText(ctx context.Context, api *ixapi.IntelligenceXAPI, Records []ixapi.SearchResult) (text string) {

	for n, record := range Records {
		resultLink := utils.IntelxURL + "?did=" + record.SystemID.String()

		title := record.Name
		if title == "" {
			title = "Untitled Document"
		}
		utils.Intelx_result = append(utils.Intelx_result, resultLink)

		if n >= utils.IntelxDefaultMaxResults-1 {
			break
		}
	}

	return
}
