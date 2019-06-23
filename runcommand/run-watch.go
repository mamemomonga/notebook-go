package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"regexp"
)

// 正規表現にマッチした文字列がでてくるまで待つ
// （出現してもプロセスは停止されません)
func runWatch(re string, c string, p ...string) (err error) {

	r := regexp.MustCompile(re)

	cmd := exec.Command(c, p...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	doneChan := make(chan bool)
	defer close(doneChan)

	monitor := func(h io.ReadCloser) {
		scanner := bufio.NewScanner(h)
		for scanner.Scan() {
			t := scanner.Text()
			log.Println(t)
			if r.MatchString(t) {
				doneChan <- true
			}
		}
	}

	waiting := func() {
		cmd.Wait()
		doneChan <- true
	}

	go monitor(stdout)
	go monitor(stderr)

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go waiting()

	// doneChanが書き込まれるまでブロック
	_ = <-doneChan

	return nil
}
