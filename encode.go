// Copyright (c) 2019, Microsoft Corporation, Sean Hinchee
// Licensed under the MIT License.
package main

import (
	"net/url"
	"os"
	"bufio"
	"log"
	"io"
	"flag"
)

// HTML-encode stdin (or arguments, if provided)
func main() {
	nl := flag.Bool("n", false, "Suppress newline")
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	out := bufio.NewWriter(os.Stdout)
	inputs := []*bufio.Reader{bufio.NewReader(os.Stdin)}

	if argc := len(args); argc > 0 {
		inputs = make([]*bufio.Reader, 0, argc)
		for _, name := range args {
			f, err := os.Open(name)
			if err != nil {
				log.Fatal("err: could not open file - ", err)
			}
			defer f.Close()

			r := bufio.NewReader(f)
			inputs = append(inputs, r)
		}
	}

	// Read input(s) and emit their encoded form
	for i, in := range inputs {
		if in == nil {
			log.Fatal("err: got nil buffer on writer #", i)
		}
		s := ""

		for {
			r, _, err := in.ReadRune()
			if err != nil {
				if err != io.EOF {
					log.Fatal("err: read failed - ", err)
				}
				break
			}
			s += string(r)
		}
		out.Write([]byte(url.QueryEscape(s)))
	}

	if !*nl {
		out.WriteRune('\n')
	}

	out.Flush()
}
