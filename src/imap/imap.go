/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package imap

import (
	"errors"
	"github.com/emersion/go-imap/client"
	"net/url"
)

type Connection struct {
	Client *client.Client
}

func (conn *Connection) Dial(uri *url.URL) error {
	var err error

	switch uri.Scheme {
	case "imap":
		port := uri.Port()
		conn.Client, err = client.Dial(uri.Host + ":" + port)
	case "imaps":
		conn.Client, err = client.DialTLS(uri.Host+":"+uri.Port(), nil)
	default:
		return errors.New("Scheme not supported: " + uri.Scheme)
	}

	return err
}
