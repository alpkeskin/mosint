// https://github.com/alpkeskin/
package main

import (
	"flag"
	"fmt"
	"time"

	"main/modules"
	"os"

	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/olekukonko/tablewriter"
)

func init() {
	templ := `{{ .Title "mosint" "" 2 }}
   {{ .AnsiColor.BrightWhite }}v2.0{{ .AnsiColor.Default }}
   {{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	println()
}

func help_menu() {
	data := [][]string{
		{"-e", "Set target email", "Yes"},
		{"-v", "Verify the target email", "No"},
		{"-ss", "Social scan for target email", "No"},
		{"-re", "Find related emails with target email", "No"},
		{"-rd", "Find related domains with target email", "No"},
		{"-l", "Find password leaks for target email", "No"},
		{"-pd", "Search pastebin dumps for target email", "No"},
		{"-er", "EmailRep.io API", "No"},
		{"-d", "More information about target email's domain", "No"},
		{"-all", "All features!", "No"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flags", "Description", "isRequired"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	color.Yellow("Example: go run main.go -e example@domain.com -all")
}

func main() {
	var email *string = flag.String("e", "", "Set email")
	var verify *bool = flag.Bool("v", false, "Verify method")
	var social_scan *bool = flag.Bool("ss", false, "Social scan method")
	var related_emails *bool = flag.Bool("re", false, "Related emails method")
	var related_domains *bool = flag.Bool("rd", false, "Related domains method")
	var leaks *bool = flag.Bool("l", false, "Find password leaks method")
	var pastebin_dumps *bool = flag.Bool("pd", false, "Pastebin dumps searching method")
	var emailrep *bool = flag.Bool("er", false, "EmailRep.io API")
	var domain *bool = flag.Bool("d", false, "More information about domain method")
	var help *bool = flag.Bool("h", false, "Help Menu")
	var all *bool = flag.Bool("all", false, "All features!")
	flag.Parse()
	if *email == "" {
		help_menu()
	} else if *help {
		help_menu()
	} else {
		whilte := color.New(color.FgWhite)
		boldWhite := whilte.Add(color.Bold)
		fmt.Print("\nEmail > ")
		boldWhite.Print(*email + "\n")

		if *all {
			start := time.Now().Unix()
			modules.Verify_email(*email)
			color.Magenta("Which Social Media Does " + *email + " Use?")
			modules.Runner(*email, "socialscan")
			color.Magenta("Related Emails:")
			modules.Related_emails(*email)
			modules.Runner(*email, "hunter")
			color.Magenta("Related Domains:")
			modules.Related_domains(*email)
			modules.Related_domains_google(*email)
			color.Magenta("Password Leaks:")
			modules.Runner(*email, "breachdirectory")
			color.Magenta("Pastebin Dumps Searching:")
			modules.Pastebin_search(*email)
			color.Magenta("EmailRep.io API Results:")
			modules.Runner(*email, "emailrep")
			color.Magenta("Domain Investigation:")
			modules.DNS_lookup(*email)
			modules.IPapi(*email)
			modules.Subdomains(*email)
			end := time.Now().Unix()
			println("---------------------------------------")
			print("Scan duration: ")
			boldWhite.Print(fmt.Sprint(end - start))
			print(" seconds.")
		} else {
			if *verify {
				modules.Verify_email(*email)
			}
			if *social_scan {
				color.Magenta("Which Social Media Does " + *email + " Use?")
				modules.Runner(*email, "socialscan")
			}
			if *related_emails {
				color.Magenta("Related Emails:")
				modules.Related_emails(*email)
				modules.Runner(*email, "hunter")
			}
			if *related_domains {
				color.Magenta("Related Domains:")
				modules.Related_domains(*email)
				modules.Related_domains_google(*email)
			}
			if *leaks {
				color.Magenta("Password Leaks:")
				modules.Runner(*email, "breachdirectory")
			}
			if *pastebin_dumps {
				color.Magenta("Pastebin Dumps Searching:")
				modules.Pastebin_search(*email)
			}
			if *emailrep {
				color.Magenta("EmailRep.io API Results:")
				modules.Runner(*email, "emailrep")
			}
			if *domain {
				color.Magenta("Domain Investigation:")
				modules.DNS_lookup(*email)
				modules.IPapi(*email)
				modules.Subdomains(*email)
			}
		}
	}
}
