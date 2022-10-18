// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type ConfigStruct struct {
	Breachdirectory string
	Hunter          string
	Emailrep        string
	Intelx          string
	Psbdmp          string
}

func SetAPIKey(key string, value string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting home directory")
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(dirname + "/mosint-config.json")
	if err != nil {
		panic(err)
	}
	var returnData map[string]interface{}
	err = json.Unmarshal(data, &returnData)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return
	}
	key = strings.Title(key)
	returnData[key] = value
	file, _ := json.MarshalIndent(returnData, "", " ")
	_ = ioutil.WriteFile(dirname+"/mosint-config.json", file, 0644)
	fmt.Println("[", color.BlueString("INFO"), "] API key set successfully!")

}

func CreateConfig() {

	dirname, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting home directory")
		os.Exit(0)
	}
	if _, err := os.Stat(dirname + "/mosint-config.json"); os.IsNotExist(err) {

		config := ConfigStruct{
			Breachdirectory: "",
			Hunter:          "",
			Emailrep:        "",
			Intelx:          "",
			Psbdmp:          "",
		}

		file, _ := json.MarshalIndent(config, "", " ")
		_ = ioutil.WriteFile(dirname+"/mosint-config.json", file, 0644)
		var example string
		fmt.Println("\n[", color.BlueString("INFO"), "] Config file created at "+dirname+"/mosint-config.json ")
		fmt.Println("[", color.BlueString("INFO"), "] If you want to use mosint with full features, set your API keys.")
		example = "mosint set hunter <hunter.io API key>\n mosint set emailrep <emailrep.io API key>\n mosint set intelx <intelx.io API key>\n mosint set psbdmp <psbdmp.ws API key>\n mosint set breachdirectory <breachdirectory.org API key>"
		fmt.Println("Examples: \n", example)

	}
}

func GetAPIKey(key string) string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting home directory")
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(dirname + "/mosint-config.json")
	if err != nil {
		panic(err)
	}
	var returnData map[string]interface{}
	err = json.Unmarshal(data, &returnData)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return ""
	}
	return returnData[key].(string)
}

func EmailRegex(email string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		color.Red("\nInvalid email address!")
		os.Exit(0)
	}
}

func PrintVerify(verify_result VerifyStruct) {
	if verify_result.IsVerified {
		fmt.Println("|-->", color.GreenString("Verified \u2714"))
	} else {
		fmt.Println("|-->", color.RedString("Not Verified \u2718"))
	}
	if verify_result.IsDisposable {
		fmt.Println("|-->", color.RedString("Disposable \u2718"))
	} else {
		fmt.Println("|-->", color.GreenString("Not Disposable \u2714"))
	}
}

func PrintEmailRep(emailrep_result EmailRepStruct) {
	if GetAPIKey("Emailrep") == "" {
		return
	}
	fmt.Println("\nEmailRep Results:")
	fmt.Println("|- Reputation:", color.YellowString(emailrep_result.Reputation))
	fmt.Println("|- Blacklisted:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.Blacklisted)))
	fmt.Println("|- Malicious Activity:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.MaliciousActivity)))
	fmt.Println("|- Credential Leaked:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.CredentialsLeaked)))
	fmt.Println("|- First Seen:", color.YellowString(emailrep_result.Details.FirstSeen))
	fmt.Println("|- Last Seen:", color.YellowString(emailrep_result.Details.LastSeen))
	fmt.Println("|- Day Since Domain Creation:", color.WhiteString(strconv.Itoa(emailrep_result.Details.DaysSinceDomainCreation)))
	fmt.Println("|- Spam:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.Spam)))
	fmt.Println("|- Free Provider:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.FreeProvider)))
	fmt.Println("|- Deliverable:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.Deliverable)))
	fmt.Println("|- Valid MX:", color.WhiteString(strconv.FormatBool(emailrep_result.Details.ValidMx)))
}

func PrintHunter(hunter_result HunterStruct) {
	if GetAPIKey("Hunter") == "" {
		return
	}
	fmt.Println("\nHunter Results:")
	fmt.Println("|- Disposable:", color.YellowString(strconv.FormatBool(hunter_result.Data.Disposable)))
	fmt.Println("|- Webmail:", color.YellowString(strconv.FormatBool(hunter_result.Data.Webmail)))
	fmt.Println("|- AcceptAll:", color.YellowString(strconv.FormatBool(hunter_result.Data.AcceptAll)))
	fmt.Println("|- Pattern:", color.WhiteString(hunter_result.Data.Pattern))
	fmt.Println("Related Emails:")
	if len(hunter_result.Data.Emails) == 0 {
		color.Red("|- No related emails found")
	} else {
		for _, v := range hunter_result.Data.Emails {
			fmt.Println("|- "+v.Value, color.GreenString("\u2714"))
		}
	}
}

func PrintGoogle(googling_result []string) {
	fmt.Println("\nGoogle Results:")
	if len(googling_result) == 0 {
		color.Red("|- No results found")
	} else {
		for _, v := range googling_result {
			fmt.Println("|- "+v, color.GreenString("\u2714"))
		}
	}
}

func PrintSocial(social_result []string) {
	fmt.Println("\nSocial Media Results:")
	for _, v := range social_result {
		fmt.Println("|- " + v)
	}
}

func PrintPsbdmp(psbdmp_result PsbdmpStruct) {
	if GetAPIKey("Psbdmp") == "" {
		return
	}
	fmt.Println("\nPastebin Dumps (psbdmp):")
	if len(psbdmp_result.Data) == 0 {
		color.Red("|- No results found")
	} else {
		for _, v := range psbdmp_result.Data {
			fmt.Println("|- pastebin.com/"+v.ID, color.GreenString("\u2714"))
		}
	}
}

func PrintIntelx(intelx_result []string) {
	if GetAPIKey("Intelx") == "" {
		return
	}
	fmt.Println("\nIntelX Results:")
	if len(intelx_result) == 0 {
		color.Red("|- No results found")
	} else {
		for _, v := range intelx_result {
			fmt.Println("|- "+v, color.GreenString("\u2714"))
		}
	}
}

func PrintBreachDirectory(breachdirectory_result BreachDirectoryStruct) {
	if GetAPIKey("Breachdirectory") == "" {
		return
	}
	fmt.Println("\nBreachDirectory Results:")
	if breachdirectory_result.Success {
		for _, v := range breachdirectory_result.Result {
			fmt.Println("|- Sources:", color.GreenString("\u2714"))
			for _, v2 := range v.Sources {
				fmt.Println("|-- "+v2, color.GreenString("\u2714"))
			}
			fmt.Println("|- "+v.Password, color.GreenString("\u2714"))
			fmt.Println("|- "+v.Sha1, color.GreenString("\u2714"))
			fmt.Println("|- "+v.Hash, color.GreenString("\u2714"))
		}
	} else {
		color.Red("|- No results found")
	}
}

func PrintLookup(lookup_temp_result [][]string) {
	fmt.Println("\nLookup Results:")
	lookup_result := tablewriter.NewWriter(os.Stdout)
	for _, v := range lookup_temp_result {
		lookup_result.Append(v)
	}
	lookup_result.Render()
}
