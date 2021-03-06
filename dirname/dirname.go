//
// dirname.go (go-coreutils) 0.1
// Copyright (C) 2014, The GO-Coreutils Developers.
//
// Written By: Trey Tacon, Abram C. Isola
//
package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

const (
	help_text = `
    usage: dirname PATH
       or: dirname OPTION
    
    convert a full pathname to just a path

        --help     display this help and exit
        --version  output version information and exit
    `
	version_text = `
    dirname (go-coreutils) 0.1

    Copyright (C) 2014, The GO-Coreutils Developers.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute 
    it under certain conditions in LICENSE.
`
)

var (
	help    = flag.Bool("help", false, help_text)
	version = flag.Bool("version", false, version_text)
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println(help_text)
		return
	}

	if *version {
		fmt.Println(version_text)
		return
	}

	if flag.NArg() >= 1 {
		// NOTE(ttacon): we need to clean the directory first
		// as filepath.Dir will not remove the last directory
		// of the path the way dirname will if it ends in
		// a trailing slash - e.g.:
		//
		// filepath.Dir("/Users/ttacon/") == "/Users/ttacon/"
		// while '$ dirname /Users/ttacon/' == /Users
		fmt.Println(filepath.Dir(filepath.Clean(flag.Arg(0))))
	} else {
		fmt.Println(help_text)
		return
	}
}
