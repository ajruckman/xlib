package xlib

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type F = map[string]interface{}

func Debug(msg string, meta ...F) {
	line("DEBUG", msg, defaultMeta(meta...))
}

func Info(msg string, meta ...F) {
	line("INFO", msg, defaultMeta(meta...))
}

func Warn(msg string, meta ...F) {
	line("WARN", msg, defaultMeta(meta...))
}

func Err(err error, meta ...F) {
	if err == nil {
		return
	}

	line("ERROR", err.Error(), defaultMeta(meta...))

	if LibConfig.EmailErrors {
		compose(err, defaultMeta(meta...))
	}

	panic(err)
}

func InfoOrErr(msg string, err error, meta ...F) {
	if err == nil {
		line("INFO", msg, defaultMeta(meta...))
	} else {
		m := defaultMeta(meta...)
		m["message"] = msg

		line("ERROR", err.Error(), defaultMeta(meta...))

		if LibConfig.EmailErrors {
			compose(err, m)
		}

		panic(err)
	}
}

func line(prefix, msg string, meta F) {
	var suffix string
	if len(meta) > 0 {
		suffix += "\t" + prettyMapFields(meta)
	}

	var level = "[" + prefix + "]"

	if LibConfig.PrependFile {
		fmt.Println(fmt.Sprintf("[%v] %-7s [%s] %s", time.Now().Format(`06-01-02T15:04:05`), level, caller(3), msg+suffix))
	} else {
		fmt.Println(fmt.Sprintf("[%v] %-7s %s", time.Now().Format(`06-01-02T15:04:05`), level, msg+suffix))
	}
}

func caller(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "?"
		line = 0
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func prettyMapFields(meta F) string {
	var res []string

	for k, v := range meta {
		res = append(res, fmt.Sprintf("%s=`%v`", k, v))
	}

	return strings.Join(res, " ")
}

func defaultMeta(meta ...F) F {
	if len(meta) == 0 {
		return nil

	} else if len(meta) > 1 {
		result := F{}

		for _, m := range meta {
			for k, v := range m {
				if _, ok := result[k]; !ok {
					result[k] = v
				}
			}

			return result
		}
	}

	return meta[0]
}
