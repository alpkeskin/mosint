// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package utils

import (
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/alpkeskin/mosint/cmd/models"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

var TitleMap = map[string]string{
	"EmailRep":        "\nEmailRep Results:",
	"Hunter":          "\nHunter Results:",
	"Googling":        "\nGoogling Results:",
	"Intelx":          "\nIntelX Results:",
	"Psbdmp":          "\nPastebin Dumps (psbdmp):",
	"Social":          "\nSocial Media Results:",
	"IPAPI":           "\nIPAPI Results:",
	"Lookup":          "\nLookup Results:",
	"BreachDirectory": "\nBreachDirectory Results:",
}

var (
	ConfigReturn   map[string]interface{}
	ProgressBar    *progressbar.ProgressBar = nil
	LookupTable    *tablewriter.Table       = tablewriter.NewWriter(os.Stdout)
	BannerTemplate string                   = `{{ .Title "mosint" "" 2 }}
{{ .AnsiColor.BrightWhite }}v2.3{{ .AnsiColor.Default }}
{{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
Now: {{ .Now "Monday, 2 Jan 2006" }}`
)

var (
	Verify_result          models.VerifyStruct
	Emailrep_result        models.EmailRepStruct
	Breachdirectory_result models.BreachDirectoryStruct
	Hunter_result          models.HunterStruct
	Googling_result        []string
	Intelx_result          []string
	Social_result          []string
	Psbdmp_result          models.PsbdmpStruct
	Lookup_temp_result     [][]string
	Ipapi_result           models.IPAPIStruct
	Verifier               = emailverifier.NewVerifier().
				EnableAutoUpdateDisposable()
)

const (
	IntelxDefaultMaxResults = 10
)

const (
	BreachDirectoryAPI    = "https://breachdirectory.p.rapidapi.com/?func=auto&term="
	EmailrepURL           = "https://emailrep.io/"
	HunterAPI             = "https://api.hunter.io/v2/domain-search?domain="
	IntelxURL             = "https://intelx.io/"
	IPAPIURL              = "https://ipapi.co/"
	PsbdmpAPI             = "https://psbdmp.ws/api/v3/search/"
	AdobeEndpoint         = "https://auth.services.adobe.com/signin/v2/users/accounts"
	DiscordEndpoint       = "https://discord.com/api/v9/auth/register"
	InstagramCSRFEndpoint = "https://www.instagram.com/accounts/emailsignup/"
	InstagramEndpoint     = "https://www.instagram.com/accounts/web_create_ajax/attempt/"
	SpotifyEndpoint       = "https://spclient.wg.spotify.com/signup/public/v1/account"
	TwitterEndpoint       = "https://api.twitter.com/i/users/email_available.json"
)
