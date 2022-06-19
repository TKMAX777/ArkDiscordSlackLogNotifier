package discord

import (
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/discord_webhook"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/joinstate"
)

type Handler struct {
	hook *discord_webhook.Handler

	settings Settings

	joinState joinstate.JoinState
}

func New(s Settings) *Handler {
	var hook = discord_webhook.New(s.Token)

	return &Handler{
		settings:  s,
		hook:      hook,
		joinState: joinstate.NewJoinState(),
	}
}

func (h *Handler) SetHookURI(uri string) {
	h.hook.SetHookURI(uri)
}
