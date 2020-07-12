package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type Config struct {
	EnableAnalytics   bool `toml:"enable_analytics"`
	EnableReleaseMode bool `toml:"enable_release_mode"`
}

func (w *WebsiteApplication) loadConfigFromFile(fname string) *Config {
	w.log.Debug().Msgf("Load Config From %s", fname)
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		w.log.Fatal().Err(err).Msg("Failed to read config.")
	}

	tomlConfig, err := toml.Load(string(dat))
	if err != nil {
		w.log.Fatal().Err(err).Msg("Failed to load TOML.")
	}

	cfg := Config{}
	err = tomlConfig.Unmarshal(&cfg)
	if err != nil {
		w.log.Fatal().Err(err).Msg("Failed to unmarshal TOML.")
	}

	return &cfg
}
