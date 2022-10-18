// mosint v2.2
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin

package cmd

import (
	"strings"
	"sync"

	emailverifier "github.com/AfterShip/email-verifier"
)

type VerifyStruct struct {
	IsVerified   bool  `json:"is_verified"`
	IsDisposable bool  `json:"is_disposable"`
	Err          error `json:"err"`
}

var (
	verifier = emailverifier.NewVerifier().
		EnableAutoUpdateDisposable()
)

func VerifyEmail(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	ret, err := verifier.Verify(email)
	if err != nil {
		verify_result.Err = err
	}
	if !ret.Syntax.Valid {
		verify_result.IsVerified = false
	} else {
		domain := strings.Split(email, "@")[1]
		verify_result.IsVerified = true
		if !verifier.IsDisposable(domain) {
			verify_result.IsDisposable = false
		} else {
			verify_result.IsDisposable = true
		}
	}

}
