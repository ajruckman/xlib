package xlib

import "os"

type config struct {
	PrependFile   bool
	EmailErrors   bool
	EmailServer   string
	SenderEmail   string
	ReceiverEmail string
	EmailSubject  string
}

var (
	Hostname  string
	LibConfig config
)

func init() {
	Hostname, _ = os.Hostname()

	LibConfig = config{
		PrependFile:   true,
		EmailErrors:   false,
		EmailServer:   "",
		SenderEmail:   "",
		ReceiverEmail: "",
		EmailSubject:  "Error on " + Hostname,
	}
}
