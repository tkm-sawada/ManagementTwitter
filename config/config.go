package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey            string
	ApiSecret         string
	AccessToken       string
	AccessTokenSecret string
	BearerToken       string
	LogFile           string
}
type Database struct {
	Date string
	Text string
	Flg  int
}

var Config ConfigList

var DatabaseName = "database.txt"

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:            cfg.Section("twitter_api").Key("api_key").String(),
		ApiSecret:         cfg.Section("twitter_api").Key("api_secret").String(),
		AccessToken:       cfg.Section("twitter_api").Key("access_token").String(),
		AccessTokenSecret: cfg.Section("twitter_api").Key("access_token_secret").String(),
		BearerToken:       cfg.Section("twitter_api").Key("bearer_token").String(),
		LogFile:           cfg.Section("config").Key("log_file").String(),
	}
}

func Readfile() ([]byte, error) {
	data, err := os.ReadFile(DatabaseName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func Savefile(data []byte) error {
	return os.WriteFile(DatabaseName, data, 0600)
}
