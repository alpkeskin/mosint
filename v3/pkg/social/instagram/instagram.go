/*
Copyright © 2023 github.com/alpkeskin
*/
package instagram

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Instagram struct {
	Exists bool
}

func New() *Instagram {
	return &Instagram{}
}

func (i *Instagram) Check(email string) {
	spinner := spinner.New("Instagram Account Checking")
	spinner.Start()

	token, err := getCSRFToken()

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	if strings.EqualFold(token, "") {
		spinner.StopFail()
		spinner.SetMessage("CSRF Token not found!")
		return
	}

	attempUrl := "https://www.instagram.com/accounts/web_create_ajax/attempt/"
	data := url.Values{}
	data.Set("email", email)
	r, err := http.NewRequest("POST", attempUrl, strings.NewReader(data.Encode())) // URL-encoded payload

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("Cookie", "csrftoken="+token+";")
	r.Header.Add("X-Csrftoken", token)

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Status code error: %d %s", res.StatusCode, res.Status)
		spinner.StopFail()
		spinner.SetMessage(msg)
		return
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	match, err := regexp.MatchString("email_is_taken", string(body))

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	if !match {
		i.Exists = false
		spinner.Stop()
		return
	}

	i.Exists = true
	spinner.Stop()
}

func (i *Instagram) Print() {
	if i.Exists {
		fmt.Println(color.GreenString("[+]"), "Instagram Account Exists", color.GreenString("\u2714"))
		return
	}
	fmt.Println(color.RedString("[!]"), "Instagram Account Not Exists", color.RedString("\u2718"))
}

func getCSRFToken() (string, error) {
	url := "https://www.instagram.com/accounts/emailsignup/"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Status code error: %d %s", res.StatusCode, res.Status)
		return "", errors.New(msg)
	}

	re := regexp.MustCompile(`(?m){\\"config\\":{\\"csrf_token\\":\\"(.*?)\\"`)
	match := re.FindStringSubmatch(string(body))

	if len(match) == 0 {
		return "", errors.New("CSRF Token not found")
	}

	return match[1], nil
}
