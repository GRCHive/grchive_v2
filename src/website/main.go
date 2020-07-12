package main

import (
	"flag"
)

var configFname string

func main() {
	flag.StringVar(&configFname, "config", "", "TOML Configuration file.")
	flag.Parse()

	app := WebsiteApplication{}
	app.setupLogging()
	app.loadConfig(configFname)
	app.loadPages()
	app.run()
}
