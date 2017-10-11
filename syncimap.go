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
	Source      *url.URL
	Destination *url.URL
}

func parse(data *Data) {
	var err error

	source := kingpin.Arg("source", "Set source URI").Required().String()
	destination := kingpin.Arg("destination", "Set destination URI").Required().String()

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data.Source, err = url.Parse(*source)
	if err != nil {
		fmt.Errorf("Source URI is not valid")
	}

	data.Destination, err = url.Parse(*destination)
	if err != nil {
		fmt.Errorf("Destination URI is not valid")
	}
}

func main() {
	var data Data

	parse(&data)
}
