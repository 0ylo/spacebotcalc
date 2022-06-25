package config

import "github.com/spf13/viper"

type Config struct {
	Bot     BotConfig `mapstructure:"bot"`
	Version string
	Build   string
}

type BotConfig struct {
	Token   string `mapstructure:"token"`
	Timeout int    `mapstructure:"timeout"`
	Debug   bool   `mapstructure:"debug"`
}

func Init(version, build string) (*Config, error) {
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("/etc/spacebotcalc")
	viper.SetConfigName("spacebotcalc")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("SPACEBOTCALC")
	viper.BindEnv("TOKEN")

	// Set defaults
	viper.SetDefault("bot.timeout", 60)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	token := viper.GetString("TOKEN")
	if len(token) > 0 {
		cfg.Bot.Token = token
	}

	cfg.Version = version
	cfg.Build = build

	return &cfg, nil
}
