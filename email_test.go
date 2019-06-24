package xlib

import (
	"io"
	"testing"
)

func TestEmail(t *testing.T) {
	LibConfig.EmailErrors = true

	compose(io.EOF, F{"test:": 2})
	compose(io.EOF, F{
		"test:": 1,
		"hello": "world",
		"msg":   "invalid memory address or nil pointer dereference",
		"long": ` 	https://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a
            https://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.go

            /std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.gohttps://github.com/comail/colog/blob/fba8e7b1f46c3607f09760ce3880066e7ff57c5a/std_formatter.go`,
	})
	compose(nil, F{"test:": 3})
	compose(nil, nil)
}
