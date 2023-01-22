// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package models

type ConfigStruct struct {
	Breachdirectory string
	Hunter          string
	Emailrep        string
	Intelx          string
	Psbdmp          string
}

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

type EmailRepStruct struct {
	Email      string `json:"email"`
	Reputation string `json:"reputation"`
	Suspicious bool   `json:"suspicious"`
	References int    `json:"references"`
	Details    struct {
		Blacklisted             bool     `json:"blacklisted"`
		MaliciousActivity       bool     `json:"malicious_activity"`
		MaliciousActivityRecent bool     `json:"malicious_activity_recent"`
		CredentialsLeaked       bool     `json:"credentials_leaked"`
		CredentialsLeakedRecent bool     `json:"credentials_leaked_recent"`
		DataBreach              bool     `json:"data_breach"`
		FirstSeen               string   `json:"first_seen"`
		LastSeen                string   `json:"last_seen"`
		DomainExists            bool     `json:"domain_exists"`
		DomainReputation        string   `json:"domain_reputation"`
		NewDomain               bool     `json:"new_domain"`
		DaysSinceDomainCreation int      `json:"days_since_domain_creation"`
		SuspiciousTld           bool     `json:"suspicious_tld"`
		Spam                    bool     `json:"spam"`
		FreeProvider            bool     `json:"free_provider"`
		Disposable              bool     `json:"disposable"`
		Deliverable             bool     `json:"deliverable"`
		AcceptAll               bool     `json:"accept_all"`
		ValidMx                 bool     `json:"valid_mx"`
		PrimaryMx               string   `json:"primary_mx"`
		Spoofable               bool     `json:"spoofable"`
		SpfStrict               bool     `json:"spf_strict"`
		DmarcEnforced           bool     `json:"dmarc_enforced"`
		Profiles                []string `json:"profiles"`
	} `json:"details"`
}

type HunterStruct struct {
	Data struct {
		Domain       string      `json:"domain"`
		Disposable   bool        `json:"disposable"`
		Webmail      bool        `json:"webmail"`
		AcceptAll    bool        `json:"accept_all"`
		Pattern      string      `json:"pattern"`
		Organization string      `json:"organization"`
		Country      string      `json:"country"`
		State        interface{} `json:"state"`
		Emails       []struct {
			Value      string `json:"value"`
			Type       string `json:"type"`
			Confidence int    `json:"confidence"`
			Sources    []struct {
				Domain      string `json:"domain"`
				URI         string `json:"uri"`
				ExtractedOn string `json:"extracted_on"`
				LastSeenOn  string `json:"last_seen_on"`
				StillOnPage bool   `json:"still_on_page"`
			} `json:"sources"`
			FirstName    string      `json:"first_name"`
			LastName     string      `json:"last_name"`
			Position     string      `json:"position"`
			Seniority    string      `json:"seniority"`
			Department   string      `json:"department"`
			Linkedin     interface{} `json:"linkedin"`
			Twitter      interface{} `json:"twitter"`
			PhoneNumber  interface{} `json:"phone_number"`
			Verification struct {
				Date   string `json:"date"`
				Status string `json:"status"`
			} `json:"verification"`
		} `json:"emails"`
		LinkedDomains []interface{} `json:"linked_domains"`
	} `json:"data"`
	Meta struct {
		Results int `json:"results"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Params  struct {
			Domain     string      `json:"domain"`
			Company    interface{} `json:"company"`
			Type       interface{} `json:"type"`
			Seniority  interface{} `json:"seniority"`
			Department interface{} `json:"department"`
		} `json:"params"`
	} `json:"meta"`
}

type IPAPIStruct struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

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

type AdobeResponse []struct {
	HasT2ELinked bool `json:"hasT2ELinked"`
}

type DiscordResponse struct {
	Errors struct {
		Email struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"email"`
	} `json:"errors"`
}

type SpotifyResponse struct {
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

type TwitterResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

type VerifyStruct struct {
	IsVerified   bool  `json:"is_verified"`
	IsDisposable bool  `json:"is_disposable"`
	Err          error `json:"err"`
}
