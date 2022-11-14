// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

var verify_result VerifyStruct
var emailrep_result EmailRepStruct
var breachdirectory_result BreachDirectoryStruct
var hunter_result HunterStruct
var googling_result []string
var intelx_result []string
var social_result []string
var psbdmp_result PsbdmpStruct
var lookup_temp_result [][]string
var ipapi_result IPAPIStruct

var rootCmd = &cobra.Command{
	Use:          "mosint [email]",
	Short:        "\nAn automated e-mail OSINT tool",
	Long:         "\nAn automated e-mail OSINT tool",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		if args[0] == "set" {
			SetAPIKey(args[1], args[2])
			os.Exit(0)
		}
		EmailRegex(args[0])
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		var wg sync.WaitGroup
		wg.Add(13)
		go VerifyEmail(&wg, args[0])
		go EmailRep(&wg, args[0])
		go Hunter(&wg, args[0])
		go Googling(&wg, args[0])
		go Adobe(&wg, args[0])
		go Discord(&wg, args[0])
		go Instagram(&wg, args[0])
		go Spotify(&wg, args[0])
		go Twitter(&wg, args[0])
		go Psbdmp(&wg, args[0])
		go Intelx(&wg, args[0])
		go BreachDirectory(&wg, args[0])
		go Lookup(&wg, args[0])

		wg.Wait()
		s.Stop()
		println()
		whiteBackground := color.New(color.FgRed).Add(color.BgWhite)
		println("Target Email:", whiteBackground.Sprint(args[0]))
		PrintVerify(verify_result)
		PrintEmailRep(emailrep_result)
		PrintHunter(hunter_result)
		PrintSocial(social_result)
		PrintPsbdmp(psbdmp_result)
		PrintIntelx(intelx_result)
		PrintBreachDirectory(breachdirectory_result)
		PrintLookup(lookup_temp_result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	templ := `{{ .Title "mosint" "" 2 }}
	{{ .AnsiColor.BrightWhite }}v2.2{{ .AnsiColor.Default }}
	{{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}`
	if os.Args[0] != "set" {
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	}
	println()
	CreateConfig()
}
