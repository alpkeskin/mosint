// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"net"
	"strings"
	"sync"
)

func Lookup(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	splt := strings.Split(email, "@")
	iprecords, _ := net.LookupIP(splt[1])
	for _, ip := range iprecords {
		row := []string{"IP", ip.String()}
		lookup_temp_result = append(lookup_temp_result, row)
	}
	nameserver, _ := net.LookupNS(splt[1])
	for _, ns := range nameserver {
		row := []string{"NS", ns.Host}
		lookup_temp_result = append(lookup_temp_result, row)
	}
	mxrecords, _ := net.LookupMX(splt[1])
	for _, mx := range mxrecords {
		row := []string{"MX", mx.Host}
		lookup_temp_result = append(lookup_temp_result, row)
	}
	txtrecords, _ := net.LookupTXT(splt[1])
	for _, txt := range txtrecords {
		row := []string{"TXT", txt}
		lookup_temp_result = append(lookup_temp_result, row)
	}
}
