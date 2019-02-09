### Introduction

**cip** converts integers between numeric bases, and also to and from Unicode characters.

The manual page is included below.

See the file **Examples.md** in this repository for many examples showing how to use cip, including two short tutorials explaining Unicode issues that may be encountered by programmers.

### Quick Start

**cip** is written in Go. To compile it, you need to have Go installed.

To compile:

```
$ go build cip.go
```
or if you have GNU *make* installed:
```
$ make
```

To install the manual page, copy the file **man1/cip.1.gz** to the directory where your manual pages are located. On Linux, this is typically **/usr/share/man/man1**.

To install both the **cip** program and its manual page, edit the Makefile to set BINDIR and MANDIR appropriately, then run

```
$ sudo make install
```

### Manual Page

```
CIP(1)                           User Commands                          CIP(1)



NAME
       cip - convert integer and print

SYNOPSIS
       cip [-opts]

       cip [-opts] [ number | Unicode_string ] ...

       cip [-opts] -r [ min ] max [ step ]

DESCRIPTION
       cip(1)  accepts numbers or Unicode characters as input, and prints them
       as base 8, 10, or 16 numbers or as Unicode characters.

       The input can be either standard input or the command´s  arguments.  If
       there  are  no  non-option  arguments,  input is read from the standard
       input as space-separated words. The words  must  be  in  a  format  cip
       understands  as numbers. By default, numbers can be integers in decimal
       (base 10) notation, octal numbers with a leading 0 (zero), or hexideci‐
       mal numbers with a leading 0x. The numeric base of the number(s) can be
       set using the -ib=base option.

       When arguments are provided (without the -r flag,  which  is  explained
       below), cip concatenates the arguments with space characters as separa‐
       tors and treats the result the same way as would an input file.

       The exception to both of the above is when the -ic flag is provided. In
       that  case, the input (either from standard input or concatenated argu‐
       ments) is a string of characters with no separations between them.  cip
       converts the Unicode characters into their code points, which are inte‐
       gers.

       With the -r (range) option, cip generates a range of numbers to  print,
       rather than reading numbers from the input.

OPTIONS
   Output base or method
       Only one of the following may be supplied:

       -b   Numeric output in binary (base 2).

       -o   Numeric output in octal (base 8).

       -d   Numeric output in decimal (base 10).

       -h  (or  -x)   Numeric output in hexidecimal (base 16), using lowercase
	              letters a-f to represent digits 10-15.

       -H (or -X)   Numeric output in hexidecimal (base 16),  using  uppercase
	            letters A-F to represent digits 10-15.

       -U   Output in Unicode standard form for code points: U+DDDD, where the
            Ds are hexidecimal digits, using uppercase A-F.

       -c   Unicode character output.

   Input Base
       Only one of the following may be supplied:

       -ib base   Use base as the numeric base of the input. base can  be  any
                  integer  from  2 to 36. The letters A-Z, either uppercase or
		  lowercase, are used to represent digits greater than 9.

       -ic   Accept Unicode characters as input. The characters are  converted
             to their Unicode code points.

   Output Formatting
       -p string   Use string as a prefix to apply to every number in the out‐
       put.

       -s string   Use string to separate numbers in the output.

       -w   Constant-width output. When generating a sequence, pad the shorter
            numbers with 0s to make all of the numbers equal width.

       -width  number    Pad  all  numbers  in the output with 0s to make them
                         equal in width to number characters.

       Only one of -w and -width may be supplied.

   Other options
       -r [ min ] max [ step ]   Range. Output a sequence of integers or char‐
       acters.

       -help   Print a help message.

ARGUMENTS
       There  are two types of non-option arguments. When the -r option is not
       in use, supplying one or more arguments causes cip to accept input from
       the argument(s).

       When  the  -r option is specified, the arguments specify the range of a
       sequence of integers or characters to generate. With one argument,  the
       range  is  from  1 to the value of the argument. When two arguments are
       given, the range is from the first to the second argument.

       If the second argument is less than the first, cip will count downwards
       from the first towards the second.

       If an additional argument is given, it specifies the amount by which to
       increment or decrement while stepping through the sequence.

EXAMPLES
       Print 237 in binary

       cip -b 237

       Convert 0xED to decimal

       cip 0xed

       or

       cip -ib=16 ed

       or

       cip -d -ib=16 ED

       Print a sequence of octal numbers from 45 (base 8) to 177 (base 8)

       cip -o -r 045 0177

       Print the lowercase Greek alphabet

       cip -c -r 0x3b1 0x3c9

BUGS
       This is an early release, so there may be some. Report bugs to the author.

SEE ALSO
       od(1), hd(1), seq(1), numconv(1)

       Go Command Line Parsing
       https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax

       Go Integer Parsing
       https://golang.org/pkg/strconv/#ParseInt

       The Unicode Standard
       https://www.unicode.org/standard/standard.html

AUTHOR
       Jay Ts (http://jayts.com)



Jay Ts                           February 2019                          CIP(1)
```
