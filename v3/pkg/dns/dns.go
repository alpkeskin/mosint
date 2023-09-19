/*
Copyright © 2023 github.com/alpkeskin
*/
package dns

import (
	"fmt"
	"strings"

	"github.com/alpkeskin/mosint/v3/internal/spinner"
	"github.com/domainr/dnsr"
	"github.com/fatih/color"
)

type Dns struct {
	Records []Record
}

type Record struct {
	Type  string
	Value string
}

func New() *Dns {
	return &Dns{}
}

func (d *Dns) Resolver(email string) {
	spinner := spinner.New("DNS Lookup")
	spinner.Start()
	defer spinner.Stop()

	domain := email[strings.IndexByte(email, '@')+1:]
	r := dnsr.NewResolver(dnsr.WithCache(10000))
	arr, err := r.ResolveErr(domain, "TXT")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	for _, rr := range arr {
		d.Records = append(d.Records, Record{
			Type:  rr.Type,
			Value: rr.Value,
		})
	}

	arr, err = r.ResolveErr(domain, "MX")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	for _, rr := range arr {
		d.Records = append(d.Records, Record{
			Type:  rr.Type,
			Value: rr.Value,
		})
	}

	arr, err = r.ResolveErr(domain, "A")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	for _, rr := range arr {
		d.Records = append(d.Records, Record{
			Type:  rr.Type,
			Value: rr.Value,
		})
	}

	arr, err = r.ResolveErr(domain, "CNAME")

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
		return
	}

	for _, rr := range r.Resolve(domain, "CNAME") {
		d.Records = append(d.Records, Record{
			Type:  rr.Type,
			Value: rr.Value,
		})
	}

	// remove duplicates
	for i := 0; i < len(d.Records); i++ {
		for j := i + 1; j < len(d.Records); j++ {
			if d.Records[i].Type == d.Records[j].Type && d.Records[i].Value == d.Records[j].Value {
				d.Records = append(d.Records[:j], d.Records[j+1:]...)
				j--
			}
		}
	}
}

func (d *Dns) Print() {
	fmt.Println("[*] DNS Lookup")
	for _, record := range d.Records {
		fmt.Println(color.GreenString("[+]"), record.Type, record.Value)
	}
}
