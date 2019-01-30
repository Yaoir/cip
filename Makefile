# Makefile for Unicode utilities

# Release date. For ronn, when making manual page
RELDATE=2019-01-30

cip: cip.go
	@go build -o cip cip.go

cip-man: cip.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="$(RELDATE)" cip.1.ronn >/dev/null 2>&1
	@gzip -f cip.1
	@mv cip.1.gz man1
	@man -l man1/cip.1.gz

wc:
	@wc -l cip.go

utap:
	@go build -o utap utap.go

upr: upr.go
	@go build upr.go

backup back bak:
	@cp cip.1.ronn *.go Makefile README.md TODO .bak
