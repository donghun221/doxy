package config

import (
	constants "HelloTencent/internal/utils"
	"slconfig/src/slconfig"
	"errors"
	"fmt"
)

const (
	TESTVAL = "testval"
)

type Config struct {
	source slconfig.ConfigInterface
	domain string
}

func NewConfig(path, domain string) (*Config, error) {
	config, err := slconfig.NewConfig(path)

	if err != nil {
		return nil, err
	}

	conf := Config{config, domain}
	return &conf, nil
}

func (conf *Config) GetTestVal() string {
	return conf.source.String(conf.getFullKey(TESTVAL))
}

// Dynamically override config
// Please do not try to add a new config dynamically
func (conf *Config) Override(key, value string) error {
	return conf.source.Set(key, value)
}

func (conf *Config) SetDomain(in string) error {
	if len(in) == 0 {
		return errors.New(fmt.Sprintf("domain cannot be empty"))
	}

	conf.domain = in
	return nil
}

func (conf *Config) getFullKey(key string) string {
	return conf.domain + constants.DOUBLE_COLON + key
}