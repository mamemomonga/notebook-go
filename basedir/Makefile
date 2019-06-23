.PHONY: run clean

bin/basedir: *.go
	go build -o bin/basedir

run: bin/basedir
	bin/basedir

clean:
	rm -rf bin var
