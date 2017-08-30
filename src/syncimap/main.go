/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	src = kingpin.Arg("source", "Set source URI").Required().String()
	dst = kingpin.Arg("destination", "Set destination URI").Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
}
