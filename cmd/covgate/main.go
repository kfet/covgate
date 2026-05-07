// covgate filters a Go coverage profile through a .covignore file
// and enforces a minimum statement-weighted coverage percentage.
//
// Usage:
//
//	go tool covgate -profile=coverage.tmp.out \
//	                -out=coverage.out         \
//	                -ignore=.covignore        \
//	                -min=100
//
// See package github.com/kfet/covgate for the filter and gate logic.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kfet/covgate"
)

func main() {
	var cfg covgate.Config
	flag.StringVar(&cfg.ProfilePath, "profile", "", "input coverage profile (required)")
	flag.StringVar(&cfg.OutPath, "out", "", "filtered profile output path (required)")
	flag.StringVar(&cfg.IgnorePath, "ignore", "", "path to .covignore (line-oriented regexes; # comments)")
	flag.Float64Var(&cfg.Min, "min", 100.0, "minimum coverage percent (statement-weighted)")
	flag.Parse()
	cfg.Stdout = os.Stdout
	cfg.Stderr = os.Stderr
	if err := covgate.Run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
