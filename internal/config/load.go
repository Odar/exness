package config

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load() (Options, error) {
	cfg, err := load()
	if err != nil {
		return defaultOptions, errors.Wrap(err, "can not load config")
	}

	return *cfg, nil
}

func load() (*Options, error) {
	path, err := getFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "get file path")
	}
	if path == "" {
		return nil, errors.New("empty path")
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Options{}
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "can not unmarshal config from file to struct")
	}

	return &config, nil
}

func getFilePath() (string, error) {
	flag.String("config", "dev.yaml", "path to config file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return "", err
	}
	viper.AutomaticEnv()

	return viper.GetString("config"), nil
}
