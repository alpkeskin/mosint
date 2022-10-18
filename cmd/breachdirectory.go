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

type BreachDirectoryStruct struct {
	Success bool `json:"success"`
	Found   int  `json:"found"`
	Result  []struct {
		HasPassword bool     `json:"has_password"`
		Sources     []string `json:"sources"`
		Password    string   `json:"password,omitempty"`
		Sha1        string   `json:"sha1,omitempty"`
		Hash        string   `json:"hash,omitempty"`
	} `json:"result"`
}

func BreachDirectory(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var key string = GetAPIKey("Breachdirectory")
	if key == "" {
		return
	}
	url := "https://breachdirectory.p.rapidapi.com/?func=auto&term=" + email

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &breachdirectory_result)
}
