package main

import (
	"github.com/uvulpos/go-svelte/src/cmd"
	_ "github.com/uvulpos/go-svelte/swagger-docs"
)

func main() {

	// mailer := sendmail.NewMailer("mail.tim-riedl.de", 465, "sysintern@tim-riedl.de", "sysintern@tim-riedl.de", "4bbT::VQSZvUjS:iZLDn:mxTVQXFX3WnTL,PLDVqJRz:DWTdgJV3CPDen.7;4R;DrtqNFh")

	// mail := mailer.NewMail()
	// mail.SetReceipiants([]string{"mail@tim-riedl.de"})
	// mail.SetSubject("hallo")
	// mail.SetHTMLMessage("hallo welt")
	// mail.SendMail()

	cmd.Execute()
}
