// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"os"

	"github.com/alpkeskin/mosint/cmd/controllers"
	"github.com/alpkeskin/mosint/cmd/utils"
	"github.com/fatih/color"
	"github.com/gammazero/workerpool"
	"github.com/spf13/cobra"
)

var Controllers = []func(string){
	controllers.VerifyEmail,
	controllers.EmailRep,
	controllers.Hunter,
	controllers.Googling,
	controllers.Adobe,
	controllers.Discord,
	controllers.Instagram,
	controllers.Spotify,
	controllers.Twitter,
	controllers.Psbdmp,
	controllers.Intelx,
	controllers.BreachDirectory,
	controllers.IPAPI,
	controllers.Lookup,
}

func Do(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
	if args[0] == "set" {
		utils.SetAPIKey(args[1], args[2])
		os.Exit(0)
	}
	utils.EmailRegex(args[0])
	wp := workerpool.New(14)
	for _, controller := range Controllers {
		controller := controller
		wp.Submit(func() {
			controller(args[0])
		})
	}
	wp.StopWait()

	println()
	whiteBackground := color.New(color.FgRed).Add(color.BgWhite)
	println("Target Email:", whiteBackground.Sprint(args[0]))
	utils.PrintVerify(utils.Verify_result)
	utils.PrintEmailRep(utils.Emailrep_result)
	utils.PrintHunter(utils.Hunter_result)
	utils.PrintSocial(utils.Social_result)
	utils.PrintPsbdmp(utils.Psbdmp_result)
	utils.PrintIntelx(utils.Intelx_result)
	utils.PrintBreachDirectory(utils.Breachdirectory_result)
	utils.PrintIPAPI(utils.Ipapi_result)
	utils.PrintLookup(utils.Lookup_temp_result)
}
