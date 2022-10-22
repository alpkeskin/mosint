// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

type adobeResponse []struct {
	HasT2ELinked bool `json:"hasT2ELinked"`
}

type discordResponse struct {
	Errors struct {
		Email struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"email"`
	} `json:"errors"`
}

type spotifyResponse struct {
	Status int `json:"status"`
	Errors struct {
		Email string `json:"email"`
	} `json:"errors"`
	Country                        string `json:"country"`
	CanAcceptLicensesInOneStep     bool   `json:"can_accept_licenses_in_one_step"`
	RequiresMarketingOptIn         bool   `json:"requires_marketing_opt_in"`
	RequiresMarketingOptInText     bool   `json:"requires_marketing_opt_in_text"`
	MinimumAge                     int    `json:"minimum_age"`
	CountryGroup                   string `json:"country_group"`
	SpecificLicenses               bool   `json:"specific_licenses"`
	TermsConditionsAcceptance      string `json:"terms_conditions_acceptance"`
	PrivacyPolicyAcceptance        string `json:"privacy_policy_acceptance"`
	SpotifyMarketingMessagesOption string `json:"spotify_marketing_messages_option"`
	PretickEula                    bool   `json:"pretick_eula"`
	ShowCollectPersonalInfo        bool   `json:"show_collect_personal_info"`
	UseAllGenders                  bool   `json:"use_all_genders"`
	UseOtherGender                 bool   `json:"use_other_gender"`
	DateEndianness                 int    `json:"date_endianness"`
	IsCountryLaunched              bool   `json:"is_country_launched"`
	AllowedCallingCodes            []struct {
		CountryCode string `json:"country_code"`
		CallingCode string `json:"calling_code"`
	} `json:"allowed_calling_codes"`
	PushNotifications bool `json:"push-notifications"`
}

type twitterResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

func Adobe(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var endpoint string = "https://auth.services.adobe.com/signin/v2/users/accounts"

	var jsonStr = []byte(`{"username":"` + email + `"}`)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr)) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("X-Ims-Clientid", "adobedotcom2")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response adobeResponse
		json.Unmarshal(body, &response)
		if len(response) > 0 {
			social_result = append(social_result, "Adobe \U0001f440")

		} else {
			social_result = append(social_result, "Adobe [Not here!]")
		}
	} else {
		social_result = append(social_result, "Adobe [Couldn't check!]")
	}
}

func Discord(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var endpoint string = "https://discord.com/api/v9/auth/register"

	var jsonStr = []byte(`{"email":"` + email + `","username":"asdsadsad","password":"q1e31e12r13*","invite":null,"consent":true,"date_of_birth":"1973-05-09","gift_code_sku_id":null,"captcha_key":null,"promotional_email_opt_in":false}`)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr)) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("X-Debug-Options", "bugReporterEnabled")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response discordResponse
		json.Unmarshal(body, &response)
		if len(response.Errors.Email.Errors) > 0 {
			if response.Errors.Email.Errors[0].Code == "EMAIL_ALREADY_REGISTERED" {
				social_result = append(social_result, "Discord \U0001f440")
			} else {
				social_result = append(social_result, "Discord [Not here!]")
			}
		} else {
			social_result = append(social_result, "Discord [Not here!]")
		}
	} else if res.StatusCode == 429 {
		social_result = append(social_result, "Discord [Rate limited!]")
	} else {
		social_result = append(social_result, "Discord [Couldn't check!]")
	}
}

func getCSRFToken() string {
	var url string = "https://www.instagram.com/accounts/emailsignup/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		re := regexp.MustCompile(`(?m){\\"config\\":{\\"csrf_token\\":\\"(.*?)\\"`)
		match := re.FindStringSubmatch(string(body))
		if len(match) > 0 {
			return match[1]
		}
	}
	return ""
}

func Instagram(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var token string = getCSRFToken()
	if token == "" {
		social_result = append(social_result, "Instagram [Couldn't check!]")
	} else {
		var endpoint string = "https://www.instagram.com/accounts/web_create_ajax/attempt/"

		data := url.Values{}
		data.Set("email", email)

		client := &http.Client{}
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
		r.Header.Add("Cookie", "csrftoken="+token+";")
		r.Header.Add("X-Csrftoken", token)

		res, err := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			match, _ := regexp.MatchString("email_is_taken", string(body))
			if match {
				social_result = append(social_result, "Instagram \U0001f440")
			} else {
				social_result = append(social_result, "Instagram [Not here!]")
			}
		} else {
			social_result = append(social_result, "Instagram [Couldn't check!]")
		}
	}
}

func Spotify(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var endpoint string = "https://spclient.wg.spotify.com/signup/public/v1/account"

	data := url.Values{}
	data.Set("validate", "1")
	data.Set("email", email)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response spotifyResponse
		json.Unmarshal(body, &response)
		if response.Status == 20 {
			social_result = append(social_result, "Spotify \U0001f440")
		} else {
			social_result = append(social_result, "Spotify [Not here!]")
		}
	} else {
		social_result = append(social_result, "Spotify [Couldn't check!]")
	}
}

func Twitter(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	var endpoint string = "https://api.twitter.com/i/users/email_available.json"

	data := url.Values{}
	data.Set("email", email)

	r, err := http.Get(endpoint + "?" + data.Encode())
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode == 200 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response twitterResponse
		json.Unmarshal(body, &response)
		if response.Taken {
			social_result = append(social_result, "Twitter \U0001f440")
		} else {
			social_result = append(social_result, "Twitter [Not here!]")
		}
	} else {
		social_result = append(social_result, "Twitter [Couldn't check!]")
	}
}
