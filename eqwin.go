package main

import (
	"fmt"
	"os"

	"goplan9.googlecode.com/hg/plan9/acme"
)

var envelope string = ".PS\n.fp 1 R R.nomath\n\n.ps 14\n.EQ\n%s\n.EN\n"

func main() {
	w, err := acme.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating window: %s\n", err)
		os.Exit(1)
	}
	w.Name("eqwin")

	c := w.EventChan()
	for {	
		ev := <-c
		if closed(c) {
			break
		}

		// mouse button 2 event in the body
		if ev.C1 == 'M' && ev.C2 == 'L' {
			pr, pw, err := os.Pipe();
			_, err = os.ForkExec("eqshow", []string{"eqshow"}, os.Envs, "",
									[]*os.File{pr, os.Stdout, os.Stderr})
			if err != nil {
				fmt.Fprintf(os.Stderr, "could now fork: %s\n", err)
				os.Exit(1)
			}

			fmt.Fprintf(pw, envelope, ev.Text)
			pw.Close()		// sends EOF to eqshow
		} else {
			w.WriteEvent(ev)
		}
	}
}
