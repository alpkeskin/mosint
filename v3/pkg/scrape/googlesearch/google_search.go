/*
Copyright © 2023 github.com/alpkeskin
*/
package googlesearch

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
)

type GoogleSearch struct {
	Response []string
}

func New() *GoogleSearch {
	return &GoogleSearch{}
}

func (g *GoogleSearch) Search(email string) {
	spinner := spinner.New("Google Searching")
	spinner.Start()

	query := fmt.Sprintf("intext:'%s'", email)
	url := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	pattern := `https?://[^\s"']+`
	re := regexp.MustCompile(pattern)
	removedPrefix := "/url?q="
	removedParams := "&sa="

	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if containsGoogle(link) {
			return
		}

		var response []string
		if re.MatchString(link) {
			link := strings.TrimPrefix(link, removedPrefix)
			link = strings.Split(link, removedParams)[0]
			response = append(response, link)
		}

		g.Response = append(g.Response, response...)
	})

	err := c.Visit(url)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	spinner.Stop()
}

func (g *GoogleSearch) Print() {
	fmt.Println("[*] Google Search Results")
	if len(g.Response) == 0 {
		fmt.Println(color.RedString("[!]"), "No results found")
		return
	}

	for _, link := range g.Response {
		fmt.Println(color.GreenString("[+]"), link)
	}
}

func containsGoogle(text string) bool {
	return regexp.MustCompile(`google\.com`).MatchString(text)
}
