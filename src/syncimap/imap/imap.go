/*
   Copyright (C) 2017 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       Syncimap
   Description   A rsync like IMAP syncronization tool
   License       Apache License 2.0 (see LICENSE for details)
*/

package imap

import (
	"errors"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"net/url"
	"strings"
)

type Connection struct {
	Client *client.Client
}

func (conn *Connection) Dial(uri *url.URL) error {
	var err error

	port := uri.Port()
	switch uri.Scheme {
	case "imap":
		if port == "" {
			port = "143"
		}
		conn.Client, err = client.Dial(uri.Host + ":" + port)
	case "imaps":
		if port == "" {
			port = "993"
		}
		conn.Client, err = client.DialTLS(uri.Host+":"+port, nil)
	default:
		return errors.New("Scheme not supported: " + uri.Scheme)
	}

	caps, err := conn.Client.Capability()
	if err != nil {
		return err
	}

	if caps["STARTTLS"] {
		conn.Client.StartTLS(nil)
	}

	user := uri.User.Username()
	password, _ := uri.User.Password()

	err = conn.Client.Login(user, password)
	if err != nil {
		return err
	}

	return err
}

func (conn *Connection) Folders(folder string) []*imap.MailboxInfo {
	var folders []*imap.MailboxInfo
	mailboxes := make(chan *imap.MailboxInfo)

	done := make(chan error, 1)
	go func() {
		done <- conn.Client.List("", folder, mailboxes)
	}()

	for mailbox := range mailboxes {
		folders = append(folders, mailbox)
	}

	return folders
}

func (conn *Connection) FolderExist(folder string) bool {
	conn.Client.Close()

	_, err := conn.Client.Select(folder, false)
	if err != nil {
		return false
	}

	return true
}

func (conn *Connection) FolderCreate(folder *imap.MailboxInfo) error {

	parent := "INBOX"
	name := folder.Name
	if strings.Contains(name, folder.Delimiter) {
		elements := strings.Split(folder.Name, folder.Delimiter)
		parent = parent + folder.Delimiter + strings.Join(elements[:len(elements)-1], folder.Delimiter)
		name = elements[len(elements)-1]
	}

	conn.Client.Select(parent, false)
	err := conn.Client.Create(folder.Name)

	return err
}
