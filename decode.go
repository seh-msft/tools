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

			r := bufio.NewReader(f)
			inputs = append(inputs, r)
		}
	}

	// Read input(s) and emit their decoded form
	for i, in := range inputs {
		if in == nil {
			log.Fatal("Got nil buffer on writer #", i)
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

		u, err := url.QueryUnescape(s)
		if err != nil {
			log.Fatal("err: could not decode - ", err)
		}

		out.Write([]byte(u))
	}

	if !*nl {
		out.WriteRune('\n')
	}

	out.Flush()
}
