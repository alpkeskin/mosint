package modules

import (
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/fatih/color"
)

var (
	verifier = emailverifier.NewVerifier()
)

func Verify_email(email string) {
	ret, err := verifier.Verify(email)
	if err != nil {
		color.Red("[-] Verify email address failed, error is: ", err)

	}
	if !ret.Syntax.Valid {
		color.Red("[-] Email address syntax is invalid!")
	}
	color.Green("[+] Email verified!")

}
