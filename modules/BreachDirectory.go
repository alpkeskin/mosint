// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func BreachDirectory(email string) BreachDirectoryStruct {
	key := GetAPIKey("BreachDirectory.org API Key")
	url := "https://breachdirectory.p.rapidapi.com/?func=auto&term=" + email

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	data := BreachDirectoryStruct{}
	json.Unmarshal(body, &data)
	return data
}
