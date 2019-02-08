# Makefile for cip (convert integer and print) command

# change the below to the directory in your PATH where you want cip installed
BINDIR=/home/jay/.bin/elf

# Release date. For ronn, when making manual page
RELDATE=2019-02-03

cip: cip.go
	@go build -o cip cip.go
# install
	@cp cip $(BINDIR)

# Manual Page

man: cip.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="$(RELDATE)" cip.1.ronn >/dev/null 2>&1
	@gzip -f cip.1
	@mv cip.1.gz man1
	@man -l man1/cip.1.gz

install:
	@cp cip $(BINDIR)

wc:
	@wc -l cip.go

backup back bak:
	@cp -a cip.1.ronn *.go Makefile Examples.md README.md TODO test .bak
