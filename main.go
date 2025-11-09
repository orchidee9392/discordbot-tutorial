package main

import(
	"fmt"
	"log"
	"os"
	"github.com/bwmarrin/discordgo"
)

func main(){
	//トークンの取得
	token := os.Getenv("DISCORD_TOKEN")
	if token == ""{
		log.Fatal("DISCORD_TOKEN が未設定です")
	}

	//Discordセッションの生成
	session, err := discordgo.New("Bot " + token)
	if err != nil{
		log.Fatal(err)
	}

	//使用イベント設定
	session.Identify.Intents = discordgo.IntentsGuilds

	//セッション開始
	if err := session.Open(); err != nil{
		log.Fatal(err)
	}
	defer session.Close()

	//ログ&終了処理
	log.Println("botが接続されました。")
	fmt.Println("Enterで終了...")
	fmt.Scanln()
}