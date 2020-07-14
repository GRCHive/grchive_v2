package main

import (
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
	"gitlab.com/grchive/grchive-v2/shared/envconfig"
	"io/ioutil"
)

type FusionAuthConfig struct {
	Host         string `env:"FUSIONAUTH_HOST,required"`
	Port         int32  `env:"FUSIONAUTH_PORT,required"`
	ExternalHost string `env:"FUSIONAUTH_EXTERNAL_HOST,required"`
	ClientId     string `env:"FUSIONAUTH_CLIENT_ID,required"`
	ClientSecret string `env:"FUSIONAUTH_CLIENT_SECRET,required"`
	ApiKey       string `env:"FUSIONAUTH_API_KEY,required"`
}

type VaultConfig struct {
	Protocol string `env:"VAULT_PROTOCOL,required"`
	Host     string `env:"VAULT_HOST,required"`
	Port     int32  `env:"VAULT_PORT,required"`
	RoleId   string `env:"VAULT_APPROLE_ROLE_ID,required"`
	SecretId string `env:"VAULT_APPROLE_SECRET_ID,required"`
}

type GrchiveConfig struct {
	Domain            string `env:"GRCHIVE_DOMAIN,required"`
	SessionAuthKey    string `env:"GRCHIVE_SESSION_AUTH_KEY,required"`
	SessionEncryptKey string `env:"GRCHIVE_SESSION_ENCRYPT_KEY,required"`
}

type Config struct {
	EnableReleaseMode bool             `toml:"enable_release_mode"`
	FusionAuth        FusionAuthConfig `toml:"fusion_auth"`
	Vault             VaultConfig
	Grchive           GrchiveConfig
}

func (w *WebappApplication) LoadConfig(fname string) {
	w.cfg = w.loadConfigFromFile(fname)
}

func (w *WebappApplication) loadConfigFromFile(fname string) *Config {
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

	cfg.loadConfigFromEnv(w.log)
	return &cfg
}

func (c *Config) loadConfigFromEnv(log zerolog.Logger) {
	if err := envconfig.Load(&c.FusionAuth); err != nil {
		log.Fatal().Err(err).Msg("Failed to read Fusion Auth env config.")
	}
	if err := envconfig.Load(&c.Vault); err != nil {
		log.Fatal().Err(err).Msg("Failed to read Vault env config.")
	}
	if err := envconfig.Load(&c.Grchive); err != nil {
		log.Fatal().Err(err).Msg("Failed to read Grchive env config.")
	}
}
