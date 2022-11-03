package main

import (
	"flag"

	_ "github.com/lib/pq"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config_sender.toml", "Path to sender configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
}
