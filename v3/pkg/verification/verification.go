/*
Copyright © 2023 github.com/alpkeskin
*/
package verification

import (
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/alpkeskin/mosint/v3/internal/spinner"
)

type Verification struct {
}

var (
	verifier = emailverifier.NewVerifier()
)

func New() *Verification {
	return &Verification{}
}

func (v *Verification) Syntax(email string) bool {
	spinner := spinner.New("Email Syntax Verification")
	spinner.Start()

	ret, err := verifier.Verify(email)

	if err != nil {
		spinner.StopFail()
		spinner.SetMessage(err.Error())
	}

	if !ret.Syntax.Valid {
		spinner.StopFail()
		return false
	}

	spinner.Stop()
	return true
}
