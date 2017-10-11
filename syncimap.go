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

type Data struct {
	Source      *string
	Destination *string
}

func parse(data *Data) {
	data.Source = kingpin.Arg("source", "Set source URI").Required().String()
	data.Destination = kingpin.Arg("destination", "Set destination URI").Required().String()
}

func main() {
	var data Data
	var srcurl, dsturl *url.URL
	var err error

	parse(&data)

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	srcurl, err = url.Parse(*data.Source)
	if err != nil {
		fmt.Errorf("Source URI is not valid")
	}

	dsturl, err = url.Parse(*data.Destination)
	if err != nil {
		fmt.Errorf("Destination URI is not valid")
	}
}
