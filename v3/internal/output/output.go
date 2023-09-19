/*
Copyright © 2023 github.com/alpkeskin
*/
package output

import (
	"encoding/json"
	"os"

	"github.com/alpkeskin/mosint/v3/internal/runner"
	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/alpkeskin/mosint/v3/pkg/dns"
	"github.com/alpkeskin/mosint/v3/pkg/services/breachdirectory"
	"github.com/alpkeskin/mosint/v3/pkg/services/emailrep"
	"github.com/alpkeskin/mosint/v3/pkg/services/haveibeenpwned"
	"github.com/alpkeskin/mosint/v3/pkg/services/hunter"
	"github.com/alpkeskin/mosint/v3/pkg/services/ipapi"
)

type Output struct {
	FilePath string
}

func New() *Output {
	return &Output{}
}

type jsonOutput struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`

	EmailRep        emailrep.EmailRepResponse               `json:"emailrep"`
	BreachDirectory breachdirectory.BreachDirectoryResponse `json:"breachdirectory"`
	HaveIBeenPwned  haveibeenpwned.HaveIBeenPwnedResponse   `json:"haveibeenpwned"`
	Hunter          hunter.HunterResponse                   `json:"hunter"`
	IpApi           ipapi.IpApiResponse                     `json:"ipapi"`
	IntelX          []string                                `json:"intelx"`
	PsbDmp          []string                                `json:"psbdmp"`
	InstagramExists bool                                    `json:"instagram_exists"`
	SpotifyExists   bool                                    `json:"spotify_exists"`
	TwitterExists   bool                                    `json:"twitter_exists"`
	GoogleSearch    []string                                `json:"google_search"`
	DnsRecords      []dns.Record                            `json:"dns_records"`
}

func (o *Output) SetFilePath(filePath string) {
	o.FilePath = filePath
}

func (o *Output) JSON(runner *runner.Runner) {
	spinner := spinner.New("JSON Output")
	spinner.Start()

	data := jsonOutput{
		Email:           runner.Email,
		Verified:        true,
		EmailRep:        runner.EmailRepC.Response,
		BreachDirectory: runner.BreachDirectoryC.Response,
		HaveIBeenPwned:  runner.HaveibeenpwnedC.Response,
		Hunter:          runner.HunterC.Response,
		IpApi:           runner.IpApiC.Response,
		IntelX:          runner.IntelxC.Response,
		PsbDmp:          runner.PsbdmpC.Urls,
		InstagramExists: runner.InstagramC.Exists,
		SpotifyExists:   runner.SpotifyC.Exists,
		TwitterExists:   runner.TwitterC.Exists,
		GoogleSearch:    runner.GoogleSearchC.Response,
		DnsRecords:      runner.DnsC.Records,
	}

	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	err = os.WriteFile(o.FilePath, file, 0644)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	spinner.Stop()
}
