// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/IntelligenceX/SDK/Go/ixapi"
)

const defaultMaxResults = 10 // max results to query and show
const frontendBaseURL = "https://intelx.io/"

func Intelx(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	key := GetAPIKey("Intelx")
	if key == "" {
		return
	}
	search(context.Background(), key, email, 2)
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
		intelx_result = append(intelx_result, resultLink)

		if n >= defaultMaxResults-1 {
			break
		}
	}

	return
}
