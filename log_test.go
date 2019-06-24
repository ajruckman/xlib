package xlib

import (
	"errors"
	"fmt"
	"testing"
)

func TestCaller(t *testing.T) {
	fmt.Println(caller(1))
}

func TestDebug(t *testing.T) {
	Debug("asdf", nil)
	Debug("asdf", F{
		"hello": "world",
		"line":  2,
	})
}

func TestInfo(t *testing.T) {
	Info("asdf", nil)
	Info("asdf", F{
		"hello": "world",
		"line":  2,
	})
}

func TestWarn(t *testing.T) {
	Warn("asdf", nil)
	Warn("asdf", F{
		"hello": "world",
		"line":  2,
	})
}

func TestErr(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		Err(errors.New("asdf"), nil)

	}()

	func() {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		Err(errors.New("asdf"), F{
			"hello": "world",
			"line":  2,
		})
	}()
}

func TestInfoOrErr(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		InfoOrErr("running test", errors.New("asdf"), nil)
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		InfoOrErr("running test", errors.New("asdf"), F{
			"hello": "world",
			"line":  2,
		})
	}()
}
