package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/meleca/bot/irc"
	"github.com/spf13/viper"
)

type Configuration interface {
	LoadFromFile(filename string) error
}

type Config struct {
	Debug         bool
	SlackToken    string
	TelegramToken string
	IRC           *irc.Config
}

func (c *Config) LoadFromFile(filename string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	if c.Debug {
		log.Printf("Configuration loaded: %s\n", c)
	}

	if c.IRC.Password == "" {
		fmt.Printf("Enter the nickserv password: ")
		password, err := terminal.ReadPassword(0)
		fmt.Println()
		if err != nil {
			return err
		}

		c.IRC.Password = string(password)
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("{%t %s %s %v}", c.Debug, c.SlackToken, c.TelegramToken, c.IRC)
}
