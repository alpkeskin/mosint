/*
Copyright © 2023 github.com/alpkeskin
*/
package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Twitter struct {
	Exists bool
}

type response struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

func New() *Twitter {
	return &Twitter{}
}

func (t *Twitter) Check(email string) {
	spinner := spinner.New("Twitter Account Checking")
	spinner.Start()

	twitterUrl := "https://api.twitter.com/i/users/email_available.json"
	data := url.Values{}
	data.Set("email", email)
	r, err := http.Get(twitterUrl + "?" + data.Encode())

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	if r.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Status code error: %d %s", r.StatusCode, r.Status)
		spinner.StopFail()
		spinner.SetMessage(msg)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	var response response
	err = json.Unmarshal(body, &response)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	if !response.Taken {
		t.Exists = false
		spinner.Stop()
		return
	}

	t.Exists = true
	spinner.Stop()
}

func (i *Twitter) Print() {
	if i.Exists {
		fmt.Println(color.GreenString("[+]"), "Twitter Account Exists", color.GreenString("\u2714"))
		return
	}
	fmt.Println(color.RedString("[!]"), "Twitter Account Not Exists", color.RedString("\u2718"))
}
