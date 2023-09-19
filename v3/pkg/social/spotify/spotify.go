/*
Copyright © 2023 github.com/alpkeskin
*/
package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/fatih/color"
)

type Spotify struct {
	Exists bool
}

type response struct {
	Status int `json:"status"`
	Errors struct {
		Email string `json:"email"`
	} `json:"errors"`
	Country                         string `json:"country"`
	CanAcceptLicensesInOneStep      bool   `json:"can_accept_licenses_in_one_step"`
	RequiresMarketingOptIn          bool   `json:"requires_marketing_opt_in"`
	RequiresMarketingOptInText      bool   `json:"requires_marketing_opt_in_text"`
	MinimumAge                      int    `json:"minimum_age"`
	CountryGroup                    string `json:"country_group"`
	SpecificLicenses                bool   `json:"specific_licenses"`
	TermsConditionsAcceptance       string `json:"terms_conditions_acceptance"`
	PrivacyPolicyAcceptance         string `json:"privacy_policy_acceptance"`
	SpotifyMarketingMessagesOption  string `json:"spotify_marketing_messages_option"`
	PretickEula                     bool   `json:"pretick_eula"`
	ShowCollectPersonalInfo         bool   `json:"show_collect_personal_info"`
	UseAllGenders                   bool   `json:"use_all_genders"`
	UseOtherGender                  bool   `json:"use_other_gender"`
	UsePreferNotToSayGender         bool   `json:"use_prefer_not_to_say_gender"`
	ShowNonRequiredFieldsAsOptional bool   `json:"show_non_required_fields_as_optional"`
	DateEndianness                  int    `json:"date_endianness"`
	IsCountryLaunched               bool   `json:"is_country_launched"`
	AllowedCallingCodes             []struct {
		CountryCode string `json:"country_code"`
		CallingCode string `json:"calling_code"`
	} `json:"allowed_calling_codes"`
	PushNotifications bool `json:"push-notifications"`
}

func New() *Spotify {
	return &Spotify{}
}

func (s *Spotify) Check(email string) {
	spinner := spinner.New("Spotify Account Checking")
	spinner.Start()

	url := fmt.Sprintf("https://spclient.wg.spotify.com/signup/public/v1/account?validate=1&email=%s", email)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 10; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.120 Mobile Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	var r response
	err = json.Unmarshal(body, &r)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	if r.Status != 20 {
		s.Exists = false
		spinner.Stop()
		return
	}

	s.Exists = true
	spinner.Stop()
}

func (i *Spotify) Print() {
	if i.Exists {
		fmt.Println(color.GreenString("[+]"), "Spotify Account Exists", color.GreenString("\u2714"))
		return
	}
	fmt.Println(color.RedString("[!]"), "Spotify Account Not Exists", color.RedString("\u2718"))
}
