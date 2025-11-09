package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	//トークンの取得
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN が未設定です")
	}

	//Discordセッションの生成
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	//使用イベントの設定
	session.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent

	//メッセージ受信ハンドラを登録
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Botの発言には反応しない（無限ループ防止）
		if m.Author.Bot {
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
	})

	//セッション開始
	if err := session.Open(); err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//ログ&終了処理
	log.Println("botが接続されました。")
	fmt.Println("Enterで終了...")
	fmt.Scanln()
}
