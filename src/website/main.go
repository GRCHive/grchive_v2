package main

import (
	"flag"
)

var configFname string

func main() {
	flag.StringVar(&configFname, "config", "", "TOML Configuration file.")
	flag.Parse()

	app := WebsiteApplication{}
	defer app.Close()

	app.SetupLogging()
	app.LoadConfig(configFname)
	app.LoadPages()
	app.Run()
}
