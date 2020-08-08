
package main

import (
	"github.com/comail/colog"
	"log"
	"os"
)

func init() {
	colors := false
	if os.Getenv("TERM") != "" {
		colors = true
	}
	if os.Getenv("DEBUG") != "" {
		colog.SetMinLevel(colog.LTrace)
		colog.SetDefaultLevel(colog.LDebug)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: colors,
			Flag:   log.Ldate | log.Ltime | log.Lshortfile,
		})

	} else {
		colog.SetMinLevel(colog.LInfo)
		colog.SetDefaultLevel(colog.LWarning)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: colors,
			Flag:   log.Ldate | log.Ltime,
		})
	}

	colog.Register()
}

