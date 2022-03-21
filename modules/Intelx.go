// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"context"
	"fmt"

	"github.com/IntelligenceX/SDK/Go/ixapi"
)

const defaultMaxResults = 10 // max results to query and show
const frontendBaseURL = "https://intelx.io/"

var urls []string

func Intelx(email string) []string {
	key := GetAPIKey("Intelx.io API Key")

	search(context.Background(), key, email, 2)

	return urls
}

func search(ctx context.Context, Key, Selector string, Sort int) {

	search := ixapi.IntelligenceXAPI{}
	search.Init("", Key)
	results, selectorInvalid, err := search.Search(ctx, Selector, Sort, defaultMaxResults, ixapi.DefaultWaitSortTime, ixapi.DefaultTimeoutGetResults)

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
		resultLink := frontendBaseURL + "?did=" + record.SystemID.String()

		title := record.Name
		if title == "" {
			title = "Untitled Document"
		}
		urls = append(urls, resultLink)

		if n >= defaultMaxResults-1 {
			break
		}
	}

	return
}
