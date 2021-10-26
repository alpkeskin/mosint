package modules

import (
	"github.com/fatih/color"
	googlesearch "github.com/rocketlaunchr/google-search"
)

func Pastebin_search(email string) {
	q := "intext:" + string('"') + email + string('"') + " site:pastebin.com"
	res, _ := googlesearch.Search(nil, q)
	size := len(res)
	for i := 0; i < size; i++ {
		color.Green(res[i].URL)
	}
	q = "intext:" + string('"') + email + string('"') + " site:throwbin.io"
	res, _ = googlesearch.Search(nil, q)
	size2 := len(res)
	for i := 0; i < size2; i++ {
		color.Green(res[i].URL)
	}
	if size == 0 && size2 == 0 {
		color.Red("[-] Not Found")
	}
}

func Related_domains_google(email string) {
	res, _ := googlesearch.Search(nil, email)
	size := len(res)
	if size > 0 {
		println("From google search:")
	} else {
		color.Red("[-] Not Found!")
	}
	for i := 0; i < size; i++ {
		color.Green("[+] " + res[i].URL)
	}
}
