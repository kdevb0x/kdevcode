// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"io"
	"net"
)

type option struct {
	label      string
	active     bool
	optionList []string
}

var daemon struct {
	listener net.Listener
	config   io.Reader
	options  map[string]option
}

// UpdateVal of an option; err != nil for malformed or nonsensical values
func (t *option) UpdateVal(active bool, val ...string) error {

}

func (t *option) UpdateState() {

}
