# Main Makefile for erc-search

VPATH=	cmd/erc-search:config
GOBIN=	${GOPATH}/bin

SRCS=	erc-search.go cli.go ldap.go config.go defaults.go

all:	erc-search

install:
	go install -v

clean:
	go clean -v
	rm -f erc-search

erc-search:    ${SRCS}
	go build -v ./...

push:
	git push --all
	git push --all origin
	git push --all backup
