package config

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

//Configuration : struct to hold all configuration model
type Configuration struct {
	Database Database
	Server   Server
}

//New : Instantiate new Configuration instance
func New(customPath string) (*Configuration, error) {
	path := "./shared/config/"
	if customPath != "" {
		path = customPath
	}

	viper.SetConfigName("default")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		glog.Fatal(err.Error())
		return nil, err
	}

	cfg := new(Configuration)
	if err := viper.Unmarshal(cfg); err != nil {
		glog.Fatalf("Failed to deserialize config struct: %s", err)
		return nil, err
	}

	return cfg, nil
}
