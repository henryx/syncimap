/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/url"
	"syncimap/imap"
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
	var err error

	parse(&data)

	err = srcconn.Dial(data.Source)
	if err != nil {
		log.Fatalf("Failed to open connection for ", data.Source.Hostname(), ": ", err.Error())
	}
	defer srcconn.Client.Close()
	defer srcconn.Client.Logout()

	err = dstconn.Dial(data.Destination)
	if err != nil {
		log.Fatalf("Failed to open connection for ", data.Source.Hostname(), ": ", err.Error())
	}
	defer dstconn.Client.Close()
	defer dstconn.Client.Logout()
}
