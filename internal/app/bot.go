package app

import "github.com/bwmarrin/discordgo"

// ボットstruct
type Bot struct {
	session *discordgo.Session
}

// 初期化関数
func New(token string, intents discordgo.Intent) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	//Intents未指定なら最小のIntentsを与える
	if intents == 0 {
		intents = discordgo.IntentsGuilds
	}
	s.Identify.Intents = intents

	return &Bot{session: s}, nil
}

// 接続開始
func (b *Bot) Start() error {
	return b.session.Open()
}

// 接続終了
func (b *Bot) Stop() {
	_ = b.session.Close()
}

// メッセージ作成イベントのハンドラを登録
func (b *Bot) OnMessageCreate(h func(*discordgo.Session, *discordgo.MessageCreate)) func() {
	return b.session.AddHandler(h)
}
