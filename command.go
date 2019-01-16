// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"os"

	"github.com/cosiner/flag"
	"github.com/spf13/pflag"
)

const stdin = os.Stdin

// flags
var (
	flagset = []pflag.Flag
)

func ParseCommandFlags() error {
	var socketType, configFile, daemonize, debug flag.Flag
	socketType.
}
