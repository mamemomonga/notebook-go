package main

import (
	"log"
)

func msgTitle(t, m string) {
	log.Printf("\033[44;1m %s \033[47;30;1m %s \033[0m", t, m)
}

func msgBlue(m string) {
	log.Printf("\033[44;1m   %s   \033[0m", m)
}

func msgYellow(m string) {
	log.Printf("\033[103;30m   %s   \033[0m", m)
}

func msgRed(m string) {
	log.Printf("\033[41;1m   %s   \033[0m", m)
}
