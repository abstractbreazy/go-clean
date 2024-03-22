package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type HTTP struct {
	Host         string        `yaml:"host" mapstructure:"host"`
	Port         string        `yaml:"port" mapstructure:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

type Limiter struct {
	Limit      uint64        `yaml:"limit" mapstructure:"limit"`
	Timeout    time.Duration `yaml:"timeout" mapstructure:"timeout"`
	SubnetMask string        `yaml:"subnet_mask" mapstructure:"subnet_mask"`
}

type Redis struct {
	URL string `yaml:"url" mapstructure:"url"`
}

type Config struct {
	HTTP    *HTTP    `yaml:"http" mapstructure:"http"`
	Limiter *Limiter `yaml:"limiter" mapstructure:"limiter"`
	Redis   *Redis   `yaml:"redis" mapstructure:"redis"`
}

func New() *Config {
	return &Config{
		HTTP:    &HTTP{},
		Limiter: &Limiter{},
		Redis:   &Redis{},
	}
}

func (c *Config) Init() (*Config, error) {
	var v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./internal/config")
	v.AddConfigPath("../../internal/config")
	v.SetEnvPrefix("limiter")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	v.AutomaticEnv()
	err = v.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
