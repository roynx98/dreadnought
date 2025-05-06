package config

import (
	"flag"
)

type FlagLoader struct{}

func (flagLoader FlagLoader) Load() ConfigDTO {
	targetHost := flag.String("host", "", "Your name")
	port := flag.String("port", "", "Port")
	flag.Parse()

	config := ConfigDTO{TargetHost: *targetHost, Port: *port}

	return config
}
