// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package controllers

import (
	"github.com/alpkeskin/mosint/cmd/utils"
	googlesearch "github.com/rocketlaunchr/google-search"
)

func Googling(email string) {
	defer utils.ProgressBar.Add(10)
	q := "intext:" + string('"') + email + string('"')
	res, err := googlesearch.Search(nil, q)
	if err != nil {
		panic(err)
	}
	size := len(res)
	for i := 0; i < size; i++ {
		utils.Googling_result = append(utils.Googling_result, res[i].URL)
	}
}
