# Main Makefile for erc-search

GO=	 go
VPATH=	config:lib
GOBIN=	${GOPATH}/bin
OPTS=	-ldflags="-s -w" -v
SRCS=	erc-search.go cli.go ldap.go config.go machine.go people.go srv.go utils.go

BIN=	erc-search
EXE=	${BIN}.exe

all:	${BIN} ${EXE}

install:	${BIN}
	${GO} install ${OPTS} .

clean:
	${GO} clean -v .
	-/bin/rm -f ${BIN} ${EXE}

${BIN}:    ${SRCS}
	${GO} build ${OPTS} .

${EXE}:    ${SRCS}
	GOOS=windows ${GO} build ${OPTS} .

push:
	git push --all
	git push --tags
