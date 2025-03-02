package config

import "github.com/spf13/viper"

type Config struct {
	LogLevel string `mapstructure:"log_level"`
}

var Cfg *Config

func LoadConfig(cfgFileName string) error {
	v := viper.New()
	v.SetConfigName("XReader Config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if cfgFileName != "" {
		v.SetConfigFile(cfgFileName)
	}
	v.SetDefault("log_level", "info")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		// Use default values when the configuration file does not exist.
	}
	Cfg = &Config{}
	return v.Unmarshal(Cfg)
}
