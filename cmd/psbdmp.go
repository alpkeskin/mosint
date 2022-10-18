// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type PsbdmpStruct struct {
	Search string `json:"search"`
	Count  int    `json:"count"`
	Data   []struct {
		ID     string `json:"id"`
		Tags   string `json:"tags"`
		Length int    `json:"length"`
		Time   string `json:"time"`
		Text   string `json:"text"`
	} `json:"data"`
}

func Psbdmp(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var key string = GetAPIKey("Psbdmp")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://psbdmp.ws/api/v3/search/"+email+"?key="+key, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "mosint")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &psbdmp_result)
}
