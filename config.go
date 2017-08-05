package main

import (
	"fmt"
	"github.com/go-chat-bot/bot/irc"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(source, c)
	if err != nil {
		return err
	}

	log.Printf("Configuration loaded: %s\n", c)
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
