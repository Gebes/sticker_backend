package env

import (
	"gebes.io/sticker_backend/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

var (
	PostgresDatabase string

	DiscordClientId          string
	DiscordClientSecret      string
	DiscordClientRedirectUrl string
	DestinationUrl           string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error.Fatalln("Unable to load environment:", err)
	}

	PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
	DiscordClientId = os.Getenv("DISCORD_CLIENT_ID")
	DiscordClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	DiscordClientRedirectUrl = os.Getenv("DISCORD_CLIENT_REDIRECT_URL")
	DestinationUrl = os.Getenv("DESTINATION_URL")

}
