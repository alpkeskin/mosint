/*
Copyright © 2023 github.com/alpkeskin
*/
package engine

import (
	"fmt"
	"os"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/config"
	. "github.com/alpkeskin/mosint/v3/internal/config"
	"github.com/alpkeskin/mosint/v3/internal/output"
	"github.com/alpkeskin/mosint/v3/internal/runner"
	"github.com/alpkeskin/mosint/v3/pkg/verification"
	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mosint [email]",
	Short:   "\nAn automated e-mail OSINT tool ",
	Long:    "\nAn automated e-mail OSINT tool written in Go with a focus on simplicity and performance.",
	Run:     magic,
	Version: "v3.0.0",
}

var (
	cfgFile    string
	outputPath string
	silent     bool
	coffee     bool
)

func Start() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func magic(cmd *cobra.Command, args []string) {
	if coffee {
		c := `
		( (
			) )
		  ........
		  |      |]
		  \      / 
		   '----'

		enjoy your coffee ☕️
		`
		fmt.Println(c)
		os.Exit(0)
	}

	if len(args) < 1 {
		cmd.Help()
		return
	}

	Cfg = config.New()
	if !Cfg.Exists(cfgFile) {
		msg := fmt.Sprintf(configHelp(), color.YellowString(".mosint.yaml"))
		println(msg)
		os.Exit(1)
	}

	err := Cfg.Parse(cfgFile)

	if err != nil {
		panic(err)
	}

	email := args[0]
	fmt.Println("Target Email:", color.YellowString(email), color.GreenString("✓"))
	println()

	verification := verification.New()
	if !verification.Syntax(email) {
		color.Red("Email syntax is not valid! Process stopped.")
		os.Exit(1)
	}

	runner := runner.New(email)
	runner.Start()

	if !silent {
		runner.Print()
	}

	if !strings.EqualFold(outputPath, "") {
		output := output.New()
		output.SetFilePath(outputPath)
		output.JSON(runner)
	}
}

func configHelp() string {
	return `
	 You must create a config file in your home directory named %s
	 Or you can use --config flag to specify a config file

	 You can find an example config file in the repository:
	 https://github.com/alpkeskin/mosint/blob/master/example-config.yaml
	`
}

func init() {
	template := `{{ .Title "mosint" "" 2 }}
	{{ .AnsiColor.BrightWhite }}v3.0{{ .AnsiColor.Default }}
	{{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, template)
	println()

	rootCmd.PersistentFlags().BoolVar(&coffee, "coffee", false, "☕️")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mosint.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "output file (.json)")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "silent mode (only output file)")
}
