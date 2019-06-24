package xlib

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"strings"

	"github.com/go-mail/mail"
)

func compose(err error, meta F) {
	if err == nil {
		return
	}

	b := make([]byte, 16384)
	n := runtime.Stack(b, false)

	stacktrace := strings.Replace(string(b[:n]), "\n", "<br>", -1)
	stacktrace = strings.Replace(stacktrace, "\t", "&emsp;&emsp;&emsp;&emsp;", -1)

	m := map[string]template.HTML{}

	if meta != nil {
		for k, v := range meta {
			val := fmt.Sprintf("%v", v)

			val = strings.Replace(fmt.Sprintf("%v", val), "\n", "<br>", -1)
			val = strings.Replace(fmt.Sprintf("%v", val), "\t", "&emsp;&emsp;&emsp;&emsp;", -1)

			m[k] = template.HTML(val)
		}
	}

	//layout, interr := template.ParseFiles("email.gohtml")
	layout, interr := template.New("emailContent").Parse(emailContent)
	if interr != nil {
		panic(interr)
	}

	var tmpl bytes.Buffer
	interr = layout.Execute(&tmpl, map[string]interface{}{
		"subject":    LibConfig.EmailSubject,
		"hostname":   Hostname,
		"error":      err.Error(),
		"stacktrace": template.HTML(stacktrace),
		"meta":       m,
	})
	if interr != nil {
		panic(interr)
	}

	email(LibConfig.EmailSubject, tmpl.String())
}

func email(subject, body string) {
	m := mail.NewMessage()
	m.SetHeader("From", LibConfig.SenderEmail)
	m.SetHeader("To", LibConfig.ReceiverEmail)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	d := mail.Dialer{Host: LibConfig.EmailServer, Port: 25}

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
