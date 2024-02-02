/*
Copyright © 2023 github.com/alpkeskin
*/
package runner

import (
	"sync"

	"github.com/alpkeskin/mosint/v3/pkg/dns"
	"github.com/alpkeskin/mosint/v3/pkg/scrape/googlesearch"
	"github.com/alpkeskin/mosint/v3/pkg/services/breachdirectory"
	"github.com/alpkeskin/mosint/v3/pkg/services/emailrep"
	"github.com/alpkeskin/mosint/v3/pkg/services/haveibeenpwned"
	"github.com/alpkeskin/mosint/v3/pkg/services/hunter"
	"github.com/alpkeskin/mosint/v3/pkg/services/intelx"
	"github.com/alpkeskin/mosint/v3/pkg/services/ipapi"
	"github.com/alpkeskin/mosint/v3/pkg/services/psbdmp"
	"github.com/alpkeskin/mosint/v3/pkg/social/instagram"
	"github.com/alpkeskin/mosint/v3/pkg/social/spotify"
	"github.com/alpkeskin/mosint/v3/pkg/social/twitter"
	"github.com/gammazero/workerpool"
)



type Runner struct {
	Email            string
	DnsC             *dns.Dns
	GoogleSearchC    *googlesearch.GoogleSearch
	BreachDirectoryC *breachdirectory.BreachDirectory
	HaveibeenpwnedC  *haveibeenpwned.HaveIBeenPwned
	EmailRepC        *emailrep.Emailrep
	HunterC          *hunter.Hunter
	IntelxC          *intelx.Intelx
	IpApiC           *ipapi.Ipapi
	PsbdmpC          *psbdmp.Psbdmp
	InstagramC       *instagram.Instagram
	SpotifyC         *spotify.Spotify
	TwitterC         *twitter.Twitter
}

func New(email string) *Runner {
	return &Runner{
		Email:            email,
		DnsC:             dns.New(),
		GoogleSearchC:    googlesearch.New(),
		BreachDirectoryC: breachdirectory.New(),
		HaveibeenpwnedC:  haveibeenpwned.New(),
		EmailRepC:        emailrep.New(),
		HunterC:          hunter.New(),
		IntelxC:          intelx.New(),
		IpApiC:           ipapi.New(),
		PsbdmpC:          psbdmp.New(),
		InstagramC:       instagram.New(),
		SpotifyC:         spotify.New(),
		TwitterC:         twitter.New(),
	}
}

func (c *Runner) Start() {
	email := c.Email
	runners := []func(string){
		c.DnsC.Resolver,
		c.GoogleSearchC.Search,
		c.BreachDirectoryC.Lookup,
		c.HaveibeenpwnedC.Lookup,
		c.EmailRepC.Lookup,
		c.HunterC.Lookup,
		c.IntelxC.Search,
		c.IpApiC.GetInfo,
		c.PsbdmpC.Search,
		c.InstagramC.Check,
		c.SpotifyC.Check,
		c.TwitterC.Check,
	}

	wp := workerpool.New(12)
	for _, runner := range runners {
		runner := runner
		wp.Submit(func() {
			runner(email)
		})
	}
	wp.StopWait()
}

func (c *Runner) Print() {
	println()
	c.EmailRepC.Print()

	println()
	c.HunterC.Print()

	println()
	c.GoogleSearchC.Print()

	println()
	c.InstagramC.Print()
	c.SpotifyC.Print()
	c.TwitterC.Print()

	println()
	c.PsbdmpC.Print()

	println()
	c.IntelxC.Print()

	println()
	c.BreachDirectoryC.Print()

	println()
	c.HaveibeenpwnedC.Print()

	println()
	c.IpApiC.Print()

	println()
	c.DnsC.Print()
}

type MultiRunner struct {
	Emails []string
	Results map[string]*Runner
	ConcurrencyLimit int
}

func NewMultiRunner(emails []string, concurrencyLimit ...int) *MultiRunner {
    var limit int
    if len(concurrencyLimit) > 0 && concurrencyLimit[0] > 0 {
        limit = concurrencyLimit[0]
    } else {
        limit = 10
    }
    results := make(map[string]*Runner)
    for _, email := range emails {
        results[email] = New(email)
    }
    return &MultiRunner{
        Emails:          emails,
        Results:         results,
        ConcurrencyLimit: limit,
    }
}

func (mr *MultiRunner) Start() {
    wp := workerpool.New(mr.ConcurrencyLimit)
    var wg sync.WaitGroup
    for _, runner := range mr.Results {
        wg.Add(1)
        wp.Submit(func() {
            defer wg.Done()
            runner.Start()
        })
    }
    wg.Wait()
    wp.StopWait()
}
