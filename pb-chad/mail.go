package main

import (
	"fmt"
	"net/mail"

	"github.com/pocketbase/pocketbase/tools/mailer"
)

func SendProduct(email string, membership string) {
	var message mailer.Message
	switch membership {
	case "Chad":
		message = mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: email}},
			Subject: "Thank You For your Purchase of ChadKit",
			HTML:    "<p>Thanks you </p>",
			// bcc, cc, attachments and custom headers are also supported...
		}
	case "GigaChad":
		message = mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: email}},
			Subject: "Thank You For your Purchase of ChadKit",
			HTML:    "<p>Thanks you </p>",
			// bcc, cc, attachments and custom headers are also supported...
		}
	case "TestChad":
		message = mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: email}},
			Subject: "Thank You For your Purchase of ChadKit",
			HTML:    "<p>Thanks you </p>",
			// bcc, cc, attachments and custom headers are also supported...
		}
	default:
		fmt.Println("Something went wrong")
	}
	app.NewMailClient().Send(&message)
}
