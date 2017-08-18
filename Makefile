# Main Makefile for erc-search

VPATH=	config:lib
GOBIN=	${GOPATH}/bin
OPTS=	-ldflags="-s -w" -v
SRCS=	erc-search.go cli.go ldap.go config.go machine.go people.go srv.go

all:	erc-search erc-search.exe

install:
	go install ${OPTS}

clean:
	go clean -v
	rm -f erc-search

erc-search:    ${SRCS}
	go build ${OPTS}

erc-search.exe:    ${SRCS}
	GOOS=windows go build ${OPTS}

push:
	git push --all
	git push --all origin
	git push --all backup
