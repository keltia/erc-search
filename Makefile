# Main Makefile for erc-search

VPATH=	cmd/erc-search:config
GOBIN=	${GOPATH}/bin
OPTS=	-ldflags="-s -w" -v
SRCS=	erc-search.go cli.go ldap.go config.go defaults.go

all:	erc-search

install:
	go install ${OPTS} ./...

clean:
	go clean -v
	rm -f erc-search

erc-search:    ${SRCS}
	go build ${OPTS} ./...

push:
	git push --all
	git push --all origin
	git push --all backup
