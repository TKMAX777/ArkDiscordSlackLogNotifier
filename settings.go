package main

import (
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/discord"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack"
)

type Settings struct {
	Slack   slack.Settings
	Discord discord.Settings
}
