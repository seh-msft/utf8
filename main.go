// Copyright (c) 2022, Microsoft Corporation, Sean Hinchee
// Licensed under the MIT License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"
	"unicode/utf8"
)

var (
	nonGraphic = flag.Bool("g", false, "Also print non-'graphic' runes in the table")
	listMode   = flag.Bool("l", false, "list mode instead of table")
	from       = flag.Uint64("from", 20, "Beginning code point for table")
	to         = flag.Uint64("to", 129, "Ending code point for table")
	columns    = flag.Uint64("cols", 5, "Number of columns in the table")
)

// utf8 - provide unicode information on runes
func main() {
	flag.Parse()
	args := flag.Args()

	// List mode
	if *listMode {
		fmt.Print(list())
		return
	}

	// Table mode
	if len(args) < 1 {
		fmt.Print(table(list()))
		return
	}

	// Process arguments
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, arg := range args {
		fmt.Fprintln(out, parse(arg))
	}
}

func table(l string) string {
	var out strings.Builder
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', tabwriter.Debug)
	r := bufio.NewReader(strings.NewReader(l))

	for {
		line, err := r.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		efatal(err, "could not read from list buffer")

		_, err = w.Write(line)
		efatal(err, "could not write to table buffer")
	}
	w.Flush()

	return out.String()
}

func list() string {
	var out strings.Builder
	w := bufio.NewWriter(&out)
	p := "\\u"

	// Inclusive ☺
	next := true
	for i := *from; i <= *to && i < utf8.MaxRune; i++ {
		r := q(fmt.Sprintf("%s%04d", p, i))
		if !*nonGraphic && !unicode.IsGraphic(r) {
			continue
		}

		if next && !*listMode {
			if i-*from != 0 {
				fmt.Fprint(w, "\n")
			}
			fmt.Fprint(w, "|")
			next = false
		}

		fmt.Fprintf(w, "%#U", r)

		if *listMode {
			// List mode is one element per line
			fmt.Fprint(w, "\n")
		} else {
			// Table mode using tabs for tabwriter
			// Trailing tab to next entry
			fmt.Fprint(w, "\t")

			// Column newlines and leading tab
			if ((i-*from)+1)%*columns == 0 {
				next = true
			}
		}
	}

	// Table mode trailing newline
	if !*listMode {
		fmt.Fprint(w, "\n")
	}

	w.Flush()
	return out.String()
}

// Parse out the \u1234 escape into a rune
func q(in string) rune {
	s, err := strconv.Unquote("'" + in + "'")
	efatal(err, "bad escape", in)
	r, _ := utf8.DecodeRuneInString(s)
	return r
}

// Parse commandline arguments
func parse(arg string) string {
	lead, _ := utf8.DecodeRune([]byte(arg))

	switch {
	case lead == '\\':
		// A '\u1234' escape code point
		return fmt.Sprintf("%#U", q(arg))

	case lead == 'u' || lead == 'U':
		// A 'u1234' or 'U+1234' code point
		if strings.Contains(arg, "+") {
			// A U+1234 escape
			return fmt.Sprintf("%#U", q("\\u"+arg[2:]))
		} else {
			// A u1234 escape
			return fmt.Sprintf("%#U", q("\\"+arg))
		}

	case unicode.IsDigit(lead):
		if len(arg) < 2 {
			// A single rune
			// Emit the escape code
			// TODO - maybe more output options
			r, _ := utf8.DecodeLastRuneInString(arg)
			return fmt.Sprintf("%#U", r)
		} else {
			// Strip a leading 0x1234 for 1234, if any
			arg = strings.TrimPrefix(arg, "0x")
			// A unicode code point '1234'
			return fmt.Sprintf("%#U", q("\\u"+arg))
		}

	default:
		// A single rune
		// Emit the escape code
		// TODO - maybe more output options
		r, _ := utf8.DecodeLastRuneInString(arg)
		return fmt.Sprintf("%#U", r)
	}
}

// If error, message and fatal
func efatal(err error, s ...interface{}) {
	if err == nil {
		return
	}
	pre := append([]interface{}{"err:"}, s...)
	fatal(pre, "→", err.Error())
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
