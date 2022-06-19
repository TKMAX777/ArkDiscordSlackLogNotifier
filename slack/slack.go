package slack

import (
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/joinstate"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack_webhook"
)

type Handler struct {
	hook *slack_webhook.Handler

	settings Settings

	joinState     joinstate.JoinState
	lastMessageTS string
}

func New(s Settings) *Handler {
	var hook = slack_webhook.New(s.Token)

	return &Handler{
		settings:  s,
		joinState: joinstate.NewJoinState(),
		hook:      hook,
	}
}
