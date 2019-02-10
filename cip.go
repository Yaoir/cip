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

// Wrapper around strconv.ParseUint() to support "U+" prefix for Unicode code points

func get_uint64(s string, base int) (uint64, error) {
//
	// recognize U+ prefix and change it to 0x
	if s[0] == 'U' && s [1] == '+' { s = "0x" + string(s[2:]) }
	return strconv.ParseUint(s,base,64)
}

func main() {
	var n1, n2, i uint64
	var step uint64 = 1	// step for -r option.
	var err error
	// options:
	var help bool		// print help message
	var printbin bool	// output octal numbers
	var printoct bool	// output octal numbers
	var printhex bool	// output hex numbers
	var printhexc bool	// output hex numbers
	var printuni bool	// output U+00xx Unicode code points
	var printchar bool	// output Unicode characters
	var inputchar bool	// input is Unicode characters
	var sequence bool	// print a range of numbers
	var downwards bool	// counting downwards, where range is specified with a larger than smaller number
	var inputbase int	// numeric base for input: 2-36
	var outputbase int	// numeric base for output: only 2, 8, 10, or 16
	var constwidth bool	// use constant width for output numbers. Pad with leading zeroes
	var separator string	// separator to use: defaults to space, not newline (correct?)
	var prefix string	// prefix string for numeric output
	var width int		// -w <width> option: print all with the same width, padded by 0s
	var nonewline bool	// -n don't print newline at end of output when using -c -ic and taking data from flag.Args()

	flag.BoolVar(&help,"help",false,"Print this help message.")
	flag.BoolVar(&printbin,"b",false,"Output binary numbers")
	flag.BoolVar(&printoct,"o",false,"Output octal numbers")
	flag.BoolVar(&printhex,"h",false,"Output hexidecimal numbers (using a-f)")
	flag.BoolVar(&printhexc,"H",false,"Output hexidecimal numbers (using A-F)")
	flag.BoolVar(&printhex,"x",false,"(alternate for -h) Output hexidecimal numbers (using a-f)")
	flag.BoolVar(&printhexc,"X",false,"(alternate for -H) Output hexidecimal numbers (using A-F)")
	flag.BoolVar(&printuni,"U",false,"Output numbers in Unicode spec (for example, U+006A)")
	flag.BoolVar(&printchar,"c",false,"Output Unicode characters")
	flag.BoolVar(&sequence,"r",false,"Print range of numbers")
	flag.IntVar(&inputbase,"ib",0,"Input `base`: 2-36")
	flag.IntVar(&outputbase,"ob",0,"Output `base`: Use 2, 8, 10, or 16")
	flag.BoolVar(&inputchar,"ic",false,"Unicode character input.")
	flag.BoolVar(&constwidth,"w",false,"Constant width output (fitting widest), padded with leading 0s")
	flag.IntVar(&width,"width",0,"Constant `width` output, padded with leading 0s")
	// for -c flag: if separator is set to " ", then separate characters with spaces (instead of no spaces)
	flag.StringVar(&separator,"s","\uffff","`separator` string")
	flag.StringVar(&prefix,"p","","numeric `prefix` (0, 0x, U+, ...)")
	flag.BoolVar(&nonewline,"n",false,"don't print newline after Unicode character(s)")
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

	// Check that ONLY ONE of -h/x, -c, -b, -o are set
	{
		var num int = 0
		if printhexc { num++ }
		if printhex  { num++ }
		if printoct  { num++ }
		if printbin  { num++ }
		if printuni  { num++ }
		if printchar { num++ }

		if num > 1 { error_exit("too many output bases specified - use only one of -b, -o, -h, -H, -x, -X, -c",2) }
	}

	// disallow -o with -ob=8

	if printoct && outputbase == 8 { error_exit("-o may not be used with -ob=8",2) }

	// But -[xXhH] may be used with -ob=16 to specify
	// capital/lowercase A-F with non-0x-prefixed hex output

	// Check that ONLY ONE of -w and -width are set
	{
		var num int = 0
		if constwidth { num++ }
		if width > 0  { num++ }
		if num > 1 { error_exit("cannot use both -w and -width options at the same time",2) }
	}

	// check that if -n flag is used, that both -c and -ic are in use as well,
	// and there are arguments

	if nonewline && ! printchar {
		error_exit("the -n option requires the -c option",2)
	}

	// Input Base: strconv.ParseUint() can handle bases from 2 to 36

	if inputbase != 0 && (inputbase < 2 || inputbase > 36) {
		error_exit("input base must be from 2 to 36",2)
	}

	// Output Base: fmt.Printf() can handle bases of 2, 8, 10, and 16
	// Also set the alternate format part of the printf format verb so that
	// by default, octal numbers are printed with a suffix of "0", and hex with "0x"

	if outputbase != 0 {
		switch outputbase {
			case 2, 8, 10, 16: // All of these are acceptable. Do nothing.
			default: error_exit("(-ob option) output base must be 2, 8, 10, or 16",2)
		}
	}

	// Set default separator if -s option hasn't overridden it.
	// For the -c option, the default is to print no separator.
	// For all others, the default separator is a newline.

	if separator == "\uffff" {
		if printchar { separator = "" } else { separator = "\n" }
	}

	// process escapes for \n, \t, \r, etc. in separator

	if len(separator) > 1 {
		separator, err = strconv.Unquote("\"" + separator + "\"")
		if err != nil { error_exit("bad character(s) in separator",2) }
	}

	if width < 0 { error_exit("-width does not work with a negative width",2) }

	// -r (range) option, for printing sequences:

	if sequence {
	//
		if len(flag.Args()) > 3 { error_exit("too many arguments for -r option",2) }
		if len(flag.Args()) < 1 { error_exit("need argument(s) for -r option",2) }

		// 3 arguments: -r min max step

		if len(flag.Args()) == 3 {
			if inputchar {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range lower limit contains more than one character",2) }
				n1 = uint64(([]rune)(flag.Arg(0))[0])
				if len(([]rune)(flag.Arg(1))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(1))[0])
				step, err = get_uint64(flag.Arg(2),inputbase)
				if err != nil { error_exit("range step is not a number",2) }
			} else {
				n1, err = get_uint64(flag.Arg(0),inputbase)
				if err != nil { error_exit("range upper limit is not a number",2) }
				n2, err = get_uint64(flag.Arg(1),inputbase)
				if err != nil { error_exit("range lower limit is not a number",2) }
				step, err = get_uint64(flag.Arg(2),inputbase)
				if err != nil { error_exit("range step is not a number",2) }
			}
		}

		// 2 arguments: -r min max
		// step defaults to 1

		if len(flag.Args()) == 2 {
			if inputchar {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n1 = uint64(([]rune)(flag.Arg(0))[0])
				if len(([]rune)(flag.Arg(1))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(1))[0])
			} else {
				n1, err = get_uint64(flag.Arg(0),inputbase)
				if err != nil { error_exit("range lower limit is not a number",2) }
				n2, err = get_uint64(flag.Arg(1),inputbase)
				if err != nil { error_exit("range upper limit is not a number",2) }
			}
		}

		// 1 argument: -r max
		// min and step both default to 1

		if len(flag.Args()) == 1 {
			n1 = 1
			if inputchar {
				if len(([]rune)(flag.Arg(0))) > 1 { error_exit("range upper limit contains more than one character",2) }
				n2 = uint64(([]rune)(flag.Arg(0))[0])
			} else {
				n2, err = get_uint64(flag.Arg(0),inputbase)
				if err != nil { error_exit("range upper limit is not a number",2) }
			}
		}

		sep := ""	// Don't print space before first number

		// for -w option: find maximum width of numbers in sequence

		if constwidth {
		//
			var f string
			num := n2
			if n1 > n2 { num = n1 }

			switch {
				case printhex || printhexc || outputbase == 16: f = "%x"
				case printoct || outputbase == 8:
					if printoct { f = "%#o" } else { f = "%o" }
				case printbin:	f = "%b"
				default:	f = "%d"
			}
			s := fmt.Sprintf(f,num)
			width = len(s)
		}
		
		// When counting downwards from n1 to n2

		if n1 > n2 { downwards = true }

		// Print Output

		i = n1
		for {
			if printhex || printhexc || outputbase == 16 {
				if printhexc {
					// print as hexidecimal number using A-F
					if outputbase == 16 {
						// print without 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%0[3]*X",sep,prefix,width,i)
					} else {
						// default: print with 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%#0[3]*X",sep,prefix,width,i)
					}
				} else {
					// print as hexidecimal number using a-f
					if outputbase == 16 {
						// print without 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%0[3]*x",sep,prefix,width,i)
					} else {
						// default: print with 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%#0[3]*x",sep,prefix,width,i)
					}
				}
			} else if printoct || outputbase == 8 {
				// print as octal number
				if outputbase == 8 {
					// print without 0 prefix
					fmt.Fprintf(os.Stdout,"%s%s%0[3]*o",sep,prefix,width,i)
				} else {
					// default: print with 0 prefix
					fmt.Fprintf(os.Stdout,"%s%s%#0[3]*o",sep,prefix,width,i)
				}
			} else if printbin {
				// print as binary number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*b",sep,prefix,width,i)
			} else if printuni {
				// print as U+HHHH
				fmt.Fprintf(os.Stdout,"%s%s%U",sep,prefix,i)
// TODO: option to allow %#U
			} else if printchar {
				// print as Unicode character
				fmt.Fprintf(os.Stdout,"%s%c",sep,rune(i))
			} else {
				// default: print as decimal number
				fmt.Fprintf(os.Stdout,"%s%s%0[3]*d",sep,prefix,width,i)
			}

			sep = separator	// after printing the first, use a separator for the rest

			// iterating is different depending on stepping upwards or downwards

			if ! downwards {
				i += step
				if i > n2 { break }
			}
			if downwards {
				// i, n2, and step are UNSIGNED, so we can't subtract anything
				// from n2 and then compare to see if it's less than 0.
				if i < n2 + step { break }
				i -= step
			}
		}
		if ! nonewline { fmt.Fprintf(os.Stdout,"\n") }
		os.Exit(0)	// command finished successfully
	}

	// When not doing a sequence (-r, or range), read from standard input or arguments

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
			if inputchar { sep = "" } else { sep = " " }
			s := strings.Join(flag.Args(),sep)
			scanner = bufio.NewScanner(strings.NewReader(s))
		}

		// input runes if inputchar is true (-ic option), otherwise words to convert to integers
		if inputchar { scanner.Split(bufio.ScanRunes) } else { scanner.Split(bufio.ScanWords) }

		sep = ""	// Don't print a separator string before the first number
		printing := false	// false until something has been printed on the line

		for scanner.Scan() {
		//
			st := scanner.Text()
			if inputchar {
				// Get the rune into an integer
				n1 = uint64(([]rune)(scanner.Text())[0])
			} else {
				// Parse number in input base
				n1, err = get_uint64(scanner.Text(),inputbase)

				if err != nil {
					if printing { fmt.Printf("\n") }

					if inputbase == 0 { inputbase = 10 }
					if numError, ok := err.(*strconv.NumError); ok {
						if numError.Err == strconv.ErrRange {
							fmt.Fprintf(os.Stderr,"%s: base %d number '%s' out of range\n",cmdname,inputbase,st)
							os.Exit(2)
						}
					}

					if numError, ok := err.(*strconv.NumError); ok {
						if numError.Err == strconv.ErrSyntax {
							fmt.Fprintf(os.Stderr,"%s: cannot convert '%s' to a number\n",cmdname,st)
							os.Exit(2)
						}
					}

					fmt.Fprintf(os.Stderr,"%s: cannot convert %s to a number\n",cmdname,st)
					os.Exit(2)
				}
			}

			// Print it in the numeric base or as a Unicode character
			if printhex || printhexc || outputbase == 16 {
				if printhexc {
					// print as hexidecimal number using A-F
					if outputbase == 16 {
						// print without 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%0[3]*X",sep,prefix,width,n1)
					} else {
						// default: print with 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%#0[3]*X",sep,prefix,width,n1)
					}
				} else {
					// print as hexidecimal number using a-f
					if outputbase == 16 {
						// print without 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%0[3]*x",sep,prefix,width,n1)
					} else {
						// default: print with 0x prefix
						fmt.Fprintf(os.Stdout,"%s%s%#0[3]*x",sep,prefix,width,n1)
					}
				}
			} else if printoct || outputbase == 8 {
				// print as octal number
				if outputbase == 8 {
					// print without 0 prefix
					fmt.Fprintf(os.Stdout,"%s%s%0[3]*o",sep,prefix,width,n1)
				} else {
					// default: print with 0 prefix
					fmt.Fprintf(os.Stdout,"%s%s%#0[3]*o",sep,prefix,width,n1)
				}
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
			printing = true
		}

		// unless -n is in effect, print a newline when outputting numbers or reading from the command's arguments
// TODO: write this more simply
		if ! nonewline && ( ! printchar || (printchar && len(flag.Args()) > 0)) { fmt.Fprintf(os.Stdout,"\n") }
		os.Exit(0)	// command finished successfully
	}
}
// EOF
