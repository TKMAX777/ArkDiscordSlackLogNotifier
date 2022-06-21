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
				JoinAndLeftState: slack.SendOption{
					IsEnabled: os.Getenv("SLACK_SEND_JOIN_AND_LEFT_STATE") == "yes",
					Emoji:     os.Getenv("SLACK_JOIN_AND_LEFT_STATE_EMOJI"),
				},
				JoinAndLeft: slack.SendOption{
					IsEnabled: os.Getenv("SLACK_SEND_JOIN_AND_LEFT") == "yes",
					Emoji:     os.Getenv("SLACK_JOIN_EMOJI"),
					EmojiSub:  os.Getenv("SLACK_LEFT_EMOJI"),
				},
				Killed: slack.SendOption{
					IsEnabled: os.Getenv("SLACK_SEND_KILLED") == "yes",
					Emoji:     os.Getenv("SLACK_KILLED_EMOJI"),
				},
				Other: slack.SendOption{
					IsEnabled: os.Getenv("SLACK_SEND_OTHER") == "yes",
					Emoji:     os.Getenv("SLACK_OTHER_EMOJI"),
				},
				All: slack.SendOption{
					IsEnabled: os.Getenv("SLACK_SEND_ALL") == "yes",
				},
			},

			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			AvaterURI: os.Getenv("SLACK_AVATER_URI"),
			UserName:  os.Getenv("SLACK_USERNAME"),
		},
		Discord: discord.Settings{
			Token: os.Getenv("DISCORD_TOKEN"),
			SendOptions: discord.SendOptions{
				JoinAndLeftState: discord.SendOption{
					IsEnabled: os.Getenv("DISCORD_SEND_JOIN_AND_LEFT_STATE") == "yes",
					Emoji:     os.Getenv("DISCORD_JOIN_AND_LEFT_STATE_EMOJI"),
				},
				JoinAndLeft: discord.SendOption{
					IsEnabled: os.Getenv("DISCORD_SEND_JOIN_AND_LEFT") == "yes",
					Emoji:     os.Getenv("DISCORD_JOIN_EMOJI"),
					EmojiSub:  os.Getenv("DISCORD_LEFT_EMOJI"),
				},
				Killed: discord.SendOption{
					IsEnabled: os.Getenv("DISCORD_SEND_KILLED") == "yes",
					Emoji:     os.Getenv("DISCORD_KILLED_EMOJI"),
				},
				Other: discord.SendOption{
					IsEnabled: os.Getenv("DISCORD_SEND_OTHER") == "yes",
					Emoji:     os.Getenv("DISCORD_OTHER_EMOJI"),
				},
				All: discord.SendOption{
					IsEnabled: os.Getenv("DISCORD_SEND_ALL") == "yes",
				},
			},
			HookURI:   os.Getenv("DISCORD_HOOK_URI"),
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

	if settings.Discord.Token != "" || settings.Discord.HookURI != "" {
		var handler = discord.New(settings.Discord)
		handler.SetHookURI(settings.Discord.HookURI)
		senders = append(senders, handler.SendMessageFunction())
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
