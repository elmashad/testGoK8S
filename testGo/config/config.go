package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *viper.Viper;

func LoadConfig()  {

	envconf, errenv := getEnvConfig()

	v := viper.New()
	v.SetConfigType("json")

	if errenv == nil {
		readErr := v.ReadConfig(envconf)
		if readErr != nil {

		}
	} else {

		v.SetConfigName("config")
		v.AddConfigPath("./")
		err := v.ReadInConfig() // Find and read the config file

		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// This is for overriding the default configs behavior
		if _, err := os.Stat("./config.override.json"); err == nil {
			v.SetConfigName("config.override")
			v.MergeInConfig()
		}
	}

	v.WatchConfig()
	Config = v
}

// Read config by key
func Get(configKey string) string {
	return Config.GetString(configKey)
}


func getEnvConfig() (*bytes.Reader, error) {
	envConf := os.Getenv("APP_CONFIG")
	if envConf == "" {
		return nil, errors.New("variable error")
	}
	envConfByte := []byte(envConf)
	r := bytes.NewReader(envConfByte)

	return r, nil
}
