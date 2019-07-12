NAME := "configtool"
SRCS := $(shell find . -type f -name '*.go')

bin/$(NAME): $(SRCS)
	go build -o $@ .

clean:
	rm -rf bin

