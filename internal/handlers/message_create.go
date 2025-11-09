package handlers

import "github.com/bwmarrin/discordgo"

//オウム返しするイベント
func Echo(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
}
