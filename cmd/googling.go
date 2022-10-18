// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"sync"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func Googling(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	q := "intext:" + string('"') + email + string('"')
	res, err := googlesearch.Search(nil, q)
	if err != nil {
		panic(err)
	}
	size := len(res)
	for i := 0; i < size; i++ {
		googling_result = append(googling_result, res[i].URL)
	}
}
