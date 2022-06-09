// Copyright 2022 LoginAuth developers.
//
// SPDX-License-Identifier: GPL-3.0

// Package loginauth provides an smtp.Auth for LOGIN authentication.
package loginauth

import (
	"errors"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

var _ = smtp.Auth(loginAuth{})

var ErrUnknown = errors.New("unknown from server")

func (s loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(s.username), nil
}
func (s loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}
	switch string(fromServer) {
	case "Username:":
		return []byte(s.username), nil
	case "Password:":
		return []byte(s.password), nil
	default:
		return nil, fmt.Errorf("%w: %q", ErrUnknown, string(fromServer))
	}
}

// Auth returns an smtp.Auth that provides LOGIN authentication.
func Auth(username, password string) smtp.Auth {
	return loginAuth{username: username, password: password}
}
