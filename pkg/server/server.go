package server

import (
	"fmt"

	"github.com/bharadwajaD/flappy-go/pkg/game"
)

type Stype int

// Set implements flag.Value.
func (Stype) Set(string) error {
	panic("unimplemented")
}

// String implements flag.Value.
func (Stype) String() string {
	panic("unimplemented")
}

const (
	WS Stype = iota
	TCP
)

type IGameServer interface {
	Run(*game.GameOpts)
}

// Config ...
type Config struct {
	Host string
	Port string
}

func NewGameServer(stype string, config *Config) (IGameServer, error) {
	var server IGameServer
	if stype == "TCP" {
		server = NewTCPServer(config)
	} else if stype == "WS" {
		server = NewWSServer(config)
	} else {
		return nil, fmt.Errorf("undefined stype")
	}

	return server, nil
}
