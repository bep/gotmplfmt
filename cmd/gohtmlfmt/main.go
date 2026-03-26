package main

import (
	"flag"
	"log"
	"os"

	"github.com/bep/gotmplfmt/tmplfmt"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	inpath := flag.Arg(0)
	outpath := inpath
	if flag.NArg() > 1 {
		outpath = flag.Arg(1)
	}
	buf, err := os.ReadFile(inpath)
	if err != nil {
		log.Fatal(err)
	}
	out, err := tmplfmt.Format(string(buf))
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(outpath, []byte(out), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
