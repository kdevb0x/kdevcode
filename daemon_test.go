package main

import (
	"log"
	"net"
	"testing"
)

func TestDaemonize(t *testing.T) {
	d := &daemon{
		socket:   "unix",
		listener: net.Listener,
		logger:   log.New(stdin, "kdevcode", 0777),
	}
}
