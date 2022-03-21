// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	googlesearch "github.com/rocketlaunchr/google-search"
)

func BinSearch(email string) []string {
	var url_array []string
	q := "intext:" + string('"') + email + string('"') + " site:pastebin.com"
	res, _ := googlesearch.Search(nil, q)
	size := len(res)
	for i := 0; i < size; i++ {
		url_array = append(url_array, res[i].URL)
	}
	q = "intext:" + string('"') + email + string('"') + " site:throwbin.io"
	res, _ = googlesearch.Search(nil, q)
	size2 := len(res)
	for i := 0; i < size2; i++ {
		url_array = append(url_array, res[i].URL)
	}
	return url_array
}

func Related_domains_from_google(email string) []string {
	var url_array []string
	res, _ := googlesearch.Search(nil, email)
	size := len(res)
	for i := 0; i < size; i++ {
		url_array = append(url_array, res[i].URL)
	}
	return url_array
}
