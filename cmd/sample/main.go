package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/orchidee9392/discordbot-tutorial/internal/app"
	"github.com/orchidee9392/discordbot-tutorial/internal/config"
	"github.com/orchidee9392/discordbot-tutorial/internal/handlers"
)

func main() {
	//設定ロード
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	//使用イベント設定
	intents := discordgo.IntentsGuilds |
	 	discordgo.IntentsGuildMessages |
	  	discordgo.IntentsMessageContent

	//Bot生成
	bot, err := app.New(cfg.Token, intents)
	if err != nil {
		log.Fatal(err)
	}

	//ハンドラ登録
	removeEcho := bot.OnMessageCreate(handlers.Echo)
	defer removeEcho()

	//接続
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
	defer bot.Stop()

	log.Println("接続完了")
	fmt.Println("Enterで終了...")
	fmt.Scanln()
}
