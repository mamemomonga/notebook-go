package main

import (
	"golang.org/x/crypto/ssh"

	"bufio"
	"io/ioutil"
	"io"
	"log"
	"os"
)

func main() {
	runCommandSSH(runCommandSSHConfig{
		KeyFile: "etc/id_ed25519",
		User:    "username",
		Host:    "192.168.1.1",
		Port:    "22",
		Cmd:     "uname -a",
		Mode:    "oneshot", // local | oneshot | persistent
	})
}

type runCommandSSHConfig struct {
	KeyFile string
	User    string
	Host    string
	Port    string
	Cmd     string
	Mode    string
}

func runCommandSSH(cnf runCommandSSHConfig) {
	auth := []ssh.AuthMethod{}
	{
		key, err := ioutil.ReadFile(cnf.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatal(err)
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}
	sshConfig := &ssh.ClientConfig{
		User:            cnf.User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", cnf.Host+":"+cnf.Port, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	var stdout io.Reader
	var stderr io.Reader
	if cnf.Mode == "local" {
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Stdin = os.Stdin
	} else {
		var err error
		stdout, err = session.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		stderr, err = session.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}
	}

	monitor := func(h io.Reader, prefix string) {
		s := bufio.NewScanner(h)
		for s.Scan() {
			log.Printf("[%s] %s", prefix, s.Text())
		}
	}

	if cnf.Mode == "persistent" {
		go monitor(stdout, "STDOUT")
		go monitor(stderr, "STDERR")
	}

	err = session.Run(cnf.Cmd)
	if err != nil {
		log.Fatal(err)
	}
	if cnf.Mode == "oneshot" {
		monitor(stdout, "STDOUT")
		monitor(stderr, "STDERR")
	}
}
