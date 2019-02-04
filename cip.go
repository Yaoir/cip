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

// Print an error message to standard error and exit

func error_exit(msg string, exit_val int) {
//
	fmt.Fprintf(os.Stderr,"%s: %s\n",cmdname,msg)
	os.Exit(exit_val)
}

func main() {
	var n1, n2, i uint64
	var step uint64 = 1	// step for -r option
	var err error
	// options:
	var help bool		// print help message
	var printbin bool	// output octal numbers
	var printoct bool	// output octal numbers
	var printdec bool	// output decimal numbers (default)
	var printhex bool	// output hex numbers
	var printhexc bool	// output hex numbers
	var printuni bool	// output U+00xx Unicode code points
	var printchar bool	// output Unicode characters
	var inputchar bool	// input is Unicode characters
	var sequence bool	// print a range of numbers
	var downwards bool	// counting downwards, where range is specified with a larger than smaller number
	var inputbase int	// constant numeric base for input
	var constwidth bool	// use constant width for output numbers. Pad with leading zeroes
	var separator string	// separator to use: defaults to space, not newline (correct?)
	var prefix string	// 0, 0x, U+, or other prefix string for numeric output
	var width int		// -w <width> option: print all with the same width, padded by 0s

	flag.BoolVar(&help,"help",false,"Print this help message.")
	flag.BoolVar(&printbin,"b",false,"Output binary numbers")
	flag.BoolVar(&printdec,"d",false,"Output decimal numbers (default)")
	flag.BoolVar(&printoct,"o",false,"Output octal numbers")
	flag.BoolVar(&printhex,"h",false,"Output hexidecimal numbers (using a-f)")
	flag.BoolVar(&printhexc,"H",false,"Output hexidecimal numbers (using A-F)")
	flag.BoolVar(&printhex,"x",false,"(alternate for -h) Output hexidecimal numbers (using a-f)")
	flag.BoolVar(&printhexc,"X",false,"(alternate for -H) Output hexidecimal numbers (using A-F)")
	flag.BoolVar(&printuni,"U",false,"Output numbers in Unicode spec (for example, U+006A)")
	flag.BoolVar(&printchar,"c",false,"Output Unicode characters")
	flag.BoolVar(&sequence,"r",false,"Print range of numbers")
	flag.IntVar(&inputbase,"ib",0,"Input `base`: Use 0, 8, 10, or 16 for numbers, and 1 for character input.")
	flag.BoolVar(&inputchar,"ic",false,"Unicode character input.")
	flag.BoolVar(&constwidth,"w",false,"Constant width output (fitting widest), padded with leading 0s")
	flag.IntVar(&width,"width",0,"Constant `width` output, padded with leading 0s")
	// for -c flag: if separator is set to " ", then separate characters with spaces (instead of no spaces)
	flag.StringVar(&separator,"s","\uffff","`separator` string")
	flag.StringVar(&prefix,"p","","numeric `prefix` (0, 0x, U+, ...")
	flag.Parse()

	if help {
		fmt.Fprintf(os.Stderr,"cip - Convert Integer and Print\n")
		fmt.Fprintf(os.Stderr,"\n")
		fmt.Fprintf(os.Stderr,"usage:\n")
		fmt.Fprintf(os.Stderr,"\tcip [ -opts ]\n")
		fmt.Fprintf(os.Stderr,"\t\tread input from standard input\n")
		fmt.Fprintf(os.Stderr,"\tcip [ -opts ] args...\n")
		fmt.Fprintf(os.Stderr,"\t\tread input from arguments\n")
		fmt.Fprintf(os.Stderr,"\tcip [ -opts ] [ min ] max [ step ]\n")
		fmt.Fprintf(os.Stderr,"\t\tgenerate a sequence from min to max by intervals of step\n")
		fmt.Fprintf(os.Stderr,"\n")
		fmt.Fprintf(os.Stderr,"Options:\n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Check that ONLY ONE of -h/x, -c, -b, -o, -d are set
	{
		var num int = 0
		if printhexc { num++ }
		if printhex  { num++ }
		if printdec  { num++ }
		if printoct  { num++ }
		if printbin  { num++ }
		if printuni  { num++ }
		if printchar { num++ }

		if num > 1 { error_exit("too many output bases specified - use only one of -b, -o, -d, -h, -H, -x, -X, -c",2) }
		// in case someone used '-d=false' option:
//		if num < 1 { error_exit("one of -b, -o, -d, -h, -H, -x, -X, -c must be specified",2) }
	}

	// Check that ONLY ONE of -w and -width are set
	{
		var num int = 0
		if constwidth { num++ }
		if width > 0  { num++ }
		if num > 1 { error_exit("cannot use both -w and -width options at the same time",2) }
	}

	// inputbase must be from 2 to 62

	if inputbase != 0 && (inputbase < 2 || inputbase > 36) {
		error_exit("input base must be from 2 to 36",2)
	}

	// -c option: default is to print no separator
	// For all others, default separator is a newline

	if ! printchar && separator == "\uffff" { separator = "\n" }
	if printchar && separator == "\uffff" { separator = "" }

	// process escapes for \n, \t, \r, etc. in separator

	if len(separator) > 1 {
		separator, err = strconv.Unquote("\"" + separator + "\"")
		if err != nil { error_exit("bad character(s) in separator",2) }
	}

	if width < 0 {
		error_exit("-width does not work with a negative width",2)
	}

	// -r (range) option, for printing sequences:

	if sequence {
	//
		if len(flag.Args()) > 3 { error_exit("too many arguments for -r option",2) }
		if len(flag.Args()) < 1 { error_exit("need argument(s) for -r option",2) }

		// -r min max step

		if len(flag.Args()) == 3 {
			if inputchar || inputbase == 1 {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range lower limit contains more than one character",2) }
				n1 = uint64(([]rune)(flag.Arg(0))[0])
				if len(([]rune)(flag.Arg(1))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(1))[0])
// TODO: step should be integer, not character
//				if len(flag.Arg(2)) > 1 { error_exit("range step contains more than one character",2) }
//				step = uint64(flag.Arg(2)[0])
				step, err = strconv.ParseUint(flag.Arg(2),inputbase,64)
				if err != nil { error_exit("range step is not a number",2) }
			} else {
				n1, err = strconv.ParseUint(flag.Arg(0),inputbase,64)
				if err != nil { error_exit("range upper limit is not a number",2) }
				n2, err = strconv.ParseUint(flag.Arg(1),inputbase,64)
				if err != nil { error_exit("range lower limit is not a number",2) }
				step, err = strconv.ParseUint(flag.Arg(2),inputbase,64)
				if err != nil { error_exit("range step is not a number",2) }
			}
		}

		// -r min max
		// step defaults to 1

		if len(flag.Args()) == 2 {
			if inputchar || inputbase == 1 {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n1 = uint64(([]rune)(flag.Arg(0))[0])
				if len(([]rune)(flag.Arg(1))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(1))[0])
			} else {
				n1, err = strconv.ParseUint(flag.Arg(0),inputbase,64)
				if err != nil { error_exit("range lower limit is not a number",2) }
				n2, err = strconv.ParseUint(flag.Arg(1),inputbase,64)
				if err != nil { error_exit("range upper limit is not a number",2) }
			}
		}

		// -r max
		// min and step both default to 1

		if len(flag.Args()) == 1 {
			n1 = 1
			if inputchar || inputbase == 1 {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(0))[0])
			} else {
				n2, err = strconv.ParseUint(flag.Arg(0),inputbase,64)
				if err != nil { error_exit("range upper limit is not a number",2) }
			}
		}

		sep := ""	// Don't print space before first number

		// find maximum width of numbers in sequence

		if constwidth {
		//
			var f string
			num := n2
			if n1 > n2 { num = n1 }

			switch {
				case printhex || printhexc:
						f = "%x"
				case printoct:	f = "%o"
				case printbin:	f = "%b"
				default:	f = "%d"
			}
			s := fmt.Sprintf(f,num)
			width = len(s)
		}
		
		if n1 > n2 {
			// counting downwards from n1 to n2
			downwards = true
			step = -step
		}

		// Print Output

		for i = n1; ( ! downwards && i <= n2) || (downwards && i >= n2); i += step {
		//
			if printhex {
				// print as hexidecimal number using a-f
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*x",sep,prefix,width,i)
			} else if printhexc {
				// print as hexidecimal number using A-F
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*X",sep,prefix,width,i)
			} else if printoct {
				// print as octal number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*o",sep,prefix,width,i)
			} else if printbin {
				// print as binary number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*b",sep,prefix,width,i)
			} else if printuni {
				// print as U+HHHH
				fmt.Fprintf(os.Stdout,"%s%s%U",sep,prefix,i)
			} else if printchar {
				// print as Unicode character
				fmt.Fprintf(os.Stdout,"%s%c",sep,rune(i))
			} else {
				// default: print as decimal number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*d",sep,prefix,width,i)
			}
			sep = separator
		}
		fmt.Fprintf(os.Stdout,"\n")
	}

	// When not doing a sequence (range), read from standard input or arguments

	if ! sequence {
	//
		var scanner *bufio.Scanner
		var sep string

		if constwidth { error_exit("-w cannot be used without -r",2) }

		if len(flag.Args()) == 0 {
			// read input from standard input
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			// read input from arguments
			if inputchar || inputbase == 1 { sep = "" } else { sep = " " }
			s := strings.Join(flag.Args(),sep)
			scanner = bufio.NewScanner(strings.NewReader(s))
		}

		// input runes if inputchar is true or inputbase is 1, otherwise words to convert to integers
		if inputchar || inputbase == 1 { scanner.Split(bufio.ScanRunes) } else { scanner.Split(bufio.ScanWords) }

		sep = ""	// Don't print a separator string before the first number

		for scanner.Scan() {
		//
			if inputchar || inputbase == 1 {
				// Get the rune into an integer
				n1 = uint64(([]rune)(scanner.Text())[0])
			} else {
				// Parse number in input base
				n1, err = strconv.ParseUint(scanner.Text(),inputbase,64)
				if err != nil {
					if len(sep) != 0 { fmt.Printf("\n") }
					if numError, ok := err.(*strconv.NumError); ok {
						if numError.Err == strconv.ErrRange {
							fmt.Fprintf(os.Stderr,"%s: base %d number %s out of range\n",cmdname,inputbase,scanner.Text())
							os.Exit(2)
						}
					}
					if numError, ok := err.(*strconv.NumError); ok {
						if numError.Err == strconv.ErrSyntax {
							fmt.Fprintf(os.Stderr,"%s: syntax error in base %d number %s\n",cmdname,inputbase,scanner.Text())
							os.Exit(2)
						}
					}
					fmt.Fprintf(os.Stderr,"%s: cannot convert %s to a number\n",cmdname,scanner.Text())
					os.Exit(2)
				}
			}

			// Print it in the numeric base or as a Unicode character

			if printhex {
				// print as hexidecimal number using a-f
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*x",sep,prefix,width,n1)
			} else if printhexc {
				// print as hexidecimal number using A-F
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*X",sep,prefix,width,n1)
			} else if printoct {
				// print as octal number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*o",sep,prefix,width,n1)
			} else if printbin {
				// print as binary number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*b",sep,prefix,width,n1)
			} else if printuni {
				// print as U+HHHH
				fmt.Fprintf(os.Stdout,"%s%s%U",sep,prefix,n1)
			} else if printchar {
				// print as Unicode character
				fmt.Fprintf(os.Stdout,"%s%c",sep,rune(n1))
			} else {
				// default: print as decimal number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*d",sep,prefix,width,n1)
			}

			sep = separator
		}
		// print a newline when outputting numbers or reading from the command's arguments
		if ! printchar || (printchar && len(flag.Args()) != 0) { fmt.Fprintf(os.Stdout,"\n") }
	}
}
