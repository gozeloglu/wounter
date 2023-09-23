package main

import "flag"

const (
	usageText = `Wounter is a tool for counting words. 

Usage of %s
	%s [options] <string | filePath>

`

	footerText = `
See github.com/gozeloglu/wounter for bug report.
`
)

var helpFlag = flag.Bool("help", false, "Print usage")
