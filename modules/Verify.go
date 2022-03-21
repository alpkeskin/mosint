// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"strings"

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

func VerifyEmail(email string) VerifyStruct {
	v := VerifyStruct{}
	ret, err := verifier.Verify(email)
	if err != nil {
		v.Err = err
	}
	if !ret.Syntax.Valid {
		v.IsVerified = false
	} else {
		domain := strings.Split(email, "@")[1]
		v.IsVerified = true
		if !verifier.IsDisposable(domain) {
			v.IsDisposable = false
		} else {
			v.IsDisposable = true
		}
	}
	return v

}
