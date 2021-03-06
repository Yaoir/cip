**Note to reader: The proper rendering of this document depends on the software you are using to view it.**

**It should appear correctly in the Atom editor's Markdown Preview.**

### Motivation

One of the main intents in developing **cip** is to help make Unicode a little easier for developers.

Here is an example showing how Unicode is not always as simple and easy as one might assume.

The Unicode code point

```
❤	U+2764	HEAVY BLACK HEART
```

was added before color emoji were added to Unicode. So originally it was just a black heart. Nowadays in Unicode, there are color emojis, and variation selectors U+FE0E and U+FE0F are used to indicate whether to show a character as a regular text character or as emoji.

U+2764 followed by U+FE0F (Variation Selector 16) indicates that the heart should be displayed as emoji, and most likely appears as a red-colored heart. When people cut and paste a red heart from a web page or other document, they are often pasting the two code points into another document.

It is also possible to indicate that the heart should be displayed in text style rather than emoji style by using U+FE0E (Variation Selector 15) after it.

As a result, a program might get any of these as a heart symbol:

* U+2764 alone, which software may render as a black or red heart.
* U+2764 followed by U+FE0F, indicating that it should be rendered _emoji style_, which software may render as ... a black or red heart.
* U+2764 followed by U+FE0E, indicating that it should be rendered _text style_, which (you guessed it) software may render as a black or red heart!

Variation selectors are not printing code points, although sometimes a replacement character will be shown when the text rendering engine cannot handle the variation selector properly. So when you see a red or black heart, you usually cannot tell what Unicode code point(s) produced it.

**cip** allow you to examine Unicode characters (or strings) to decipher the codes. Let's use it on a heart by cutting and pasting a heart from a document into a **cip** command.

```
$ cip -U -ic ❤️
U+2764
U+FE0F
```

Or if the heart is in a text file,

```
$ cip -U -ic <heart.txt
U+2764
U+FE0F
U+000A
```

(The U+000A character is a newline, which is typical as the last character in a text file.)

You can see that in this case, the heart was specified as U+2764 HEAVY BLACK HEART, followed by U+FE0F to specify emoji style rendering.

You can also use cip to output specific Unicode characters. Here is how to create **heart.txt**:

```
$ cip -c 0x2764 0xfe0f >heart.txt
```

or just print it to the terminal:

```
$ cip -c 0x2764 0xfe0f
❤️
```

You may also use Unicode code point notation.

```
$ cip -c U+2764 U+fe0f
❤️
```

#### Further Reading

Here is a page that shows variation selectors used with emoji:

[Unicode Character Database - Standardized Variants](http://www.unicode.org/Public/6.3.0/ucd/StandardizedVariants.html)

And that's just getting started. Consider that:

* Unicode has many other code points for hearts.
* Variation selectors U+FE0E and U+FE0F can also be used with other characters, not just emoji.
* There are many other variation selectors.
* There are many other kinds of Unicode code points that are invisible, either because they are defined as "non-printing" characters, or they modify how another character appears, or because software typically ignores them.

There are two more examples of "weird Unicode stuff" in [Examples.md](https://github.com/Yaoir/cip/blob/master/Examples.md) in this repository.

## Introduction to **cip**

**cip** converts integers between numeric bases, and also to and from Unicode characters.

Applications include:

* Numeric base conversion: from base 2-36 to octal, decimal, or hexidecimal, including U+*n* Unicode code point format.
* Sequence generation
* Unicode explorer: convert Unicode characters to or from octal, decimal, or hex Unicode code point.
* Analysis of Unicode strings

The manual page is included below.

See [Examples.md](https://github.com/Yaoir/cip/blob/master/Examples.md) for many examples showing how to use **cip**, including two short tutorials explaining Unicode issues that may be encountered by programmers.

## Compiling and Installing **cip**

**cip** is written in Go. To compile it, you need to have Go installed.

To compile:

```
$ go build cip.go
```
or if you have GNU **make** installed:
```
$ make
```

To install the manual page, copy the file **man1/cip.1.gz** to the directory where your manual pages are located. On Linux, this is typically **/usr/share/man/man1**.

To install **cip** program using **make**, edit **Makefile** to set BINDIR appropriately, then run

```
$ make install
```

To install the manual page using **make**, edit **Makefile** to set MANDIR appropriately, then run

```
$ sudo make installman
```

## Manual Page

A copy of the manual page is included here for convenience. To display it better, install it on your system and use the command

```
$ man cip
```

```
CIP(1)                           User Commands                          CIP(1)



NAME
       cip - convert integer and print

SYNOPSIS
       cip [-opts]

       cip [-opts] [ number | Unicode_string ] ...

       cip [-opts] -r [ min ] max [ step ]

DESCRIPTION
       cip  accepts numbers or Unicode characters as input, and prints them as
       base 2, 8, 10, or 16 numbers or as Unicode UTF-8 characters.

       When cip is not being used to generate a sequence (with the -r option),
       the  input can be either from the standard input or the command´s argu‐
       ments.

       The input is read as whitespace-separated words. The words must be in a
       format  cip understands as numbers. By default, numbers can be integers
       in decimal (base 10) notation, octal (base 8) numbers with a leading  0
       (zero),  hexidecimal  (base  16)  numbers with a leading 0x, or Unicode
       code points (a prefix of U+ followed by hexidecimal digits). A  mix  of
       numeric bases may be present.

       Alternatively,  the  numeric  base of the number(s) in the input can be
       set using the -ib=base option, in which case prefixes are not  allowed,
       and  all of the numbers must have the specified numeric base, which can
       be from 2 to 36.

       cip can output numbers only in bases 2, 8, 10, or  16.  The  -o  option
       specifies  base  8  (octal),  and  numbers are printed with a leading 0
       (zero). Hexidecimal output can be specified using any of -x, -X, -h, or
       -H,  and  the numbers are printed with a leading 0x prefix. Using -X or
       -H results in the numbers being  printed  using  capital  A-F  letters,
       rather than lowercase.

       To  avoid  printing of prefixes, the output base can be specified using
       either the -ob=8 or -ob=16 option. Either of -X or -H may be  added  to
       -ob=16 to specify that the numbers should be printed with uppercase A-F
       digits.

       When arguments are provided (without the -r flag,  which  is  explained
       below),  cip  concatenates the arguments with space characters (Unicode
       U+0020) as separators and treats the result the same way as it would an
       input file.

       The  exception  to  both of the above is when the -ic flag is provided,
       the input (either from standard input or concatenated arguments)  is  a
       string  of  characters with no separators between them. Internally, cip
       converts the Unicode characters into their (integer) code points.

       With the -r (range) option, cip generates a range of numbers to  print,
       rather than reading numbers from the input.

OPTIONS
       Options  may  be specified in any way supported by the flags package in
       the Go standard library.

   Output base or method
       By default, cip outputs decimal (base 10) integers. To output in  other
       formats or numeric bases, only one of the following may be supplied:

       -b   Numeric output in binary (base 2).

       -o   Numeric output in octal (base 8).

       -h  (or  -x)   Numeric output in hexidecimal (base 16), using lowercase
       letters a-f to represent digits 10-15.

       -H (or -X)   Numeric output in hexidecimal (base 16),  using  uppercase
       letters A-F to represent digits 10-15.

       -U   Output in Unicode standard form for code points: U+DDDD, where the
       Ds are hexidecimal digits, using uppercase A-F.

       -ob _base_   Set the numeric base for output to base. Only values of 2,
       8, 10, and 16 are allowed.

       -c   Unicode character output.

   Input Base
       By default, cip interprets input numbers as decimal (base 10) integers.
       To accept input in other bases or as Unicode characters,  only  one  of
       the following may be supplied:

       -ib  base    Use base as the numeric base of the input. base can be any
       integer from 2 to 36. The letters A-Z, either uppercase  or  lowercase,
       are used to represent digits greater than 9.

       -ic    Accept Unicode characters as input. The characters are converted
       to their Unicode code points.

   Output Formatting
       -n   When using -c to output Unicode characters, don´t print a  newline
       at the end of the line.

       -p string   Use string as a prefix to apply to every number in the out‐
       put.

       -s string   Use string to separate numbers in the  output.  string  may
       include  \t, \n and other characters recognized by strconv.Unquote() in
       the Go standard library.

       -w   Constant-width output. When generating a sequence, pad the shorter
       numbers with 0s to make all of the numbers equal width.

       -width  number   Pad all numbers in the output with 0s to make them (at
       least) equal in width to number characters.

       Only one of -w and -width may be supplied. The  -w  option  works  only
       with the -r option.

   Other options
       -r [ min ] max [ step ]   Range. Output a sequence of integers or char‐
       acters starting at min and ending at (or before) max,  incrementing  or
       decrementing  by  step  (min  can be greater than max, resulting in cip
       counting downwards by step. By using the -ic option, min  and  max  can
       both be specified as Unicode characters, but step must always be speci‐
       fied by a positive base 10 integer.

       -help   Print a help message.

ARGUMENTS
       There are two types of non-option arguments. When the -r option is  not
       in use, supplying one or more arguments causes cip to accept input from
       the argument(s).

       When the -r option is specified, the arguments specify the range  of  a
       sequence  of integers or characters to generate. With one argument, the
       range is from 1 to the value of the argument. When  two  arguments  are
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

       cip -ib=16 ED

       Print a sequence of octal numbers from 90 (base 10) to 177 (base 8)

       cip -o -r 90 0177

       or

       cip -ob=8 -r 90 0177

       Print the lowercase Greek alphabet

       cip -c -r 0x3b1 0x3c9

CAVEAT UTILITOR
       Omitting the < for standard input redirection.

       cip -ic -c input

       prints "input", not the contents of a file named input.

SEE ALSO
       od(1), hd(1), seq(1), numconv(1)

       Option  Parsing https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax

       Integer Parsing https://golang.org/pkg/strconv/#ParseInt

       strconv.Unquote() https://golang.org/pkg/strconv/#Unquote

       The Unicode Standard https://www.unicode.org/standard/standard.html

AUTHOR
       Jay Ts (http://jayts.com)



Jay Ts                           February 2019                          CIP(1)
```
