package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

// 標準出力経由でtarを取得する
func runStdout2Expand(base, c string, p ...string) {
	cmd := exec.Command(c, p...)
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		tr := tar.NewReader(stdout)
		for {
			hdr, err := tr.Next()
			if err != nil {
				errChan <- err
				break
			}

			typeflag := string(hdr.Typeflag)
			fn := path.Join(base, hdr.Name)

			// http://www.redout.net/data/tar.html#typeflag
			// ディレクトリ
			if typeflag == "5" {
				log.Printf("Create: %s\n", fn)
				if err := os.MkdirAll(fn, 0755); err != nil {
					log.Fatal(err)
				}
				// ファイル
			} else if typeflag == "0" {
				log.Printf("Write: %s\n", fn)
				outfile, err := os.Create(fn)
				if err != nil {
					errChan <- err
					break
				}
				defer outfile.Close()
				if _, err := io.Copy(outfile, tr); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	err = <-errChan
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}

}
