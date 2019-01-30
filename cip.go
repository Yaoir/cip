package main

// cip: convert integer(s) and print

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	)

var cmdname string = "cip"

// Print an error message to standard error, then exit with exit_value

func error_exit(msg string, exit_val int) {
//
	fmt.Fprintf(os.Stderr,"%s: %s\n",cmdname,msg)
	os.Exit(exit_val)
}

func main() {
	var n1, n2, i uint64
	var step uint64 = 1	// step for -r option
	var err error
	var printoct bool	// output octal numbers
	var printdec bool	// output decimal numbers (default)
	var printhex bool	// output hex numbers
	var printchar bool	// output Unicode characters
	var opt_range bool	// print a range of numbers
	var inputbase int	// constant numeric base for input
	var constwidth bool	// use constant width for output numbers. Pad with leading zeroes
	var separator string	// separator to use: defaults to space, not newline (correct?)

	flag.BoolVar(&printdec,"d",false,"output decimal numbers (default)")
	flag.BoolVar(&printoct,"o",false,"output octal numbers")
	flag.BoolVar(&printhex,"h",false,"output hexidecimal numbers")
	flag.BoolVar(&printhex,"x",false,"(alternate for -h) output hexidecimal numbers")
	flag.BoolVar(&printchar,"c",false,"output runes")
	flag.BoolVar(&opt_range,"r",false,"print range of numbers")
	flag.IntVar(&inputbase,"b",0,"input base (0, 8, 10, or 16)")
	flag.BoolVar(&constwidth,"w",false,"constant width output")
// if set to " ", then separate characters with spaces (instead of no spaces)
	flag.StringVar(&separator,"s","\n","separator character")
	flag.Parse()
//flag.PrintDefaults()

	// Check that ONLY ONE of -h/x, -c, -o, -d are true
	{
		var num int = 0
		if printhex  { num++ }
		if printdec  { num++ }
		if printoct  { num++ }
		if printchar { num++ }

		if num > 1 { error_exit("too many output bases specified - use only one of -o, -c, -d, -h, -x",1) }
	}

	if len(separator) > 1 {
// TODO: if separator[0] == '\' { process escapes for \n, \t, \r }
		error_exit("separator cannot be more than one character long",1)
	}

	if constwidth {
		error_exit("-w option is not implemented yet",1)
	}

	if inputbase != 0 {
		switch inputbase {
			case 8, 10, 16: // no problem!
			default:	// anything else: no good.
				error_exit("-b accepts only bases 8, 10, or 16",1)
		}
//		error_exit("-b option is not implemented yet",1)
	}

	if opt_range {
		if len(flag.Args()) > 3 { error_exit("too many arguments for -r option",1) }
		if len(flag.Args()) < 1 { error_exit("need argument(s) for -r option",1) }

		if len(flag.Args()) == 3 {
			n1, err = strconv.ParseUint(flag.Arg(0),0,64)
			if err != nil { error_exit("bad argument: upper limit",2) }
			n2, err = strconv.ParseUint(flag.Arg(1),0,64)
			if err != nil { error_exit("bad argument: lower limit",2) }
			step, err = strconv.ParseUint(flag.Arg(2),0,64)
			if err != nil { error_exit("bad argument: step",2) }
			if n2 < n1 { error_exit("bad range",2) }
		}

		if len(flag.Args()) == 2 {
			n1, err = strconv.ParseUint(flag.Arg(0),0,64)
			if err != nil { error_exit("bad argument",2) }
			n2, err = strconv.ParseUint(flag.Arg(1),0,64)
			if err != nil { error_exit("bad argument",2) }
			if n2 < n1 { error_exit("bad range",2) }
		}

		if len(flag.Args()) == 1 {
			n1 = 1
			n2, err = strconv.ParseUint(flag.Arg(0),0,64)
			if err != nil { error_exit("bad argument",2) }
		}

		for i = n1; i <= n2; i += step {
			if printhex {
				// print as hexidecimal
				fmt.Fprintf(os.Stdout,"%x",i)
			} else if printoct {
				// print as rune
				fmt.Fprintf(os.Stdout,"%o",i)
			} else if printchar {
				// print as char
				fmt.Fprintf(os.Stdout,"%c",rune(i))
			} else {
				// default: print as decimal
				fmt.Fprintf(os.Stdout,"%d",i)
			}
			if i < n2 && ! printchar { fmt.Fprintf(os.Stdout,"%s",separator) }
		}
		fmt.Fprintf(os.Stdout,"\n")
	}

	if ! opt_range {
		var scanner *bufio.Scanner
		if len(flag.Args()) == 0 {
			// read input from standard input
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			// read input from arguments
			s := strings.Join(flag.Args()," ")
			scanner = bufio.NewScanner(strings.NewReader(s))
		}

		scanner.Split(bufio.ScanWords)
// TODO: for runes, use scanner.Split(bufio.ScanRunes) ?  or convert to rune, then int(rune)?

		sep := ""	// Don't print space before first number

		for scanner.Scan() {
		//
// TODO: use input base instead of 0
//	add prefix for octal and decimal?
// what happens if abaff is scanned? Or 377?

			// Parse number in input base

			n1, err = strconv.ParseUint(scanner.Text(),inputbase,64)
			if err != nil { if len(sep) != 0 { fmt.Printf("\n") }; error_exit("bad input",2) }

			// Print it in the numberic base or as a Unicode character

			if printhex {
				// print as hexidecimal number
				fmt.Fprintf(os.Stdout,"%s%x",sep,n1)
			} else if printoct {
				// print as octal number
				fmt.Fprintf(os.Stdout,"%s%o",sep,n1)
			} else if printchar {
				// print as Unicode character
				fmt.Fprintf(os.Stdout,"%c",rune(n1))
			} else {
				// default: print as decimal number
				fmt.Fprintf(os.Stdout,"%s%d",sep,n1)
			}

			sep = separator
		}
		fmt.Fprintf(os.Stdout,"\n")
	}
}
