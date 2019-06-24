package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var (
		i = flag.Int("int", 0, "数値")
		s = flag.String("str", "default", "文字列")
		b = flag.Bool("bool", false, "可否")
	)
	flag.Parse()

	if *i <= 0 {
		log.Println("int は 0以上でなければなりません")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !*b {
		log.Println("bool は trueである必要があります")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("INT:  %d\n", *i)
	log.Printf("STR:  %s\n", *s)
	log.Printf("BOOL: %v\n", *b)
}
