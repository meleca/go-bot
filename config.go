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

/* Request an user input that might be secret or not, depending on 'secret' boolean parameter.
In case 'secret' is true, the user input won't be presented on the terminal. */
func requestInput(message string, secret bool) (string, error) {
	var input string
	var err error

	fmt.Print(message)
	if secret == true {
		var hidenInput []byte

		hidenInput, err = terminal.ReadPassword(0)
		if err != nil {
			return "", err
		}
		input = string(hidenInput)
	} else {
		fmt.Scan(input)
	}
	fmt.Println()
	return input, err
}

/* For now just secrets (password and tokens) are supported to be stored as env vars.
Read user's env vars and if it's also empty request input from the user. The env vars are:
	GOBOT_IRC_PASSWORD
	GOBOT_SLACK_TOKEN
	GOBOT_TELEGRAM_TOKEN
*/
func (c *Config) loadSecretsFromEnv() error {
	var err error

	viper.SetEnvPrefix("gobot")
	viper.BindEnv("irc_password")
	viper.BindEnv("slack_token")
	viper.BindEnv("telegram_token")

	if c.IRC.Password == "" {
		value := viper.GetString("irc_password")
		if value != "" {
			c.IRC.Password = value
			if c.Debug {
				log.Print("IRC password env variable found")
			}
		} else {
			message := "Enter the nickserv password: "
			c.IRC.Password, err = requestInput(message, true)
			if err != nil {
				return err
			}
		}
	}

	if c.SlackToken == "" {
		value := viper.GetString("slack_token")
		if value != "" {
			c.SlackToken = value
			if c.Debug {
				log.Print("Slack token env variable found")
			}
		} else {
			message := "Insert Slack token ID:"
			c.SlackToken, err = requestInput(message, true)
			if err != nil {
				return err
			}
		}
	}

	if c.TelegramToken == "" {
		value := viper.GetString("telegram_token")
		if value != "" {
			c.TelegramToken = value
			if c.Debug {
				log.Print("Telegram token env variable found")
			}
		} else {
			message := "Insert Telegram token ID: "
			c.TelegramToken, err = requestInput(message, true)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

/* Read config file based on its type and let 'viper' library does the magic here */
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
		log.Print("Configuration loaded from file: ", c)
	}

	if c.IRC.Password == "" || c.SlackToken == "" || c.TelegramToken == "" {
		err = c.loadSecretsFromEnv()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("{%t %s %s %v}", c.Debug, c.SlackToken, c.TelegramToken, c.IRC)
}
