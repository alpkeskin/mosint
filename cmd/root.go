// mosint v2.3
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"os"

	"github.com/alpkeskin/mosint/cmd/utils"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "mosint [email]",
	Short:        "\nAn automated e-mail OSINT tool",
	Long:         "\nAn automated e-mail OSINT tool",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		Do(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if os.Args[0] != "set" {
		banner.InitString(colorable.NewColorableStdout(), true, true, utils.BannerTemplate)
	}
	println()
	utils.ProgressBar = progressbar.Default(100)
}
