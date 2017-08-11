package main

import (
	"os"

	"github.com/meleca/bot/irc"
	"github.com/meleca/bot/slack"
	"github.com/meleca/bot/telegram"
	_ "github.com/meleca/plugins-br/cnpj"
	_ "github.com/meleca/plugins-br/cotacao"
	_ "github.com/meleca/plugins-br/cpf"
	_ "github.com/meleca/plugins-br/dilma"
	_ "github.com/meleca/plugins-br/lula"
	_ "github.com/meleca/plugins-br/megasena"
	_ "github.com/meleca/plugins/9gag"
	_ "github.com/meleca/plugins/catfacts"
	_ "github.com/meleca/plugins/catgif"
	_ "github.com/meleca/plugins/chucknorris"
	_ "github.com/meleca/plugins/crypto"
	_ "github.com/meleca/plugins/encoding"
	_ "github.com/meleca/plugins/example"
	_ "github.com/meleca/plugins/gif"
	_ "github.com/meleca/plugins/godoc"
	_ "github.com/meleca/plugins/guid"
	_ "github.com/meleca/plugins/puppet"
	_ "github.com/meleca/plugins/treta"
	_ "github.com/meleca/plugins/url"
)

func main() {
	filename := os.Args[1]

	var config Config
	err := config.LoadFromFile(filename)
	if err != nil {
		panic(err)
	}

	go irc.Run(config.IRC)
	go slack.Run(config.SlackToken)
	telegram.Run(config.TelegramToken, config.Debug)
}
