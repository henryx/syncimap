/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/url"
)

var (
	src = kingpin.Arg("source", "Set source URI").Required().String()
	dst = kingpin.Arg("destination", "Set destination URI").Required().String()
)

func main() {
	var srcurl, dsturl *url.URL
	var err error

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	srcurl, err = url.Parse(*src)
	if err != nil {
		fmt.Errorf("Source URI is not valid")
	}

	dsturl, err = url.Parse(*dst)
	if err != nil {
		fmt.Errorf("Destination URI is not valid")
	}
}
