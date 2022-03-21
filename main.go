// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"main/modules"
	"strconv"

	"os"

	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

func init() {
	templ := `{{ .Title "mosint" "" 2 }}
   {{ .AnsiColor.BrightWhite }}v2.1{{ .AnsiColor.Default }}
   {{ .AnsiColor.BrightCyan }}https://github.com/alpkeskin/{{ .AnsiColor.Default }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	println()
}

func help_menu() {
	data := [][]string{
		{"-e", "Set target email", "Yes"},
		{"-verify", "Verify the target email", "No"},
		{"-social", "Social scan for target email", "No"},
		{"-relateds", "Find related emails and domains with target email", "No"},
		{"-leaks", "Find password leaks for target email", "No"},
		{"-dumps", "Search pastebin dumps for target email", "No"},
		{"-domain", "More information about target email's domain", "No"},
		{"-o", "Output to text file", "No"},
		{"-v", "Version of mosint", "No"},
		{"-h", "Help Menu", "No"},
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

func verifyPrint(verifyData modules.VerifyStruct, emailRepData modules.EmailRepStruct, email string) {
	if verifyData.IsVerified {
		fmt.Println(email+" =>", color.GreenString("Verified \u2714"))
		outputText += email + " => Verified \n"
	} else {
		fmt.Println(email+" =>", color.RedString("Not Verified \u2718"))
		outputText += email + " => Not Verified \n"
	}
	if verifyData.IsDisposable {
		fmt.Println(email+" =>", color.RedString("Disposable \u2718"))
		outputText += email + " => Disposable \n"
	} else {
		fmt.Println(email+" =>", color.GreenString("Not Disposable \u2714"))
		outputText += email + " => Not Disposable \n"
	}

	if modules.GetAPIKey("EmailRep.io API Key") != "" {
		fmt.Println("\nEmailRep Data for", color.WhiteString(email))
		outputText += "\nEmailRep Data for " + email + "\n"
		fmt.Println("|- Reputation:", color.YellowString(emailRepData.Reputation))
		outputText += "|- Reputation: " + emailRepData.Reputation + "\n"
		fmt.Println("|- Blacklisted:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Blacklisted)))
		outputText += "|- Blacklisted: " + strconv.FormatBool(emailRepData.Details.Blacklisted) + "\n"
		fmt.Println("|- Malicious Activity:", color.WhiteString(strconv.FormatBool(emailRepData.Details.MaliciousActivity)))
		outputText += "|- Malicious Activity: " + strconv.FormatBool(emailRepData.Details.MaliciousActivity) + "\n"
		fmt.Println("|- Credential Leaked:", color.WhiteString(strconv.FormatBool(emailRepData.Details.CredentialsLeaked)))
		outputText += "|- Credential Leaked: " + strconv.FormatBool(emailRepData.Details.CredentialsLeaked) + "\n"
		fmt.Println("|- First Seen:", color.YellowString(emailRepData.Details.FirstSeen))
		outputText += "|- First Seen: " + emailRepData.Details.FirstSeen + "\n"
		fmt.Println("|- Last Seen:", color.YellowString(emailRepData.Details.LastSeen))
		outputText += "|- Last Seen: " + emailRepData.Details.LastSeen + "\n"
		fmt.Println("|- Day Since Domain Creation:", color.WhiteString(strconv.Itoa(emailRepData.Details.DaysSinceDomainCreation)))
		outputText += "|- Day Since Domain Creation: " + strconv.Itoa(emailRepData.Details.DaysSinceDomainCreation) + "\n"
		fmt.Println("|- Spam:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Spam)))
		outputText += "|- Spam: " + strconv.FormatBool(emailRepData.Details.Spam) + "\n"
		fmt.Println("|- Free Provider:", color.WhiteString(strconv.FormatBool(emailRepData.Details.FreeProvider)))
		outputText += "|- Free Provider: " + strconv.FormatBool(emailRepData.Details.FreeProvider) + "\n"
		fmt.Println("|- Deliverable:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Deliverable)))
		outputText += "|- Deliverable: " + strconv.FormatBool(emailRepData.Details.Deliverable) + "\n"
		fmt.Println("|- Valid MX:", color.WhiteString(strconv.FormatBool(emailRepData.Details.ValidMx)))
		outputText += "|- Valid MX: " + strconv.FormatBool(emailRepData.Details.ValidMx) + "\n"
	} else {
		color.Red("EmailRep.io API Key is not set!")
		outputText += "EmailRep.io API Key is not set!\n"
	}
}

func socialPrint(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	outputText += "Social Media Search Results: \n"
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("|- "+scanner.Text(), color.GreenString("\u2714"))
		outputText += "|- " + scanner.Text() + "\n"

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}

func relatedPrint(relEmails []string, relDomains []string, fromGoogle []string, hunterData modules.HunterStruct) {

	fmt.Println("Related Emails:")
	outputText += "Related Emails: \n"
	for _, v := range relEmails {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
	if modules.GetAPIKey("Hunter.io API Key") != "" {
		for _, v := range hunterData.Data.Emails {
			fmt.Println("|- "+v.Value, color.GreenString("\u2714"))
			outputText += "|- " + v.Value + "\n"
		}
	}
	println("")
	fmt.Println("Related Domains:")
	outputText += "Related Domains: \n"
	for _, v := range relDomains {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
	for _, v := range fromGoogle {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
}

func leakPrint(breachData modules.BreachDirectoryStruct, intelxData []string) {
	fmt.Println("Password Leaks:")
	outputText += "Password Leaks: \n"
	if breachData.Success {
		for _, v := range breachData.Result {
			for _, w := range v.Sources {
				fmt.Println("|- " + w)
				outputText += "|- " + w + "\n"
			}
			if v.HasPassword {
				fmt.Println("|-- "+v.Password, color.GreenString("\u2714"))
				outputText += "|-- " + v.Password + "\n"
			}
		}
	} else {
		fmt.Println("|- No Password Leaks from Breach Directory")
		outputText += "|- No Password Leaks from Breach Directory \n"
	}
	println("\nPassword Leaks from Intelx:")
	if len(intelxData) > 0 {
		for _, v := range intelxData {
			fmt.Println("|- "+v, color.GreenString("\u2714"))
			outputText += "|- " + v + "\n"
		}
	} else {
		color.Red("No intelx file!")
		outputText += "No intelx file! \n"
	}
}

func dumpPrint(binData []string) {
	fmt.Println("Pastebin and Throwbin Search Results:")
	outputText += "Pastebin and Throwbin Search Results: \n"
	for _, v := range binData {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
}

func domainPrint(table *tablewriter.Table, ipapi modules.IPAPIStruct) {
	println("\nDomain Information:")
	outputText += "Domain Information: \n"
	fmt.Println("|- IP: "+ipapi.IP, color.GreenString("\u2714"))
	outputText += "|- IP: " + ipapi.IP + "\n"
	fmt.Println("|- City: "+ipapi.City, color.GreenString("\u2714"))
	outputText += "|- City: " + ipapi.City + "\n"
	fmt.Println("|- Region: "+ipapi.Region, color.GreenString("\u2714"))
	outputText += "|- Region: " + ipapi.Region + "\n"
	fmt.Println("|- Region Code: "+ipapi.RegionCode, color.GreenString("\u2714"))
	outputText += "|- Region Code: " + ipapi.RegionCode + "\n"
	fmt.Println("|- Country: "+ipapi.Country, color.GreenString("\u2714"))
	outputText += "|- Country: " + ipapi.Country + "\n"
	fmt.Println("|- Country Code: "+ipapi.CountryCode, color.GreenString("\u2714"))
	outputText += "|- Country Code: " + ipapi.CountryCode + "\n"
	fmt.Println("|- Country Name: "+ipapi.CountryName, color.GreenString("\u2714"))
	outputText += "|- Country Name: " + ipapi.CountryName + "\n"
	fmt.Println("|- Postal: "+ipapi.Postal, color.GreenString("\u2714"))
	outputText += "|- Postal: " + ipapi.Postal + "\n"
	fmt.Println("|- TimeZone: "+ipapi.Timezone, color.GreenString("\u2714"))
	outputText += "|- TimeZone: " + ipapi.Timezone + "\n"
	fmt.Println("|- Country Calling Code: "+ipapi.CountryCallingCode, color.GreenString("\u2714"))
	outputText += "|- Country Calling Code: " + ipapi.CountryCallingCode + "\n"
	fmt.Println("|- Currency: "+ipapi.Currency, color.GreenString("\u2714"))
	outputText += "|- Currency: " + ipapi.Currency + "\n"
	fmt.Println("|- Organization: "+ipapi.Org, color.GreenString("\u2714"))
	outputText += "|- Organization: " + ipapi.Org + "\n"

	println("\nDNS Records:")
	table.Render()
}

var outputText string = ""

func main() {
	var email *string = flag.String("e", "", "Set email")
	var verify *bool = flag.Bool("verify", false, "Verify method")
	var social_accounts *bool = flag.Bool("social", false, "Finding registered accounts from email")
	var relateds *bool = flag.Bool("relateds", false, "Finding related emails and domains from domain")
	var leaks *bool = flag.Bool("leaks", false, "Finding password leaks from email")
	var dumps *bool = flag.Bool("dumps", false, "Finding Pastebin dumps from email")
	var domain *bool = flag.Bool("domain", false, "More information about domain")
	var output *bool = flag.Bool("o", false, "Output to text file")
	var version *bool = flag.Bool("v", false, "Version of mosint")
	var help *bool = flag.Bool("h", false, "Help Menu")
	var all *bool = flag.Bool("all", false, "All features!")
	flag.Parse()
	println("")
	if len(*email) == 0 {
		help_menu()
		os.Exit(0)
	} else if *help {
		help_menu()
		os.Exit(0)
	} else if *version {
		color.White("version: 2.1")
		os.Exit(0)
	} else if *all {

		var bar = progressbar.Default(100, "mosinting")
		bar.Add(5)
		var verifyData = modules.VerifyEmail(*email)
		bar.Add(7)
		var emailRepData = modules.EmailRep(*email)
		bar.Add(8)
		modules.Runner(*email, "SocialScan")
		bar.Add(8)
		modules.Runner(*email, "Holehe")
		var relEmails = modules.RelatedEmails(*email)
		bar.Add(8)
		var relDomains = modules.RelatedDomains(*email)
		bar.Add(8)
		var fromGoogle = modules.Related_domains_from_google(*email)
		bar.Add(8)
		var hunterData = modules.Hunter(*email)
		bar.Add(8)
		var breachData = modules.BreachDirectory(*email)
		bar.Add(8)
		var binData = modules.BinSearch(*email)
		bar.Add(8)
		var ipapi = modules.IPAPI(*email)
		bar.Add(8)
		var table = modules.DNS_lookup(*email)
		bar.Add(8)
		var intelxData = modules.Intelx(*email)
		bar.Finish()
		verifyPrint(verifyData, emailRepData, *email)
		socialPrint("socialscantempresult.txt")
		socialPrint("holehetempresult.txt")
		relatedPrint(relEmails, relDomains, fromGoogle, hunterData)
		leakPrint(breachData, intelxData)
		dumpPrint(binData)
		domainPrint(table, ipapi)
		if *output {
			var filename = modules.FileWriter(*email, outputText)
			color.Green("\nOutput file: " + filename)
		}
		os.Exit(0)
	} else {
		if *verify {
			verifyPrint(modules.VerifyEmail(*email), modules.EmailRep(*email), *email)
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
			os.Exit(0)
		}
		if *social_accounts {
			fmt.Println("Social media accounts opened with", color.WhiteString(*email))
			modules.Runner(*email, "SocialScan")
			modules.Runner(*email, "Holehe")
			socialPrint("socialscantempresult.txt")
			socialPrint("holehetempresult.txt")
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
			os.Exit(0)
		}
		if *relateds {
			relatedPrint(modules.RelatedEmails(*email), modules.RelatedDomains(*email), modules.Related_domains_from_google(*email), modules.Hunter(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
			os.Exit(0)
		}
		if *leaks {
			leakPrint(modules.BreachDirectory(*email), modules.Intelx(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
			os.Exit(0)
		}
		if *dumps {
			dumpPrint(modules.BinSearch(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
		}
		if *domain {
			domainPrint(modules.DNS_lookup(*email), modules.IPAPI(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nOutput file: " + filename)
			}
			os.Exit(0)
		}
	}

}
