package config

import (
	"fmt"
	"os"
)

//設定struct
type Config struct {
	Token string
}

//環境変数を読み込んで設定structを作成します
func FromEnv() (Config, error){
	t := os.Getenv("DISCORD_TOKEN")
	if t == "" {
		return Config{}, fmt.Errorf("DISCORD_TOKEN が設定されていません。")
	}

	return Config{Token: t}, nil
}