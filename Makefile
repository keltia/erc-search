# Main Makefile for erc-search

VPATH=	erc-search:config
DEST=	bin
GOBIN=	${GOPATH}/bin

SRCS=	config.go erc-search.go cli.go

all:	${DEST}/erc-search

install:
	go install erc-search/erc-search.go erc-search/cli.go

clean:
	go clean -v
	rm -f ${DEST}/erc-search

${DEST}/erc-search:    ${SRCS}
	go build -v

push:
	git push --all
	git push --all origin
	git push --all backup
