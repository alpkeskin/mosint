/*
Copyright © 2023 github.com/alpkeskin
*/
package engine

import (
	"bufio"
	"fmt"
	"os"

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
	Use:     "mosint [email|emails.txt]",
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
	concurrencyLimit int
)

func Start() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func magic(cmd *cobra.Command, args []string) {
    if coffee {
        printCoffee()
        os.Exit(0)
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
	if len(args) == 0 {
		cmd.Help()
		return
	}
    if fileExists(args[0]) {
        processEmailsFromFile(args[0])
    } else {
        processSingleEmail(args[0])
    }
}

func processSingleEmail(email string) {
    fmt.Println("Target Email:", color.YellowString(email), color.GreenString("✓"))

    if !verification.New().Syntax(email) {
        color.Red("Email syntax is not valid! Process stopped.")
        os.Exit(1)
    }

    processAndOutput(email)
}

func processEmailsFromFile(filePath string) {
    emails := []string{}
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        email := scanner.Text()
        if email != "" {
            emails = append(emails, email)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
        os.Exit(1)
    }
    if len(emails) > 0 {
        concurrencyLimit := 10
        multiRunner := runner.NewMultiRunner(emails, concurrencyLimit)
        multiRunner.Start()
        if outputPath != "" {
            output := output.New()
            output.SetFilePath(outputPath)
            output.JSONMulti(multiRunner)
        }
    }
}


func processAndOutput(email string) {
    runner := runner.New(email)
    runner.Start()

    if !silent {
        runner.Print()
    }

    if outputPath != "" {
        output := output.New()
        output.SetFilePath(outputPath)
        output.JSON(runner)
    }
}

func printCoffee() {
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
}

func configHelp() string {
	return `
	 You must create a config file in your home directory named %s
	 Or you can use --config flag to specify a config file

	 You can find an example config file in the repository:
	 https://github.com/alpkeskin/mosint/blob/master/example-config.yaml
	`
}

func fileExists(filePath string) bool {
    info, err := os.Stat(filePath)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func init() {
	template := `{{ .Title "mosint" "" 2 }}
	{{ .AnsiColor.BrightWhite }}v3.0{{ .AnsiColor.Default }}
	{{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, template)
	println()
	rootCmd.PersistentFlags().IntVarP(&concurrencyLimit, "concurrency", "C", 10, "concurrency limit for processing emails from file")
	rootCmd.PersistentFlags().BoolVar(&coffee, "coffee", false, "☕️")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mosint.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "output file (.json)")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "silent mode (only output file)")
}
