package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"

	"github.com/pocketbase/pocketbase/tools/mailer"

	"pb-chad/email_templates"
)

func SendProduct(name string, email string, membership string) error {
	params := struct {
		Name       string
		Membership string
	}{
		Name: name, Membership: membership,
	}
	body, renderErr := resolveTemplateContent(
		params,
		email_templates.Layout,
		email_templates.ProductSuccessEmail,
	)

	if renderErr != nil {
		return renderErr
	}

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
			HTML:    body,
			// bcc, cc, attachments and custom headers are also supported...
		}
	default:
		fmt.Println("Something went wrong")
	}
	app.NewMailClient().Send(&message)
	return nil
}

func resolveTemplateContent(data any, content ...string) (string, error) {
	if len(content) == 0 {
		return "", nil
	}

	t := template.New("inline_template")

	var parseErr error
	for _, v := range content {
		t, parseErr = t.Parse(v)
		if parseErr != nil {
			return "", parseErr
		}
	}

	var wr bytes.Buffer

	if executeErr := t.Execute(&wr, data); executeErr != nil {
		return "", executeErr
	}

	return wr.String(), nil
}
