package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/ark"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/discord"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/message_sender"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack"
)

func main() {
	var settings = Settings{
		Slack: slack.Settings{
			Token: os.Getenv("SLACK_TOKEN"),
			SendOptions: slack.SendOptions{
				JoinAndLeftState: os.Getenv("SLACK_SEND_JOIN_AND_LEFT_STATE") == "yes",
				JoinAndLeft:      os.Getenv("SLACK_SEND_JOIN_AND_LEFT") == "yes",
				Other:            os.Getenv("SLACK_SEND_OTHER") == "yes",
				All:              os.Getenv("SLACK_SEND_ALL") == "yes",
			},

			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			AvaterURI: os.Getenv("SLACK_AVATER_URI"),
			UserName:  os.Getenv("SLACK_USERNAME"),

			UserListIconEmoji: os.Getenv("SLACK_USER_EMOJI"),
		},
		Discord: discord.Settings{
			Token: os.Getenv("DISCORD_TOKEN"),
			SendOptions: discord.SendOptions{
				JoinAndLeftState: os.Getenv("DISCORD_SEND_JOIN_AND_LEFT_STATE") == "yes",
				JoinAndLeft:      os.Getenv("DISCORD_SEND_JOIN_AND_LEFT") == "yes",
				Other:            os.Getenv("DISCORD_SEND_OTHER") == "yes",
				All:              os.Getenv("DISCORD_SEND_ALL") == "yes",
			},

			ChannelID: os.Getenv("DISCORD_CHANNEL_ID"),
			AvaterURI: os.Getenv("DISCORD_AVATER_URI"),
			UserName:  os.Getenv("DISCORD_USERNAME"),
		},
	}

	var senders = make([]message_sender.MessageSender, 0)

	if settings.Slack.Token != "" {
		senders = append(senders, slack.New(settings.Slack).SendMessageFunction())
		fmt.Println("Start sending to Slack")
	}

	if settings.Slack.Token != "" {
		senders = append(senders, discord.New(settings.Discord).SendMessageFunction())
		fmt.Println("Start sending to Discord")
	}

	var ARK = ark.New(os.Getenv("ARK_LOG_DIRECTORY"))
	message, err := ARK.Start()
	if err != nil {
		panic(err)
	}

	for mes := range message {
		for _, s := range senders {
			err = s(mes)
			if err != nil {
				log.Printf("Error: %s\n", err)
			}
		}
	}
}
