/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"imap"
	"log"
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
		log.Fatalf("Source URI is not valid")
	}

	data.Destination, err = url.Parse(*destination)
	if err != nil {
		log.Fatalf("Destination URI is not valid")
	}
}

func main() {
	var data Data
	var srcconn, dstconn imap.Connection

	parse(&data)

	if err := srcconn.Dial(data.Source); err != nil {
		log.Fatalf(err.Error())
	}
	defer srcconn.Client.Close()
	defer srcconn.Client.Logout()

	if err := dstconn.Dial(data.Destination); err != nil {
		log.Fatalf(err.Error())
	}
	defer dstconn.Client.Close()
	defer dstconn.Client.Logout()
}
